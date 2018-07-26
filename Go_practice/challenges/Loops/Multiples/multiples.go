package main

import "fmt"

func main() {
	n := 2

	for i := 0; i < 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i+1, (i+1)*n)
	}
}
