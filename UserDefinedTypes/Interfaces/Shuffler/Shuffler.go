/*
program that shuffles and prints a slice of integers and the a slice of string
using an interface defining 2 methods used to shuffle the elements
-> demonstrate that interfaces are type agnostic

// no class or inheritance in go, but types, and methods attached to them
// interface: define behaviors (methods) without defining types themselves
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// any type that satisfies this interface has its methods attached to it
type shuffler interface {
	Len() int
	Swap(i, j int)
}

// functions that has "shuffler" interface type as parameter
// any argument that satisfies "shuffler" interface can be passed into this function
func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {

		// initializing the Seed based on the current time expressed in nano seconds
		// to generate each time a new random int with rand.Intn
		rand.Seed(time.Now().UnixNano())

		// generating a random integer between [0; n-1]
		j := rand.Intn(s.Len())

		// printing j to show its random value
		fmt.Printf("%d ", j)

		// i, j: used as indexes inside swap function
		// to swap the location of their respective element in the slice
		s.Swap(i, j)
	}
}

//intSlice integer slice: satisfies the "shuffler" interface
// as it has its two methods Len() and Swap() attached
type intSlice []int

// defining the method Len() attached to "intSlice" type
// returns the length of the slice
func (is intSlice) Len() int {
	return len(is)
}

// defining the method Swap() attached to "intSlice" type
//
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// same operation with a string slice -> interface is type agnostic
// it only look at the methods

//intSlice integer slice: satisfies the "shuffler" interface
// as it has its two methods Len() and Swap() attached
type stringSlice []string

// defining the method Len() attached to "intSlice" type
// returns the length of the slice
func (is stringSlice) Len() int {
	return len(is)
}

// defining the method Swap() attached to "intSlice" type
//
func (is stringSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	fmt.Printf("\n")

	slice1 := intSlice{1567, 22, 31, 404, 52, 63456, 71, 86, 92, 101}
	shuffle(slice1)

	// %v: the value in a default format
	fmt.Printf("%v\n", slice1)

	fmt.Printf("\n\n")

	slice2 := stringSlice{"red", "blue", "green", "pink", "orange", "yellow"}
	shuffle(slice2)
	fmt.Printf("%v\n", slice2)

	fmt.Printf("\n\n")
}
