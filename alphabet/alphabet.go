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
	Decimal   = "0123456789"
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// using byte, rune causes indices mapping vector too larget
type Alphabet struct {
	s string
	m [256]int // indices mapping vector
}

func New(s string) (ab *Alphabet) {
	ab = &Alphabet{s: s}
	for i, _ := range ab.m {
		ab.m[i] = -1
	}
	for i, n := 0, len(s); i < n; i++ {
		ab.m[s[i]] = i
	}
	return
}

func (ab *Alphabet) Len() int {
	return len(ab.s)
}

func (ab *Alphabet) ToIndex(b byte) (index int) {
	return ab.m[b]
}
func (ab *Alphabet) ToIndices(s string) (indices []int) {
	for i, n := 0, len(s); i < n; i++ {
		indices = append(indices, ab.m[s[i]])
	}
	return
}

func (ab *Alphabet) ToByte(index int) (b byte) {
	return ab.s[index]
}
func (ab *Alphabet) ToBytes(indices []int) (s []byte) {
	s = make([]byte, 0, len(indices))
	for _, index := range indices {
		s = append(s, ab.ToByte(index))
	}
	return
}
func (ab *Alphabet) ToString(indices ...int) (s string) {
	return string(ab.ToBytes(indices))
}

func (ab *Alphabet) Contains(b byte) bool {
	return ab.ToIndex(b) != -1
}
