package main

import (
	"fmt"
)

type person struct {
	age int
}

func main() {
	p := person{}
	p1 := p.NewPerson(17)
	p1.amIOld()
	p1 = p1.yearPasses()
	p1 = p1.yearPasses()
	p1.amIOld()
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if initialAge < 0 {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
	} else {
		p.age = initialAge
	}
	return p
}

func (p person) amIOld() {
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}

}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++
	return p
}
