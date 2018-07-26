package main

import (
	"bufio"
	"fmt"
	"io"
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
	strArr := make([][]string, 0)

	for {
		stringSlice := make([]string, 0)

		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading file: %v\n", err)
		}

		stringSlice = strings.Fields(line)

		strArr = append(strArr, stringSlice)
		if err == io.EOF {
			break
		}
	}

	strArr = append(strArr[:0], strArr[1:]...)
	fmt.Println(strArr)

	for _, slice := range strArr {
		fmt.Println(alternatingCharacters(slice[0]))
	}
}

func alternatingCharacters(s string) int32 {
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			count++
		}
	}

	return int32(count)
}
