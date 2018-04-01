/*
Program that defines and print out the type of the argument passed to the function
using an empty interface, type switch and type assertion
*/

package main

import "fmt"

// empty interfaces are useful to have an unknown type that's determined at runtime
// all types satisfy such interface
// go statically typed: check types during compilation before the program is run

// function that has an empty interface as parameter
func whatIsThis(i interface{}) {

	// Type Switch: i.(type) -> what is the type of i?
	// swtich case in function of the type
	switch i.(type) {

	// case is string
	case string:
		// i.(string): Type Assertion
		fmt.Printf("It's a string: %v\n", i.(string))

		//case is integer
	case int:
		// i.(int): Type Assertion
		fmt.Printf("It's an integer: %v\n", i.(int))

		// case any other type
	default:
		fmt.Printf("Don't know: %v\n", i)
	}
	// %T: type of i
	fmt.Printf("%T\n", i)
}

func main() {
	whatIsThis(120)
}
