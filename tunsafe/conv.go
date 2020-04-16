package tunsafe

import (
	"reflect"
	"unsafe"
)

func String(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}

func Bytes(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return
}
