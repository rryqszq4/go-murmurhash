package murmurhash

func MurmurHash1 (key []byte, seed uint32) (hash uint32) {
	const m = 0xc6a4a793
	const r = 16
	l := len(key)
	h := seed ^ uint32(l * m)
	data := key

	for l >= 4 {
		k := uint32(data[0]) + uint32(data[1]) << 8 + uint32(data[2]) << 16 + uint32(data[3]) << 24
		h += k
		h *= m
		h ^= h >> 16

		data = data[4:]
		l -= 4
	}

	switch l {
	case 3:
		h += uint32(data[2]) << 16
		fallthrough
	case 2:
		h += uint32(data[1]) << 8
		fallthrough
	case 1:
		h += uint32(data[0])
		h *= m 
		h ^= h >> r
	}

	h *= m
	h ^= h >> 10
	h *= m
	h ^= h >> 17

	return h
}
