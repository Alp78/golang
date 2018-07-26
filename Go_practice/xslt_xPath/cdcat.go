package main

import (
	"fmt"
	"os"
	"strconv"

	xmlpath "gopkg.in/xmlpath.v2"
)

func main() {
	file, err := os.Open("cdcat.xml")
	if err != nil {
		fmt.Println("Error openning the XML file: ", err.Error())
	}

	root, err := xmlpath.Parse(file)
	if err != nil {
		fmt.Println("Error parsing the XML file: ", err.Error())
	}

	for i := 1; i <= 2; i++ {
		pathStr := "/catalog/cd[" + strconv.Itoa(i) + "]/artist"
		path := xmlpath.MustCompile(pathStr)
		if value, ok := path.String(root); ok {
			fmt.Printf("Found: %v\n", value)
		}
	}
}
