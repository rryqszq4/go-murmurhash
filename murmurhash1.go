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

func MurmurHash1Aligned(key []byte, seed uint32) (hash uint32) {
	const m uint32 = 0xc6a4a793
	const r = 16

	l := len(key)
	data := key
	var h uint32 = seed ^ (uint32(l) * m)
	var align int = int(uint64(data[0])) & 3

	if align != 0 && l >= 4 {
		var t, d uint32 = 0, 0

		switch align {
			case 1: 
				t |= uint32(data[2]) << 16
				fallthrough
			case 2:
				t |= uint32(data[1]) << 8
				fallthrough
			case 3:
				t |= uint32(data[0])
		}

		t <<= 8 * uint32(align)

		data = data[4 - align:]
		l -= 4 - align

		var sl int = 8 * (4 - align)
		var sr int = 8 * align

		for l >= 4 {
			d = uint32(data[0]) + uint32(data[1]) << 8 + 
				uint32(data[2]) << 16 + uint32(data[3]) << 24
			t = (t >> uint32(sr)) | (d << uint32(sl))
			h += t
			h *= m 
			h ^= h >> r
			t = d

			data = data[4:]
			l -= 4
		}

		var pack int
		if l < align {
			pack = l
		} else {
			pack = align
		}

		d = 0

		switch pack {
		case 3:
			d |= uint32(data[2]) << 16
			fallthrough
		case 2:
			d |= uint32(data[1]) << 8
			fallthrough
		case 1:
			d |= uint32(data[0])
			fallthrough
		case 0:
			h += (t >> uint32(sr)) | (d << uint32(sl))
			h *= m 
			h ^= h >> r
		}

		data = data[pack:]
		l -= pack
	} else {
		for l >= 4 {
			h += uint32(data[0]) + uint32(data[1]) << 8 + 
				uint32(data[2]) << 16 + uint32(data[3]) << 24
		}
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
