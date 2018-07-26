package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input/input14.txt")
	if err != nil {
		fmt.Printf("Error openning the file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var line string
	str, _, err := reader.ReadLine()
	line = strings.TrimRight(string(str), "\r\n")

	fmt.Println(isValid(line))
}

func isValid(s string) string {
	if len(s) <= 2 {
		return "YES"
	}

	var count int
	charSlice := make([]string, 0)
	tempSlice := make([]string, 0)
	countSlice := make([]int, 0)
	countDist := make([]int, 0)
	numSlice := make([]int, 0)
	// finSlice := make([]int, 0)

	for _, c := range s {
		tempSlice = append(tempSlice, string(c))
	}
	fmt.Println(tempSlice)
	for i, r := range tempSlice {
		count = 1
		for j := i + 1; j < len(tempSlice); j++ {
			if tempSlice[j] == r && tempSlice[j] != "!" {
				count++
				tempSlice[j] = "!"
			}
		}
		if count > 1 || (count == 1 && r != "!") {
			countSlice = append(countSlice, count)
		}
	}

	for _, r := range tempSlice {
		if r != "!" {
			charSlice = append(charSlice, r)
		}
	}
	fmt.Println(countSlice)
	fmt.Println(charSlice)

	for i := 0; i < len(countSlice); i++ {
		count = 1
		for j := i + 1; j < len(countSlice); j++ {
			if countSlice[j] == countSlice[i] && countSlice[i] != 0 {
				count++
				countSlice[j] = 0
			}
		}
		if countSlice[i] != 0 {
			numSlice = append(numSlice, count)
		}
	}

	fmt.Println(countSlice)
	fmt.Println(numSlice)
	if len(numSlice) > 2 {
		return "NO"
	}
	for _, n := range countSlice {
		if n != 0 {
			countDist = append(countDist, n)
		}
	}
	if countDist[0] < countDist[1] {
		countDist[0], countDist[1] = countDist[1], countDist[0]
		numSlice[0], numSlice[1] = numSlice[1], numSlice[0]
	}

	fmt.Println(numSlice)
	fmt.Println(countDist)

	if numSlice[0] > 1 && numSlice[1] > 1 {
		return "NO"
	}
	if (countDist[0] > 2 && numSlice[0] == 1 && numSlice[1] > 1 && countDist[0]-countDist[1] > 1) ||
		(countDist[1] > 2 && numSlice[1] == 1 && numSlice[0] > 1 && countDist[1]-countDist[0] > 1) {
		return "NO"
	}
	return "YES"
}
