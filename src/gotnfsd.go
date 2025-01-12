package gotnfsd

/*
#cgo CFLAGS: -DUNIX -DENABLE_CHROOT -DNEED_ERRTABLE -DBSD -DGO
#include <stdio.h>
#include <stdlib.h>
#include "tnfsd.h"
*/
import "C"
import (
	"bufio"
	"fmt"
	"log"
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

func Init(_ *os.File) {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err.Error())
	}

	C.tnfsd_init_logs(C.int(w.Fd()))
	C.tnfsd_init()

	go func() {
		buf := bufio.NewReader(r)
		for {
			line, _, err := buf.ReadLine()
			if err != nil {
				log.Println("Error from buf.ReadLine()")
				log.Fatal(err.Error())
			}
			fmt.Println("LOGGED LINE: " + string(line))
		}
	}()
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
