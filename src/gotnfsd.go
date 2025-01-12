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
	C.tnfsd_init()

	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err.Error())
	}

	C.tnfsd_init_logs(C.int(w.Fd()))

	// w.Write([]byte("hi\n"))
	// w2 := os.NewFile(w.Fd(), "via fd")
	// w2.Write([]byte("hi via fd\n"))

	go func() {
		buf := bufio.NewReader(r)
		for {
			line, _, err := buf.ReadLine()
			if err != nil {
				log.Println("Error from buf.ReadLine()")
				log.Fatal(err.Error())
			}
			fmt.Println(string(line))
		}
	}()

	// go func() {
	// 	scanner := bufio.NewScanner(r)
	// 	scanner.Split(bufio.ScanBytes)
	// 	for scanner.Scan() {
	// 		fmt.Print(scanner.Text())
	// 	}
	// }()

	// var buf bytes.Buffer

	// go func() {
	// 	io.Copy(&buf, w)
	// }()

	// go func() {
	// 	for {
	// 		ioutil.ReadAll(r)
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()
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
