package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./input/input14.txt")
	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	instructions := make([][]int, 0)
	for {
		strLine := make([]string, 0)

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Problem reading the lines.")
			os.Exit(1)
		}

		strLine = strings.Fields(line)

		intLine := make([]int, len(strLine))

		for i := range strLine {
			tempInt, _ := strconv.Atoi(strLine[i])
			intLine[i] = tempInt
		}
		instructions = append(instructions, intLine)
		if err == io.EOF {
			break
		}
	}

	s := make([]int, len(instructions[1]))
	copy(s, instructions[1])

	start := time.Now()
	_, min, max := qsort(s)
	elapsed := time.Since(start)

	fmt.Printf("min: %d, max: %d\n", min, max)
	fmt.Println("Done in", elapsed)

}

func sort(a []int) {
	count := 0

	// Bubble Sort
	// for i := 0; i < len(a); i++ {
	// 	for j := 0; j < len(a)-1; j++ {
	// 		if a[j] > a[j+1] {
	// 			a[j], a[j+1] = a[j+1], a[j]
	// 			count++
	// 		}
	// 	}
	// }

	// Selection Sort
	// b := make([]int, 0, len(a))
	// min := 0
	// index := 0
	// for len(b) < len(a) {

	// 	for _, c := range a {
	// 		if c != 0 {
	// 			min = c
	// 			break
	// 		}
	// 	}

	// 	for i := 0; i < len(a); i++ {
	// 		if a[i] <= min && a[i] != 0 {
	// 			min = a[i]
	// 			index = i
	// 		}
	// 	}
	// 	count++

	// 	b = append(b, min)
	// 	a[index] = 0
	// }
	fmt.Println("Slice was sorted in", count, "operations")
	fmt.Println("min:", a[0], ", max:", a[len(a)-1])
}

// Quicksort
func qsort(a []int) ([]int, int, int) {

	if len(a) < 2 {
		return a, 0, 0
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	source := int64(time.Now().Nanosecond())
	rand.Seed(source)
	pivotIndex := rand.Intn((len(a) - 1))

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	// sort the left part of the slice
	qsort(a[:left])
	// sort the right slice
	qsort(a[left+1:])

	// return the sorted slice
	return a, a[0], a[len(a)-1]
}
