package main

import (
	"fmt"
	"reflect"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {

	t := triangle{
		height: 20.5,
		base:   4.3,
	}

	s := square{
		sideLength: 2.7,
	}

	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Printf("The area of %v is: %v\n\n", reflect.TypeOf(s).Name(), s.getArea())
}
