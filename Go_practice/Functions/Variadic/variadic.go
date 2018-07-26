package main

import "fmt"

func main() {
	s := make([]int, 0)

	// append() is a variadic function
	// func append(slice []Type, elems ...Type) []Type
	s = append(s, 5, 6, 7, 8)
	fmt.Println(s)

	fmt.Println(findInt(s, 3, 7, 8))

	// can pass a slice to a variadic parameter
	// arg must be suffixed by ...
	fmt.Println(findInt(s, []int{6, 7}...))

	welcome := []string{"hello", "world"}
	// what is passed is not the actual slice, but the string inside of it
	fmt.Println(change(welcome...))
	// thus the original slice is not modified
	fmt.Println(welcome)
}

// variadic function that returns true if all the int args are in the slice
// variadic parameter must be at the end of the parameters' list
func findInt(arr []int, ints ...int) bool {
	var t bool
	var count int

	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(arr); j++ {
			if ints[i] == arr[j] {
				count++
				break
			}
		}
	}
	if count == len(ints) {
		t = true
	} else {
		t = false
	}
	return t
}

func change(s ...string) []string {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
	return s
}
