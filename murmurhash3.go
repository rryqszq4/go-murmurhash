package murmurhash

func rotl32(x uint32, r int8) (z uint32) {
	s := uint8(r) & (32 - 1)
	return ( x << s ) | ( x >>(32 - s) )
}

func fmix32(h uint32) (z uint32) {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return h
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
		k1 = uint32(data[i*4+0]) + uint32(data[i*4+1]) << 8 + uint32(data[i*4+2]) << 16 + uint32(data[i*4+3]) << 24
		k2 = uint32(data[i*4+1]) + uint32(data[i*4+2]) << 8 + uint32(data[i*4+3]) << 16 + uint32(data[i*4+4]) << 24
		k3 = uint32(data[i*4+2]) + uint32(data[i*4+3]) << 8 + uint32(data[i*4+4]) << 16 + uint32(data[i*4+5]) << 24
		k4 = uint32(data[i*4+3]) + uint32(data[i*4+4]) << 8 + uint32(data[i*4+5]) << 16 + uint32(data[i*4+6]) << 24

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
