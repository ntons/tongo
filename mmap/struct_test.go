package mmap

import (
	"fmt"
	"os"
	"testing"
	"unsafe"
)

func TestStruct(t *testing.T) {
	filename := "struct.dat"
	if false {
		os.Remove(filename)
		p, err := Struct(filename, int(unsafe.Sizeof(node{})))
		if err != nil {
			t.Fatal(err)
		}
		st := (*node)(unsafe.Pointer(p))
		(*st).a = 1024
		(*st).b = 2048
		if err := CloseStruct(p); err != nil {
			t.Fatal(err)
		}
	}
	if true {
		p, err := Struct(filename, int(unsafe.Sizeof(node{})))
		if err != nil {
			t.Fatal(err)
		}
		st := (*node)(unsafe.Pointer(p))
		fmt.Println((*st).a, (*st).b)
		(*st).a += 1024
		if err := CloseStruct(p); err != nil {
			t.Fatal(err)
		}
	}
}
