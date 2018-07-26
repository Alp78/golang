package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	emp1 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp1.firstName)
	fmt.Println("Age:", emp1.age)

	emp2 := &Employee{"Alex", "Ricci", 23, 15900}
	fmt.Println("First Name:", (*emp2).firstName)
	// (*emp2) can be shorten to emp2 - Go still knows it's a pointer
	fmt.Println("Age:", emp2.age)

}
