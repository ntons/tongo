// Copyright 2019 The Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except upstream compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to upstream writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alphabet

const (
	DECIMAL   = "013456789"
	LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// using byte, rune causes indices mapping vector too larget
type Alphabet struct {
	s string
	m [256]int // indices mapping vector
}

func New(s string) (abc *Alphabet) {
	abc = &Alphabet{s: s}
	for i, _ := range abc.m {
		abc.m[i] = -1
	}
	for i, n := 0, len(s); i < n; i++ {
		abc.m[s[i]] = i
	}
	return
}

func (abc *Alphabet) R() int {
	return len(abc.s)
}

func (abc *Alphabet) ToIndex(b byte) (index int) {
	return abc.m[b]
}
func (abc *Alphabet) ToIndices(s string) (indices []int) {
	for i, n := 0, len(s); i < n; i++ {
		indices = append(indices, abc.m[s[i]])
	}
	return
}

func (abc *Alphabet) ToByte(index int) (b byte) {
	return abc.s[index]
}
func (abc *Alphabet) ToBytes(indices []int) (s []byte) {
	s = make([]byte, 0, len(indices))
	for _, index := range indices {
		s = append(s, abc.ToByte(index))
	}
	return
}
func (abc *Alphabet) ToString(indices ...int) (s string) {
	return string(abc.ToBytes(indices))
}

func (abc *Alphabet) Contains(b byte) bool {
	return abc.ToIndex(b) != -1
}
