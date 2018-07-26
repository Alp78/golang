package main

import (
	"fmt"
	"reflect"
)

func main() {

	// slice of empty interface: elements can be of any type
	eis := make([]interface{}, 4, 4)
	eis[0] = 23
	eis[1] = 5.46
	eis[2] = "blahblah"
	eis[3] = true

	// type assertion -> t1 gets the ubderlying value of the interface
	t1, ok := eis[0].(int)
	fmt.Println(t1, ok)

	t2, ok := eis[0].(string)
	fmt.Println(t2, ok)

	for _, item := range eis {
		// type switch
		switch t := item.(type) {
		case int:
			fmt.Printf("Type: %T - %v\n", t, t)
		case float64:
			fmt.Printf("Type: %T - %v\n", t, t)
		case string:
			fmt.Printf("Type: %T - %v\n", t, t)
		case bool:
			fmt.Printf("Type: %T - %v\n", t, t)
		}
	}

	// check with ok keyword
	if _, ok := eis[0].(int); ok {
		fmt.Println(eis[0], "is of type", reflect.TypeOf(eis[0]))
	}

}
