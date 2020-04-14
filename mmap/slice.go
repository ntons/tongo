// 注意：一旦超了Cap重新分配内存就完蛋了！

package mmap

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

const SliceHeaderSize = unsafe.Sizeof(reflect.SliceHeader{})

func Slice(
	path string, elemSize, elemCap int) (arbitrary *[]struct{}, err error) {
	var f *os.File
	if f, err = os.OpenFile(
		path, syscall.O_RDWR|syscall.O_CREAT, 0666); err != nil {
		err = errors.New(fmt.Sprintf("OpenFile: %v", err))
		return
	}
	var fi os.FileInfo
	if fi, err = f.Stat(); err != nil {
		err = errors.New(fmt.Sprintf("Stat: %v", err))
		return
	}
	var newFile = fi.Size() == 0

	var size = MmapHeaderSize + SliceHeaderSize + uintptr(elemSize*elemCap)
	if err = f.Truncate(int64(size)); err != nil {
		err = errors.New(fmt.Sprintf("Truncate: %v", err))
		return
	}
	var p uintptr
	if p, err = mmap(0, uintptr(size),
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_SHARED, int(f.Fd()), 0); err != nil {
		err = errors.New(fmt.Sprintf("Mmap: %v", err))
		return
	}
	mh := (*MmapHeader)(unsafe.Pointer(p))
	sh := (*reflect.SliceHeader)(unsafe.Pointer(p + MmapHeaderSize))
	if newFile {
		copy(mh.Tag[:], MmapTag)
		sh.Len = 0
	} else {
		if !bytes.Equal(mh.Tag[:], []byte(MmapTag)) {
			err = errors.New("InvalidTag")
			return
		}
	}
	mh.Len = int(size)
	sh.Data = p + MmapHeaderSize + SliceHeaderSize
	sh.Cap = elemCap
	if sh.Len > sh.Cap {
		sh.Len = sh.Cap
	}
	arbitrary = (*[]struct{})(unsafe.Pointer(sh))
	return
}

func CloseSlice(arbitrary *[]struct{}) (err error) {
	p := (uintptr)(unsafe.Pointer(arbitrary)) - MmapHeaderSize
	mh := (*MmapHeader)(unsafe.Pointer(p))
	if !bytes.Equal(mh.Tag[:], []byte(MmapTag)) {
		panic("InvalidTag")
	}
	return munmap(p, uintptr(mh.Len))
}
