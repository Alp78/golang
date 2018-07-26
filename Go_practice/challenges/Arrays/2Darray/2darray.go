package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input03.txt")
	if err != nil {
		fmt.Printf("There was an error openning the file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	intArr := make([][]int, 0, 6)
	var line string

	// scans each line of the input, converts strings into int
	// store the case into a slice of slice (2d array)

	for scanner.Scan() {
		stringLineSlice := make([]string, 6)
		intLineSlice := make([]int, 6)
		line = scanner.Text()
		if len(line) > 0 {
			stringLineSlice = strings.Fields(line)
		}

		for i, strNum := range stringLineSlice {
			intLineSlice[i], _ = strconv.Atoi(strNum)
		}

		intArr = append(intArr, intLineSlice)
	}

	var sumHourGlass int
	hgArr := make([]int, 0, 16)
	// sums the 16 Hour glasses in a slice
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sumHourGlass = intArr[i][j] + intArr[i][j+1] + intArr[i][j+2] +
				intArr[i+1][j+1] + intArr[i+2][j] + intArr[i+2][j+1] + intArr[i+2][j+2]
			hgArr = append(hgArr, sumHourGlass)

		}

	}

	// sorting the Hour Glass Sums slice
	for i := 0; i < len(hgArr); i++ {
		for j := i + 1; j < len(hgArr); j++ {
			if hgArr[j] < hgArr[i] {
				hgArr[i], hgArr[j] = hgArr[j], hgArr[i]
			}
		}
	}
	// return hte last element which is the highest sum
	fmt.Println(hgArr[len(hgArr)-1])
}

// for HackerRank
func hourglassSum(intArr [][]int32) int32 {
	var sumHourGlass int32
	hgArr := make([]int32, 0, 16)
	// sums the 16 Hour glasses in a slice
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sumHourGlass = intArr[i][j] + intArr[i][j+1] + intArr[i][j+2] +
				intArr[i+1][j+1] + intArr[i+2][j] + intArr[i+2][j+1] + intArr[i+2][j+2]
			hgArr = append(hgArr, sumHourGlass)
		}
	}

	// sorting the Hour Glass Sums slice
	for i := 0; i < len(hgArr); i++ {
		for j := i + 1; j < len(hgArr); j++ {
			if hgArr[j] < hgArr[i] {
				hgArr[i], hgArr[j] = hgArr[j], hgArr[i]
			}
		}
	}

	// return the last element which is the highest sum
	return hgArr[len(hgArr)-1]
}
