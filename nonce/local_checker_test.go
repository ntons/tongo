package nonce

import (
	"testing"
	"time"
)

func TestLocalChecker(t *testing.T) {
	var checker Checker = NewLocalChecker(time.Second)

	if !checker.Check("1") {
		t.Fatal("1")
	}
	if !checker.Check("2") {
		t.Fatal("2")
	}
	if !checker.Check("1024") {
		t.Fatal("1024")
	}
	if checker.Check("2") {
		t.Fatal("2")
	}

	time.Sleep(time.Second / 2)

	if checker.Check("1024") {
		t.Fatal("1024")
	}
	if !checker.Check("512") {
		t.Fatal("512")
	}

	time.Sleep(time.Second / 2)

	if !checker.Check("1024") {
		t.Fatal("1024")
	}
	if checker.Check("512") {
		t.Fatal("512")
	}
}
