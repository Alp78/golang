package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("There was an error opening the file: %s", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
