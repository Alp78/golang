package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input08.txt")
	if err != nil {
		fmt.Printf("Error openning the file: %v\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(file)

	var line string
	intArr := make([][]int32, 0, 2)

	for {
		stringSlice := make([]string, 0)
		intSlice := make([]int32, 0)

		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
		}

		stringSlice = strings.Fields(line)

		var number int

		for _, strNum := range stringSlice {
			number, _ = strconv.Atoi(strNum)
			intSlice = append(intSlice, int32(number))
		}

		intArr = append(intArr, intSlice)
		if err == io.EOF {
			break
		}

	}
	fmt.Printf("Length of array: %v\n", intArr[0][0])
	fmt.Printf("Number of rotations: %v\n", intArr[0][1])

	fmt.Println(intArr)
	// // number of left rotation
	// var d int32
	// // array of numbers to rotate left
	// a := make([]int32, len(intArr[1]))

	// d = intArr[0][1]
	// a = intArr[1]
	// fmt.Println(d)
	// fmt.Println(a)

}
