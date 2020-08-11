package nocopy

import (
	"reflect"
	"unsafe"
)

func BytesToString(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return
}
