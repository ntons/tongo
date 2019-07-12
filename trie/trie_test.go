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

package trie

import (
	"math/rand"
	"testing"

	"github.com/ntons/tons-go/alphabet"
)

const ABC = alphabet.LOWERCASE + alphabet.DECIMAL + "-_"

func RandString(maxLen int) string {
	n := rand.Intn(maxLen + 1)
	b := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		b = append(b, ABC[rand.Intn(len(ABC))])
	}
	return string(b)
}

func TestTrie(tt *testing.T) {
	t := New(ABC)
	m := make(map[string]interface{})

	for i := 0; i < 1000000; i++ {
		s := RandString(16)
		switch rand.Intn(3) {
		case 0: // put
			t.Put(s, s)
			m[s] = s
		case 1: // get
			// do nothing
		case 2: // del
			t.Del(s)
			delete(m, s)
		}
		if v1, v2 := t.Size(), len(m); v1 != v2 {
			tt.Fatalf("length mismatch: %q, %v, %v", s, v1, v2)
			return
		}
		if v1, v2 := t.Get(s), m[s]; v1 != v2 {
			tt.Fatalf("value mismatch: %q, %v, %v", s, v1, v2)
			return
		}
	}
}
