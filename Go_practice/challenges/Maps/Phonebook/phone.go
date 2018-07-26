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
	file, err := os.Open("./input/input00.txt")
	if err != nil {
		fmt.Printf("There was an error openning the file: %v", err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	tempSlice := make([]string, 0)

	strSlice := make([][]string, 0)
	for {
		str, _, err := reader.ReadLine()
		line := strings.TrimRight(string(str), "\r\n")
		if err != nil && err != io.EOF {
			fmt.Println("Error reading the line.", err)
			os.Exit(1)
		}
		tempSlice = strings.Fields(line)
		if len(tempSlice) != 0 {
			strSlice = append(strSlice, tempSlice)
		}

		if err == io.EOF {
			break
		}
	}
	nbEntries, e := strconv.Atoi(strSlice[0][0])
	if e != nil {
		fmt.Println("Error converting the entries number.")
		os.Exit(1)
	}

	strSlice = append(strSlice[:0], strSlice[1:]...)

	phoneBook := make(map[string]string)
	for i := 0; i < nbEntries; i++ {
		phoneBook[strSlice[i][0]] = strSlice[i][1]
	}

	var found bool
	var value string
	for i := nbEntries; i < len(strSlice); i++ {
		found = false
		if value, found = phoneBook[strSlice[i][0]]; found {
			fmt.Printf("%s=%s\n", strSlice[i][0], value)
			found = true
		}

		if !found {
			fmt.Println("Not found")
		}
	}
}
