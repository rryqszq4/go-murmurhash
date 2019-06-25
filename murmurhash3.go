package murmurhash

func rotl32(x uint32, r int8) (z uint32) {
	s := uint8(r) & (32 - 1)
	return ( x << s ) | ( x >> (32 - s) )
}

func rotl64(x uint64, r int8) (z uint64) {
	s := uint8(r) & (64 - 1)
	return ( x << s ) | ( x >> (64 - s) )
}

func fmix32(h uint32) (z uint32) {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return h
}

func fmix64(k uint64) (z uint64) {
	k ^= k >> 33
	k *= 0xff51afd7ed558ccd
	k ^= k >> 33
	k *= 0xc4ceb9fe1a85ec53
	k ^= k >> 33

	return k
}

// MurmurHash3_x86_32
func MurmurHash3_x86_32(key []byte, seed uint32) (hash uint32) {
	var l int = len(key)
	var nblocks = l / 4

	var h1 uint32 = seed

	const c1 uint32 = 0xcc9e2d51
	const c2 uint32 = 0x1b873593

	var data []byte = key
	var k1 uint32

	for i := 0; i < nblocks; i++ {
		k1 = uint32(data[i*4+0]) + uint32(data[i*4+1]) << 8 + uint32(data[i*4+2]) << 16 + uint32(data[i*4+3]) << 24

		k1 *= c1
		k1 = rotl32(k1, 15)
		k1 *= c2

		h1 ^= k1
		h1 = rotl32(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}

	tail := data[nblocks*4:]

	switch l & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1; k1 = rotl32(k1, 15); k1 *= c2; h1 ^= k1
	}

	h1 ^= uint32(l)

	h1 = fmix32(h1)

	return h1
}

// MurmurHash3_x86_128
func MurmurHash3_x86_128(key []byte, seed uint32) (hash [4]uint32) {
	var l int = len(key)
	var nblocks = l / 16
	var data []byte = key

	var h1 uint32 = seed
	var h2 uint32 = seed
	var h3 uint32 = seed
	var h4 uint32 = seed

	const c1 uint32 = 0x239b961b
	const c2 uint32 = 0xab0e9789
	const c3 uint32 = 0x38b34ae5
	const c4 uint32 = 0xa1e38b93

	var k1,k2,k3,k4 uint32

	for i := 0; i < nblocks; i++ {
		k1 = uint32(data[i*16+0]) + uint32(data[i*16+1]) << 8 + uint32(data[i*16+2]) << 16 + uint32(data[i*16+3]) << 24
		k2 = uint32(data[i*16+4]) + uint32(data[i*16+5]) << 8 + uint32(data[i*16+6]) << 16 + uint32(data[i*16+7]) << 24
		k3 = uint32(data[i*16+8]) + uint32(data[i*16+9]) << 8 + uint32(data[i*16+10]) << 16 + uint32(data[i*16+11]) << 24
		k4 = uint32(data[i*16+12]) + uint32(data[i*16+13]) << 8 + uint32(data[i*16+14]) << 16 + uint32(data[i*16+15]) << 24

		k1 *= c1; k1 = rotl32(k1, 15); k1 *= c2; h1 ^= k1
		h1 = rotl32(h1, 19); h1 += h2; h1 = h1 * 5 + 0x561ccd1b
		k2 *= c2; k2 = rotl32(k2, 16); k2 *= c3; h2 ^= k2
		h2 = rotl32(h2, 17); h2 += h3; h2 = h2 * 5 + 0x0bcaa747
		k3 *= c3; k3 = rotl32(k3, 17); k3 *= c4; h3 ^= k3
		h3 = rotl32(h3, 15); h3 += h4; h3 = h3 * 5 + 0x96cd1c35
		k4 *= c4; k4 = rotl32(k4, 18); k4 *= c1; h4 ^= k4
		h4 = rotl32(h4, 13); h4 += h1; h4 = h4 * 5 + 0x32ac3b17
	}

	tail := data[nblocks*16:]
	
	k1 = 0; k2 = 0; k3 = 0; k4 = 0

	switch l & 15 {
	case 15:
		k4 ^= uint32(tail[14]) << 16
		fallthrough
	case 14:
		k4 ^= uint32(tail[13]) << 8
		fallthrough
	case 13:
		k4 ^= uint32(tail[12]) << 0
		k4 *= c4; k4 = rotl32(k4, 18); k4 *= c1; h4 ^= k4
		fallthrough
	case 12:
		k3 ^= uint32(tail[11]) << 24
		fallthrough
	case 11:
		k3 ^= uint32(tail[10]) << 16
		fallthrough
	case 10:
		k3 ^= uint32(tail[9]) << 8
		fallthrough
	case 9:
		k3 ^= uint32(tail[8]) << 0
		k3 *= c3; k3 = rotl32(k3, 17); k3 *= c4; h3 ^= k3
		fallthrough
	case 8:
		k2 ^= uint32(tail[7]) << 24
		fallthrough
	case 7:
		k2 ^= uint32(tail[6]) << 16
		fallthrough
	case 6:
		k2 ^= uint32(tail[5]) << 8
		fallthrough
	case 5:
		k2 ^= uint32(tail[4]) << 0
		k2 *= c2; k2 = rotl32(k2, 16); k2 *= c3; h2 ^= k2
		fallthrough
	case 4:
		k1 ^= uint32(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0]) << 0
		k1 *= c1; k1 = rotl32(k1, 15); k1 *= c2; h1 ^= k1
	}

	h1 ^= uint32(l); h2 ^= uint32(l); h3 ^= uint32(l); h4 ^= uint32(l)

	h1 += h2; h1 += h3; h1 += h4
	h2 += h1; h3 += h1; h4 += h1

	h1 = fmix32(h1)
	h2 = fmix32(h2)
	h3 = fmix32(h3)
	h4 = fmix32(h4)

	h1 += h2; h1 += h3; h1 += h4
	h2 += h1; h3 += h1; h4 += h1

	return [4]uint32{h1,h2,h3,h4}
}

// MurmurHash3_x64_128
func MurmurHash3_x64_128(key []byte, seed uint64) (hash [2]uint64) {
	var l int = len(key)
	var nblocks = l / 16
	var data []byte = key

	var h1 uint64 = seed
	var h2 uint64 = seed

	const c1 uint64 = 0x87c37b91114253d5
	const c2 uint64 = 0x4cf5ad432745937f

	var k1,k2 uint64

	for i := 0; i < nblocks; i++ {
		k1 := uint64(data[i*16+0]) + uint64(data[i*16+1]) << 8 + 
			uint64(data[i*16+2]) << 16 + uint64(data[i*16+3]) << 24 +
			uint64(data[i*16+4]) << 32 + uint64(data[i*16+5]) << 40 + 
			uint64(data[i*16+6]) << 48 + uint64(data[i*16+7]) << 56
		k2 := uint64(data[i*16+8]) + uint64(data[i*16+9]) << 8 + 
			uint64(data[i*16+10]) << 16 + uint64(data[i*16+11]) << 24 +
			uint64(data[i*16+12]) << 32 + uint64(data[i*16+13]) << 40 + 
			uint64(data[i*16+14]) << 48 + uint64(data[i*16+15]) << 56

		k1 *= c1; k1 = rotl64(k1, 31); k1 *= c2; h1 ^= k1
		h1 = rotl64(h1, 27); h1 += h2; h1 = h1 * 5 + 0x52dce729
		k2 *= c2; k2 = rotl64(k2, 33); k2 *= c1; h2 ^= k2
		h2 = rotl64(h2, 31); h2 += h1; h2 = h2 * 5 + 0x38495ab5
	}

	tail := data[nblocks*16:]

	k1 = 0; k2 = 0

	switch l & 15 {
	case 15:
		k2 ^= uint64(tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(tail[8]) << 0
		k2 *= c2; k2 = rotl64(k2, 33); k2 *= c1; h2 ^= k2
		fallthrough
	case 8:
		k1 ^= uint64(tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(tail[0]) << 0
		k1 *= c1; k1 = rotl64(k1, 31); k1 *= c2; h1 ^= k1
	}

	h1 ^= uint64(l); h2 ^= uint64(l)

	h1 += h2; h2 += h1

	h1 = fmix64(h1); h2 = fmix64(h2)

	h1 += h2; h2 += h1

	return [2]uint64{h1, h2}
}
