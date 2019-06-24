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
