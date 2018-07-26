package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input/input06.txt")
	if err != nil {
		fmt.Println("Error Openning the file.")
		os.Exit(1)
	}

	reader := bufio.NewReader(file)

	word := make([]string, 0)
	wordsArr := make([][]string, 0)

	for {
		line, err := reader.ReadString('\n')
		word = strings.Fields(line)

		wordsArr = append(wordsArr, word)
		if err == io.EOF {
			break
		}
	}

	wordsArr = append(wordsArr[:0], wordsArr[1:]...)
	for i := 0; i < len(wordsArr); i += 2 {
		fmt.Println(twoStrings(wordsArr[i][0], wordsArr[i+1][0]))
	}

}

func twoStrings(s1 string, s2 string) string {
	if len(s2) < len(s1) {
		for _, c := range s2 {

			if strings.Contains(s1, string(c)) {
				return "YES"
			}

		}
	} else {
		for _, c := range s1 {
			if strings.Contains(s2, string(c)) {
				return "YES"
			}
		}
	}
	return "NO"
}
