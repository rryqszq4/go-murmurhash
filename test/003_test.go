package main 

import (
	"testing"
	. "../../go-murmurhash"
)

func TestMurmurHash3_x86_32(t *testing.T) {
	var tests = []struct {
		seed uint32
		key string
		hash uint32
	}{
		{0x12345678, "foo", 4084948345},
		{0x12345678, "foofoo", 3208788144},
		{0x12345678, "foofoofoofoofoofoofoofoo", 1298251034},
	}

	for _, test := range tests {
		hash := MurmurHash3_x86_32([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func TestMurmurHash3_x86_128(t *testing.T) {
	var tests = []struct {
		seed uint32
		key string
		hash [4]uint32
	}{
		{0x12345678, "foo", [4]uint32{1918736524,2601350885,2601350885,2601350885}},
		{0x12345678, "foofoo", [4]uint32{3145568264,1744258630,3744689824,3744689824}},
		{0x12345678, "foofoofoofoofoofoofoofoo", [4]uint32{2421145219,3553640520,1602356257,4216251088}},
	}

	for _, test := range tests {
		hash := MurmurHash3_x86_128([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %v, real: %v", test.key, test.seed, test.hash, hash)
		}
	}
}

func TestMurmurHash3_x64_128(t *testing.T) {
	var tests = []struct {
		seed uint64
		key string
		hash [2]uint64
	}{
		{0x12345678, "foo", [2]uint64{15043494824502257628, 1408089765573727336}},
		{0x12345678, "foofoo", [2]uint64{7114225430032014667, 2403719037022842152}},
		{0x12345678, "foofoofoofoofoofoofoofoo", [2]uint64{3158673277875037204, 15287788684307515255}},
	}

	for _, test := range tests {
		hash := MurmurHash3_x64_128([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %v, real: %v", test.key, test.seed, test.hash, hash)
		}
	}
}

func BenchmarkMurmurHash3_x86_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHashAligned2([]byte("foo"), 0x12345678)
	}
}

func BenchmarkMurmurHash3_x86_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHashAligned2([]byte("foo"), 0x12345678)
	}
}

func BenchmarkMurmurHash3_x64_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHashAligned2([]byte("foo"), 0x12345678)
	}
}

