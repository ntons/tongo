package strutil

import (
	"bytes"
	"reflect"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	b := []byte("Hello World")
	s := String(b)
	if s != "Hello World" {
		t.Fatalf("content mismatch")
	}
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	if bh.Data != sh.Data {
		t.Fatalf("data mismatch: %v,%v", bh.Data, sh.Data)
	}
	if bh.Len != sh.Len {
		t.Fatalf("len mismatch: %v,%v", bh.Len, sh.Len)
	}
}

func TestBytes(t *testing.T) {
	s := "Hello World"
	b := Bytes(s)
	if !bytes.Equal(b, []byte("Hello World")) {
		t.Fatalf("content mismatch")
	}
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	if bh.Data != sh.Data {
		t.Fatalf("data mismatch: %v,%v", bh.Data, sh.Data)
	}
	if bh.Len != sh.Len {
		t.Fatalf("len mismatch: %v,%v", bh.Len, sh.Len)
	}
}
