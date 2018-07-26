package main

import "fmt"

// defining a new custom type of underlying type struct
type person struct {
	// defining fields
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	// declaration 1 - initialize fields in order of definition
	//alex := person{"Alex", "Anderson"}

	// declaration 2 - with labels
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// declaration 3 - null values and assignment
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex.anderson@blah.com"
	alex.contact.zip = 12300

	// %+v: print all fields with their value
	//fmt.Printf("%+v", alex)

	// %+s: only values (if string)
	// fmt.Printf("%+s", alex)

	// embed a struct in another
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "jim.party@blah.com",
			// multiline declaration must end with comma, even last line
			zip: 15800,
			// multiline declaration must end with comma, even after closing curly bracket
		},
	}
	jim.print()

	// declaration of a pointer with '&' operator
	// -> gives the memory address of the value this variable is pointing at
	// it is not necessary to declare the pointer - if the function has a receiver
	// or parameter as pointer -> go will automatically convert the value into a pointer
	jimPointer := &jim
	jimPointer.updateFirstName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*
// go: "pass by value" language -> copy the data to a new container
// so the receiver won't be affected by the function
func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}
*/

// receiver function with pointer -> affects the original value
// '*' pointer operator -> gives the value this memory address is pointing at
// ATTENTION: '*person' is a type description - not the actual pointer
func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}
