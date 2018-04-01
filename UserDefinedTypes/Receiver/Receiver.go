/*
program that sums the content of a user-defined integeres slice, using a Receiver and a method
*/
package main

import (
	"fmt"
)

// SummableSlice int slice
type SummableSlice []int

// (s SummableSlice): receiver
func (s SummableSlice) sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func main() {
	// declaring and initializing a SummableLice variable
	var sumNumbers = SummableSlice{1, 4, 5, 7, 45, 2, 48, 494}

	// call for the method as part of the SummableSlice "class"
	fmt.Printf("%d", sumNumbers.sum())
	fmt.Printf("\n")
}
