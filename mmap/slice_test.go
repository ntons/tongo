package mmap

import (
	"fmt"
	"os"
	"testing"
	"unsafe"
)

type node struct {
	a uint64
	b int64
}

func TestSlice(t *testing.T) {
	filename := "slice.dat"
	if false {
		os.Remove(filename)
		p, err := Slice(filename, int(unsafe.Sizeof(node{})), 10)
		if err != nil {
			t.Fatal(err)
		}
		vec := (*[]node)(unsafe.Pointer(p))
		fmt.Println(len(*vec), cap(*vec))
		(*vec) = append(*vec, node{
			0xAAAAAAAAAAAAAAAA,
			0x1111111111111111,
		})
		fmt.Println(len(*vec), cap(*vec))
		if err := CloseSlice(p); err != nil {
			t.Fatal(err)
		}
	}
	if true {
		p, err := Slice(filename, int(unsafe.Sizeof(node{})), 10)
		if err != nil {
			t.Fatal(err)
		}
		vec := (*[]node)(unsafe.Pointer(p))
		fmt.Println(len(*vec), cap(*vec))
		(*vec) = append(*vec, node{
			0x123456789ABCDEF0,
			0x123456789ABCDEF0,
		})
		fmt.Println(len(*vec), cap(*vec))
		if err := CloseSlice(p); err != nil {
			t.Fatal(err)
		}
	}
}
