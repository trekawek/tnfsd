package gotnfsd

/*
#cgo CFLAGS: -DUNIX -DENABLE_CHROOT -DNEED_ERRTABLE -DBSD -DGO
#include <stdio.h>
#include <stdlib.h>
#include "tnfsd.h"
*/
import "C"
import (
	"os"
	"unsafe"
)

type InvalidDirError struct{}
type SocketError struct{}

func (e InvalidDirError) Error() string {
	return "invalid directory"
}
func (e SocketError) Error() string {
	return "couldn't open socket"
}

func Init(log *os.File) {
	C.tnfsd_init_logs(C.int(log.Fd()))
	C.tnfsd_init()
}

func Start(rootPath string, port int, read_only bool) error {
	path := C.CString(rootPath)
	defer C.free(unsafe.Pointer(path))

	err := C.tnfsd_start(path, C.int(port), C.bool(read_only))
	if err == -1 {
		return &InvalidDirError{}
	} else if err == -2 {
		return &SocketError{}
	}
	return nil
}

func Stop() {
	C.tnfsd_stop()
}
