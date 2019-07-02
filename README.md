# go-murmurhash

go-murmurhash is Go implementation of murmurhash algorithm.

## Installation
```sh
go get github.com/rryqszq4/go-murmurhash
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
