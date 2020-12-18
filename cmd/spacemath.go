package main

import (
	"flag"
	"fmt"

	"github.com/mikerourke/spacemath"
)

func main() {
	pathPtr := flag.String("path", "", "Path of the file to open")
	flag.Parse()

	if err := spacemath.Open(*pathPtr); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
