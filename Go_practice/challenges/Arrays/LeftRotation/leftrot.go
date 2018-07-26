package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	file, err := os.Open("input00.txt")
	if err != nil {
		fmt.Printf("Error openning the file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var line string
	intArr := make([][]int32, 0, 2)

	for {
		stringSlice := make([]string, 0)
		intSlice := make([]int32, 0)

		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
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
	// number of left rotation
	var d int32
	// array of numbers to rotate left
	var a []int32

	d = intArr[0][1]
	a = intArr[1]

	var result []int32

	start := time.Now()
	result = rotLeft(a, d)
	elapsed := time.Since(start)
	fmt.Println(result)

	fmt.Printf("Function took %v\n", elapsed)
	// fmt.Printf("Size of array: %v\n", intArr[0][0])
	// fmt.Printf("Number of rotations: %v\n", intArr[0][1])
	// var str string
	// str = arrayToString(result, " ")
	// ioutil.WriteFile("output08.txt", []byte(str), 0666)

}

// func rotLeft(a []int32, d int32) []int32 {
// 	n := len(a)
// 	rotations := int(d)
// 	// execute rotation d times
// 	for j := 0; j < rotations; j++ {
// 		// inversion of first and last elements
// 		a[n-1], a[0] = a[0], a[n-1]
// 		// switch of all elements left but the last
// 		for i := 0; i < len(a)-2; i++ {
// 			a[i+1], a[i] = a[i], a[i+1]
// 		}
// 	}
// 	return a
// }

// func rotLeft(a []int32, d int32) []int32 {
// 	n := len(a)
// 	rot := int(d)
// 	var temp int32
// 	var j, k int

// 	// execute rotation d times
// 	for i := 0; i < len(a)-1; i++ {
// 		temp = a[i]
// 		j = i
// 		for {
// 			k = j + rot
// 			if k >= n {
// 				k = k - n
// 			}
// 			if k == i {
// 				break
// 			}
// 			a[j] = a[k]
// 			j = k
// 		}
// 		a[j] = temp
// 	}
// 	return a
// }

func rotLeft(a []int32, d int32) []int32 {

	sLeft := a[:d]
	fmt.Println(sLeft)
	for i, j := 0, len(sLeft)-1; i < j; i, j = i+1, j-1 {
		sLeft[i], sLeft[j] = sLeft[j], sLeft[i]
	}

	sRight := a[d:]
	fmt.Println(sRight)
	for i, j := 0, len(sRight)-1; i < j; i, j = i+1, j-1 {
		sRight[i], sRight[j] = sRight[j], sRight[i]
	}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func arrayToString(a []int32, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
