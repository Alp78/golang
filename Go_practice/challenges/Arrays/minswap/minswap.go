package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/input00.txt")
	if err != nil {
		fmt.Println("Error openning the file", err.Error())
	}

	reader := bufio.NewReader(file)
	var tempInt int
	var intArr []int32
	for i := 0; i < 2; i++ {
		line, _ := reader.ReadString('\n')
		if i == 0 {
			continue
		}

		strSlice := strings.Fields(line)
		for _, s := range strSlice {
			tempInt, _ = strconv.Atoi(s)
			intArr = append(intArr, int32(tempInt))
		}
	}
	fmt.Println(intArr)
	fmt.Println(minimumSwaps(intArr))
}

func minimumSwaps(arr []int32) int32 {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
				count++
			}
		}
	}
	fmt.Println(arr)
	return int32(count)
}
