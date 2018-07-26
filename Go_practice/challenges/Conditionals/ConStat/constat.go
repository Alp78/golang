package main

import "fmt"

func main() {
	var N int32 = 8

	switch true {
	case (N%2 == 1) || (N%2 == 0 && N >= 6 && N <= 20):
		fmt.Println("Weird")
	case (N%2 == 0 && N >= 2 && N <= 5) || (N%2 == 0 && N > 20):
		fmt.Println("Not Weird")

	}

}
