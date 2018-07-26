package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type intMapSlice []map[int][]int

func main() {
	// opens the input.txt file from app folder
	inputFile := "input.txt"
	content, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("There was an error openning the file: %v", err)
		os.Exit(1)
	}
	// closes the file at program exit
	defer content.Close()
	// read the file content
	reader := bufio.NewReader(content)
	var line string
	tempSlice := make([]string, 0)
	// while (infinite for)
	for {
		line, err = reader.ReadString('\n')
		tempSlice = append(tempSlice, line)
		// when ReadString reaches EOF -> exit the loop
		if err == io.EOF {
			break
		}
	}

	//storing the number of test cases (first element of slice)
	testCaseNumber, err := strconv.Atoi(strings.Fields(tempSlice[0])[0])
	fmt.Println(testCaseNumber)
	if err != nil {
		fmt.Printf("There was an error converting to int: %v", err)
	}

	// removing the first element of the slice (test case number)
	tempSlice = append(tempSlice[:0], tempSlice[1:]...)

	//convert the slice of strings into a slice of slice of int
	// slice hosting slices of int to store each converted string slice
	caseIntSlice := make([][]int, len(tempSlice))

	for i := 0; i < len(tempSlice); i++ {
		// make a new empty slice of string each iteration to host temporarily each number in each line
		tempStringSlice := make([]string, len(tempSlice[i]))
		// get rid of new lines \n, spaces \s and carriage return \r
		tempStringSlice = strings.Fields(tempSlice[i])
		// need to make a new slice for each slice in a 2-dimensional array
		caseIntSlice[i] = make([]int, len(tempStringSlice))
		// fill each slice of slice with the converted slices of integers
		for j := 0; j < len(tempStringSlice); j++ {
			// convert string to int
			tempInt, _ := strconv.Atoi(tempStringSlice[j])
			// assign the converted integer
			caseIntSlice[i][j] = tempInt
		}
	}
	// create a slice of map[int]int
	caseCollection := make(intMapSlice, 0)

	// fill the slice with maps hosting each a pair of int slice
	// 0: N, K
	// 1: K-sums
	for i := 0; i < len(caseIntSlice); i = i + 2 {
		tempMap := make(map[int][]int)
		tempMap[0] = caseIntSlice[i]
		tempMap[1] = caseIntSlice[i+1]
		caseCollection = append(caseCollection, tempMap)
	}

	resultArray := make([][]int, testCaseNumber)

	// checks the length of the progression slices
	for i, testCase := range caseCollection {
		n := testCase[0][0]
		k := testCase[0][1]
		switch n {
		case 1:
			resultSlice := make([]int, n)
			resultSlice[0] = testCase[1][0] / k
			resultArray[i] = resultSlice
			continue
		case 2:
			resultSlice := make([]int, n)
			resultSlice[0] = testCase[1][0] / k
			resultSlice[1] = testCase[1][len(testCase[1])-1] / k
			resultArray[i] = resultSlice
			continue
		default:
			resultSlice := make([]int, n)
			first := minIntSlice(testCase[1]) / k
			last := maxIntSlice(testCase[1]) / k
			resultSlice[0] = first
			resultSlice[1] = last
			resultArray[i] = resultSlice
			continue
		}
	}
	fmt.Println(resultArray)
	// case there is only 1 number

	// logic for solving K-sums in one of the cases (maps in collection)

	// first number of the progression is always the min value of k-sums divided by k
	// fmt.Println(caseCollection[2][1])
	fmt.Println(combRep(2, 2))

}

// minIntSlice copies the []int given as arg, then orders its numbers
// finally returns the first element which is the lowest value of the collection
func minIntSlice(s []int) int {
	ts := make([]int, len(s))
	copy(ts, s)
	for i := 0; i < len(ts); i++ {
		for j := i + 1; j < len(ts); j++ {
			if ts[j] <= ts[i] {
				ts[i], ts[j] = ts[j], ts[i]
			}
		}
	}
	return ts[0]
}

func maxIntSlice(s []int) int {
	ts := make([]int, len(s))
	copy(ts, s)
	for i := 0; i < len(ts); i++ {
		for j := i + 1; j < len(ts); j++ {
			if ts[j] <= ts[i] {
				ts[i], ts[j] = ts[j], ts[i]
			}
		}
	}
	return ts[len(ts)-1]
}

// combination with repetitions

/*
func combrep(n int, lst []int) [][]int {
	r := combrep(n, lst[1:])
	for _, x := range combrep(n-1, lst) {
		r = append(r, append(x, lst[0]))
	}
	return r
}
*/

// combRep returns the number of combinations with repetition for n, k
func combRep(n, k int) int {
	d := n + k - 1
	dividend := 1
	for i := 1; i <= d; i++ {
		dividend = i * dividend
	}

	d1 := 1
	for i := 1; i <= n-1; i++ {
		d1 = i * d1
	}

	d2 := 1
	for i := 1; i <= k; i++ {
		d2 = i * d2
	}
	divisor := d1 * d2

	return dividend / divisor
}
