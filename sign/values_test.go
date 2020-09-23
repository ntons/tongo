package sign

import (
	"math/rand"
	"testing"
)

func TestValues(t *testing.T) {
	var expected string
	a := []string{"a", "b", "c", "d", "e", "f", "g"}
	for i, v := range a {
		if i > 0 {
			expected += "&"
		}
		expected += v + "=" + v
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	vals := Values{}
	for _, v := range a {
		vals = append(vals, KV{v, v})
	}
	if s := vals.buffer().String(); s != expected {
		t.Fatalf("unexpected buffer result: %s", s)
	}
}
