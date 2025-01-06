package main

import (
	"fmt"

	gotnfsd "github.com/fujinetwifi/tnfsd"
)

func main() {
<<<<<<< HEAD
	fmt.Println("Starting server")
	gotnfsd.StartServer(".")
=======
	gotnfsd.StartServer(".")
	fmt.Println(".")
>>>>>>> b804bc8 (feat(gui): add Go wrapper)
}
