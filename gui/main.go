package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nwah/gotnfsd"
)

func main() {
	r, w, _ := os.Pipe()

	go func() {
		buf := bufio.NewReader(r)
		for {
			line, _, err := buf.ReadLine()
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("LOGGED LINE: " + string(line))
		}
	}()

	gotnfsd.Init(w)
	gotnfsd.Start(".", 16384, true)
}
