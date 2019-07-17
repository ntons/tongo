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
	"github.com/ntons/tons-go/alphabet"
)

// Node of Trie
type node struct {
	size  int
	value interface{}
	next  []*node
}

func newNode(R int) *node {
	return &node{next: make([]*node, R)}
}

func (x *node) getSize() int {
	if x == nil {
		return 0
	} else {
		return x.size
	}
}
func (x *node) getValue() interface{} {
	if x == nil {
		return nil
	} else {
		return x.value
	}
}

// Trie Searching Tree
type Trie struct {
	abc  *alphabet.Alphabet
	root *node
}

func New(abc *alphabet.Alphabet) *Trie {
	return &Trie{abc: abc}
}

func (t *Trie) Size() int {
	return t.root.getSize()
}

func (t *Trie) Get(s string) (v interface{}) {
	return t.get(t.root, s).getValue()
}
func (t *Trie) get(xx *node, s string) (x *node) {
	if x = xx; x == nil {
		return
	}
	if len(s) == 0 {
		return
	}
	index := t.abc.ToIndex(s[0])
	return t.get(x.next[index], s[1:])
}

func (t *Trie) Put(s string, v interface{}) {
	t.root = t.put(t.root, s, v)
}
func (t *Trie) put(xx *node, s string, v interface{}) (x *node) {
	if x = xx; x == nil {
		x = newNode(t.abc.R())
	}
	if len(s) == 0 {
		if x.value == nil {
			x.size++
		}
		x.value = v
		return
	}
	index := t.abc.ToIndex(s[0])
	size := x.next[index].getSize()
	x.next[index] = t.put(x.next[index], s[1:], v)
	if x.next[index].getSize() != size {
		x.size += x.next[index].getSize() - size
	}
	return
}

func (t *Trie) Del(s string) {
	t.root = t.del(t.root, s)
}
func (t *Trie) del(xx *node, s string) (x *node) {
	if x = xx; x == nil {
		return nil
	}
	if len(s) == 0 {
		if x.value != nil {
			x.size--
		}
		x.value = nil
	} else {
		index := t.abc.ToIndex(s[0])
		size := x.next[index].getSize()
		x.next[index] = t.del(x.next[index], s[1:])
		if x.next[index].getSize() != size {
			x.size += x.next[index].getSize() - size
		}
	}
	if x.size == 0 {
		return nil
	}
	return x
}

func (t *Trie) Keys() (a []string) {
	return t.KeysWithPrefix("")
}
func (t *Trie) KeysWithPrefix(s string) (a []string) {
	return t.collect(t.get(t.root, s), s)
}
func (t *Trie) collect(x *node, s string) (a []string) {
	if x == nil {
		return
	}
	if x.value != nil {
		a = append(a, s)
	}
	for i, _ := range x.next {
		a = append(a, t.collect(x.next[i], s+t.abc.ToString(i))...)
	}
	return
}
