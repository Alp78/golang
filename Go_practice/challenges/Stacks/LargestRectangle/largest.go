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
	file, err := os.Open("./input/input08.txt")
	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	instructions := make([][]int32, 0)
	for {
		strLine := make([]string, 0)

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Problem reading the lines.")
			os.Exit(1)
		}

		strLine = strings.Fields(line)

		intLine := make([]int32, len(strLine))

		for i := range strLine {
			tempInt, _ := strconv.Atoi(strLine[i])
			intLine[i] = int32(tempInt)
		}
		instructions = append(instructions, intLine)
		if err == io.EOF {
			break
		}
	}

	buildings := make([]int32, len(instructions[1]))
	copy(buildings, instructions[1])

	fmt.Println(largestRectangle(buildings))
}

func largestRectangle(h []int32) int64 {
	var n int32
	var count int
	areas := make([]int32, 0, len(h))

	for i, b := range h {
		count = 1
		for j := i + 1; j < len(h); j++ {
			if h[j] >= b {
				count++
			} else {
				break
			}
		}
		for j := i - 1; j >= 0; j-- {
			if h[j] >= b {
				count++
			} else {
				break
			}
		}
		n = b * int32(count)
		areas = append(areas, n)
	}

	for i := 0; i < len(areas); i++ {
		for j := 0; j < len(areas)-1; j++ {
			if areas[j+1] < areas[j] {
				areas[j+1], areas[j] = areas[j], areas[j+1]
			}
		}
	}

	max := areas[len(areas)-1]

	return int64(max)
}
