package main

import (
	"fmt"
	"os"

	"github.com/cyphar/parcel/discovery"
)

func main() {
	for _, arg := range os.Args[1:] {
		name, err := discovery.Resolve(arg)
		fmt.Printf("%s -> (%s, %v)\n", arg, name, err)
	}
}