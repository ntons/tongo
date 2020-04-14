// mmap对象存储

package mmap

import (
	"syscall"
	"unsafe"
)

func mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error) {
	xaddr, _, err = syscall.Syscall6(syscall.SYS_MMAP, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flags), uintptr(fd), uintptr(offset))
	if err == syscall.Errno(0) {
		err = nil
	}
	return
}

func munmap(addr uintptr, length uintptr) (err error) {
	_, _, err = syscall.Syscall(syscall.SYS_MUNMAP, uintptr(addr), uintptr(length), 0)
	if err == syscall.Errno(0) {
		err = nil
	}
	return
}

type MmapHeader struct {
	Tag [4]byte
	Len int
}

const MmapTag = "MMAP"
const MmapHeaderSize = unsafe.Sizeof(MmapHeader{})
