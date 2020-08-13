package nonce

import (
	"testing"
)

func TestRandom(t *testing.T) {
	nonce := Random(32)
	if len(nonce) != 32 {
		t.Fatal("bad nonce length: ", len(nonce))
	}
}
