package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input01.txt")
	if err != nil {
		fmt.Printf("Error openning the file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var line string
	tempArr := make([][]int32, 0)

	for {
		stringSlice := make([]string, 0)
		intSlice := make([]int32, 0)

		line, err = reader.ReadString('\n')

		stringSlice = strings.Fields(line)

		var number int

		for _, strNum := range stringSlice {
			number, _ = strconv.Atoi(strNum)
			intSlice = append(intSlice, int32(number))
		}

		tempArr = append(tempArr, intSlice)
		if err == io.EOF {
			break
		}

	}
	caseNum := tempArr[0][0]

	// removing the first element of the slice (test case number)
	tempArr = append(tempArr[:0], tempArr[1:]...)

	intArr := make([][]int32, 0, caseNum)

	for i := 0; i < len(tempArr); i++ {
		if i%2 == 1 {
			intArr = append(intArr, tempArr[i])
		}
	}

	for _, arr := range intArr {
		fmt.Println(minimumBribes(arr))
	}
}

func minimumBribes(q []int32) string {
	var swapNum int
	var swapStr string
	var tempInt int

	for i := 0; i < len(q); i++ {
		tempInt = int(q[i])
		if tempInt-(i+1) > 2 {
			return "Too chaotic"
		}
		for j := int(math.Max(0, float64(q[i]-2))); j < i; j++ {
			if q[j] > q[i] {
				swapNum++
			}
		}
	}
	// for i := 0; i < len(q)-1; i++ {
	// 	for j := i + 1; j < len(q); j++ {
	// 		if q[j] < q[i] {
	// 			q[i], q[j] = q[j], q[i]
	// 			fmt.Println(q)
	// 		}
	// 	}
	// }
	swapStr = strconv.Itoa(swapNum)
	return swapStr
}
