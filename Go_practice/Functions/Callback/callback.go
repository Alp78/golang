package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	s := "I am not thinking of anything else than indexes."

	// IndexFunc is a callback function
	// it has a function as parameter
	// which is called inside itself
	c := strings.IndexFunc(s, unicode.IsSpace)

	fmt.Println(c)

}
