package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input/input00.txt")
	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(1)
	}
	// reader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(file)
	words := make([][]string, 0)
	for {
		line, err := reader.ReadString('\n')

		strLine := strings.Fields(line)

		words = append(words, strLine)

		if err == io.EOF {
			break
		}
	}

	words = append(words[:0], words[1:]...)
	for _, w := range words {
		evenOddSep(w[0])
	}
}

func evenOddSep(s string) {
	var even, odd string

	for i, c := range s {
		if i%2 == 0 {
			even += string(c)
		} else {
			odd += string(c)
		}
	}
	fmt.Printf("%s %s\n", even, odd)
}
