package main

import (
	"fmt"

	gotnfsd "github.com/fujinetwifi/tnfsd"
)

func main() {
	gotnfsd.StartServer(".")
	fmt.Println(".")
}
