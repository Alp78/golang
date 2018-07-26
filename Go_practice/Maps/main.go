package main

import (
	"fmt"
)

func main() {

	// map declaration 1
	/*
		colors := map[string]string{
			"red":   "#ff0000",
			"green": "#008000",
			"blue":  "#0000FF",
		}
	*/

	// map declaration 2
	// var colors map[string]string

	// map declaration 3
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#008000"
	colors["blue"] = "#0000FF"

	fmt.Println(colors["red"])

	// delete a key-value pair from a map
	// delete(colors, "red")
	// fmt.Println(colors)

	printMap(colors)

	// check is a Key exists in the map
	fmt.Println(colors)

	_, exists := colors["yellow"]
	if exists {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// idiomatic if statement
	if _, exists := colors["yellow"]; exists {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}

// iterate over key-value pairs
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for '%s' is '%s'\n", color, hex)
	}
}
