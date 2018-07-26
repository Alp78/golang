package main

import (
	"fmt"
	"runtime/debug"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from: ", r)

		// as the recover() eliminates the stack trace, we can print it with debug pack:
		debug.PrintStack()
	}
}

func a() {
	// deferred call to recovery()
	defer recovery()
	s := []int{1, 2, 3, 4}
	// generates a panic
	fmt.Println(s[4])
}

func main() {
	a()

	fmt.Println("Returned to caller main() after panic and recovery in a().")
}
