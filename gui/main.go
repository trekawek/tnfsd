package main

import (
	"github.com/nwah/gotnfsd"
)

func main() {
	gotnfsd.Init(nil)
	gotnfsd.Start(".", 16384, true)
}
