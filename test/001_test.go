package main

import (
	"fmt"
	"testing"
	. "../../go-murmurhash"
)

func TestMurmurHash1(t *testing.T) {
	var tests = []struct {
		seed uint32
		key string
		hash uint32
	}{
		{0x12345678, "foo", 2077645772},
		{0x12345678, "foofoo", 2304193616},
		{0x12345678, "foofoofoofoofoofoofoofoo", 801342995},
	}

	for _, test := range tests {
		hash := MurmurHash1([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func TestMurmurHash1Aligned(t *testing.T) {
	var tests = []struct {
		seed uint32
		key string
		hash uint32
	}{
		{0x12345678, "foo", 2077645772},
		{0x12345678, "foofoo", 2304193616},
		{0x12345678, "foofoofoofoofoofoofoofoo", 801342995},
	}

	for _, test := range tests {
		hash := MurmurHash1([]byte(test.key), test.seed)
		if hash != test.hash {
			t.Errorf("key: '%s', seed: %d, expect: %d, real: %d", test.key, test.seed, test.hash, hash)
		}
	}
}

func BenchmarkMurmurHash1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash1([]byte("foo"), 0x12345678)
	}
}

func BenchmarkMurmurHash1Aligned(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash1Aligned([]byte("foo"), 0x12345678)
	}
}

func ExampleMurmurHash1() {
	fmt.Printf("%d\n", MurmurHash1([]byte("foo"), 0x12345678))
	// output:
	// 2077645772
}

func ExampleMurmurHash1Aligned() {
	fmt.Printf("%d\n", MurmurHash1Aligned([]byte("foo"), 0x12345678))
	// output:
	// 2077645772
}

