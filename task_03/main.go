package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	f, error := os.Open(fileName)
	if error != nil {
		fmt.Println("An error occours:", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
