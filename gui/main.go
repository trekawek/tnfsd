package main

import (
	"fmt"

	gotnfsd "github.com/fujinetwifi/tnfsd"
)

func main() {
	fmt.Println("Starting server")
	gotnfsd.StartServer(".")
}
