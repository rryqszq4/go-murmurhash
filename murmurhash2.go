package murmurhash

func MurmurHash2 (key []byte, seed uint32) (hash uint32) {
	const m uint32 = 0x5bd1e995
	const r = 24

	var l int = len(key)
	var h uint32 = seed ^ uint32(l)

	var data []byte = key

	var k uint32

	for l >= 4 {
		k = uint32(data[0]) + uint32(data[1]) << 8 + uint32(data[2]) << 16 + uint32(data[3]) << 24

		k *= m
		k ^= k >> r
		k *= m

		h *= m
		h ^= k

		data = data[4:]
		l -= 4
	}

	switch l {
	case 3:
		h ^= uint32(data[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(data[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(data[0])
		h *= m
	}

	h ^= h >> 13
	h *= m
	h ^= h >> 15

	return h
}

func MurmurHash64A (key []byte, seed uint64) (hash uint64) {
	const m uint64 = 0xc6a4a7935bd1e995
	const r = 47

	var l int = len(key)
	var h uint64 = seed ^ uint64(l) * m

	var data []byte = key
	var l8 int = l / 8

	var k uint64

	for i := 0; i < l8; i++ {
		i8 := i * 8
		k = uint64(data[i8 + 0]) + uint64(data[i8 + 1]) << 8 + 
			uint64(data[i8 + 2]) << 16 + uint64(data[i8 + 3]) << 24 + 
			uint64(data[i8 + 4]) << 32 + uint64(data[i8 + 5]) << 40 + 
			uint64(data[i8 + 6]) << 48 + uint64(data[i8 + 7]) << 56

		k *= m
		k ^= k >> r
		k *= m

		h ^= k
		h *= m
	}

	switch l & 7 {
	case 7:
		h ^= uint64(data[6]) << 48
		fallthrough
	case 6:
		h ^= uint64(data[5]) << 40
		fallthrough
	case 5:
		h ^= uint64(data[4]) << 32
		fallthrough
	case 4:
		h ^= uint64(data[3]) << 24
		fallthrough
	case 3:
		h ^= uint64(data[2]) << 16
		fallthrough
	case 2:
		h ^= uint64(data[1]) << 8
		fallthrough
	case 1:
		h ^= uint64(data[0])
		h *= m
	}

	h ^= h >> r
	h *= m
	h ^= h >> r

	return h
}

func MurmurHash64B(key []byte, seed uint64) (hash uint64) {
	const m uint32 = 0x5bd1e995
	const r = 24

	var l int = len(key)
	var h1 uint32 = uint32(seed) ^ uint32(l)
	var h2 uint32 = uint32(seed) >> 32

	var data []byte = key

	var k1, k2 uint32

	for l >= 8 {
		k1 = uint32(data[0]) + uint32(data[1]) << 8 + uint32(data[2]) << 16 + uint32(data[3]) << 24
		k1 *= m; k1 ^= k1 >> r; k1 *=m
		h1 *= m; h1 ^= k1
		data = data[4:]
		l -= 4

		k2 = uint32(data[0]) + uint32(data[1]) << 8 + uint32(data[2]) << 16 + uint32(data[3]) << 24
		k2 *= m; k2 ^= k2 >> r; k2 *= m
		h2 *= m; h2 *= k2
		data = data[4:]
		l -= 4
	}

	if l >= 4 {
		k1 = uint32(data[0]) + uint32(data[1]) << 8 + uint32(data[2]) << 16 + uint32(data[3]) << 24
		k1 *= m; k1 ^= k1 >> r; k1 *= m
		h1 *= m; h1 ^= k1
		l -= 4
	}

	switch l {
	case 3:
		h2 ^= uint32(data[2]) << 16
		fallthrough
	case 2:
		h2 ^= uint32(data[1]) << 8
		fallthrough
	case 1:
		h2 ^= uint32(data[0])
		h2 *= m
	}

	h1 ^= h2 >> 18; h1 *= m
	h2 ^= h1 >> 22; h2 *= m
	h1 ^= h2 >> 17; h1 *= m
	h2 ^= h1 >> 19; h2 *= m

	var h uint64 = uint64(h1)

	h = (h << 32) | uint64(h2)

	return h

}
