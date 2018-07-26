package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("test.txt")
	// if err != nil {
	// 	// assign a custom error message to err with Errorf()
	// 	err = fmt.Errorf("Error openning: %v", f)
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(f.Name())
	// }

	// returning only the path of the failed file from PathError struct
	// if err != nil {
	// 	fmt.Println(err.(*os.PathError).Path)
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(f.Name())
	// }

	// using type assertion to return the path if there is an error
	// as an interace can have different underlying types
	// type assertion is the way to find out which it is.
	// once the underlying type asserted, we can access to their properties (fields, methods, ...)
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err.Path)
	} else {
		fmt.Println(f.Name())
	}

}
