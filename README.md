# go-murmurhash

go-murmurhash is Go implementation of murmurhash algorithm.

## Installation
```sh
go get github.com/rryqszq4/go-murmurhash
```
## Usage
```go
package main

import (
   "fmt"
   . "github.com/rryqszq4/go-murmurhash"
)

func main() {
   var foo string = "foo"
   var seed uint32 = 0x12345678
   
   var key = []byte(foo)
   
   fmt.Printf("%d\n", MurmurHash64A(key, uint64(seed)))
}
```

## API
```
func MurmurHash1 (key []byte, seed uint32) (hash uint32)
func MurmurHash1Aligned(key []byte, seed uint32) (hash uint32)

func MurmurHash2 (key []byte, seed uint32) (hash uint32)
func MurmurHash64A (key []byte, seed uint64) (hash uint64)
func MurmurHash64B(key []byte, seed uint64) (hash uint64)
func MurmurHash2A(key []byte, seed uint32) (hash uint32)
func MurmurHashNeutral2(key []byte, seed uint32) (hash uint32)
func MurmurHashAligned2(key []byte, seed uint32) (hash uint32)

func MurmurHash3_x86_32(key []byte, seed uint32) (hash uint32)
func MurmurHash3_x86_128(key []byte, seed uint32) (hash [4]uint32)
func MurmurHash3_x64_128(key []byte, seed uint64) (hash [2]uint64)
```
## Test
cd test; go test -v
```
=== RUN   TestMurmurHash1
--- PASS: TestMurmurHash1 (0.00s)
=== RUN   TestMurmurHash1Aligned
--- PASS: TestMurmurHash1Aligned (0.00s)
=== RUN   TestMurmurHash2
--- PASS: TestMurmurHash2 (0.00s)
=== RUN   TestMurmurHash64A
--- PASS: TestMurmurHash64A (0.00s)
=== RUN   TestMurmurHash64B
--- PASS: TestMurmurHash64B (0.00s)
=== RUN   TestMurmurHash2A
--- PASS: TestMurmurHash2A (0.00s)
=== RUN   TestMurmurHashNeutral2
--- PASS: TestMurmurHashNeutral2 (0.00s)
=== RUN   TestMurmurHashAligned2
--- PASS: TestMurmurHashAligned2 (0.00s)
=== RUN   TestMurmurHash3_x86_32
--- PASS: TestMurmurHash3_x86_32 (0.00s)
=== RUN   TestMurmurHash3_x86_128
--- PASS: TestMurmurHash3_x86_128 (0.00s)
=== RUN   TestMurmurHash3_x64_128
--- PASS: TestMurmurHash3_x64_128 (0.00s)
=== RUN   ExampleMurmurHash1
--- PASS: ExampleMurmurHash1 (0.00s)
=== RUN   ExampleMurmurHash1Aligned
--- PASS: ExampleMurmurHash1Aligned (0.00s)
PASS
```


go test -bench=. -benchmem
```
goos: darwin
goarch: amd64
BenchmarkMurmurHash1-8           	300000000	         5.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHash1Aligned-8    	200000000	         6.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHash2-8           	300000000	         5.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHash64A-8         	300000000	         5.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHash64B-8         	200000000	         7.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHash2A-8          	200000000	         6.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHashNeutral2-8    	200000000	         6.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHashAligned2-8    	300000000	         5.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHash3_x86_32-8    	300000000	         5.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHash3_x86_128-8   	300000000	         5.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkMurmurHash3_x64_128-8   	300000000	         5.64 ns/op	       0 B/op	       0 allocs/op
PASS
```

## License
```
BSD 2-Clause License

Copyright (c) 2019, rryqszq4 <rryqszq@gmail.com>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
