package main

import "fmt"

type car struct {
	model string
	year  string
}

type employee struct {
	name string
	*car
}

func main() {

	car1 := car{
		model: "ferrari",
		year:  "1973",
	}

	andre := employee{
		name: "Andre",
		car:  &car1,
	}

	andre.printCar()
}

func (c *car) printCar() {
	fmt.Printf("%v from %v", c.model, c.year)
}
