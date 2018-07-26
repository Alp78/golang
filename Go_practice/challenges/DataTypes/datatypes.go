package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		i  uint64 = 4
		d         = 4.0
		s         = "HackerRank "
		i2 uint64
		d2 float64
		s2 string
	)

	file, _ := os.Open("input00.txt")
	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	i2, _ = strconv.ParseUint(input[0], 10, 64)
	d2, _ = strconv.ParseFloat(input[1], 64)
	s2 = input[2]

	// Print the sum of both integer variables on a new line.
	fmt.Printf("%d\n", i+i2)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+d2)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Printf("%s", s+s2)
}
