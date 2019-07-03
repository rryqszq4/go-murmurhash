package main

import (
	"testing"
	. "../../go-murmurhash"
)

func TestMurmurHash2(t *testing.T) {
	var tests = []struct {
		seed uint32
		key string
		hash uint32
	}{
		{0x12345678, "foo", 1532637697},
		{0x12345678, "foofoo", 1972356138},
		{0x12345678, "foofoofoofoofoofoofoofoo", 925860986},
	}

	for _, test := range tests {
		hash := MurmurHash2([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func TestMurmurHash64A(t *testing.T) {
	var tests = []struct {
		seed uint64
		key string
		hash uint64
	}{
		{0x12345678, "foo", 14576140976603494442},
		{0x12345678, "foofoo", 3164582456347398077},
		{0x12345678, "foofoofoofoofoofoofoofoo", 10934192187992172313},
	}

	for _, test := range tests {
		hash := MurmurHash64A([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func TestMurmurHash64B(t *testing.T) {
	var tests = []struct {
		seed uint64
		key string
		hash uint64
	}{
		{0x12345678, "foo", 7448014171559636702},
		{0x12345678, "foofoo", 6143656703560561235},
		{0x12345678, "foofoofoofoofoofoofoofoo", 11180785033823365454},
	}

	for _, test := range tests {
		hash := MurmurHash64B([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func TestMurmurHash2A(t *testing.T){
	var tests = []struct {
		seed uint32
		key string
		hash uint32
	}{
		{0x12345678, "foo", 2221687037},
		{0x12345678, "foofoo", 1924592810},
		{0x12345678, "foofoofoofoofoofoofoofoo", 1984285793},
	}

	for _, test := range tests {
		hash := MurmurHash2A([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func TestMurmurHashNeutral2(t *testing.T){
	var tests = []struct {
		seed uint32
		key string
		hash uint32
	}{
		{0x12345678, "foo", 1532637697},
		{0x12345678, "foofoo", 1972356138},
		{0x12345678, "foofoofoofoofoofoofoofoo", 925860986},
	}

	for _, test := range tests {
		hash := MurmurHashNeutral2([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func TestMurmurHashAligned2(t *testing.T){
	var tests = []struct {
		seed uint32
		key string
		hash uint32
	}{
		{0x12345678, "foo", 1532637697},
		{0x12345678, "foofoo", 1972356138},
		{0x12345678, "foofoofoofoofoofoofoofoo", 925860986},
	}

	for _, test := range tests {
		hash := MurmurHashAligned2([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func BenchmarkMurmurHash2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash2([]byte("foo"), 0x12345678)
	}
}

func BenchmarkMurmurHash64A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash64A([]byte("foo"), 0x12345678)
	}
}

func BenchmarkMurmurHash64B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash64B([]byte("foo"), 0x12345678)
	}
}

func BenchmarkMurmurHash2A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash2A([]byte("foo"), 0x12345678)
	}
}

func BenchmarkMurmurHashNeutral2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHashNeutral2([]byte("foo"), 0x12345678)
	}
}

func BenchmarkMurmurHashAligned2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHashAligned2([]byte("foo"), 0x12345678)
	}
}
