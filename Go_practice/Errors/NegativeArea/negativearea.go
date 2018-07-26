package main

import (
	"fmt"
	"math"
)

// create a custom struct for holding error data
type areaError struct {
	radius float64
	err    string
}

func main() {
	area, err := area(-5)
	if err != nil {
		// assert the error type to the new custom error type
		if t, ok := err.(*areaError); ok {
			fmt.Println(t.Error())
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(area)
}

func area(r float64) (float64, error) {
	if r < 0 {
		// initialize a new areaError struct and returns its adress
		return 0, &areaError{r, "negative number"}
	}
	return math.Pi * r * r, nil
}

// implement the error interface by attaching an Error() string method to the error struct
func (a *areaError) Error() string {
	return fmt.Sprintf("radius %f: %s", a.radius, a.err)
}
