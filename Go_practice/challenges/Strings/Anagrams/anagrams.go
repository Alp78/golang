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
		fmt.Println("Error Openning the file.")
		os.Exit(1)
	}

	reader := bufio.NewReader(file)

	words := make([]string, 0)

	for {
		line, err := reader.ReadString('\n')
		words = append(words, strings.Fields(line)[0])

		if err == io.EOF {
			break
		}
	}
	fmt.Println(words)

	fmt.Println(makeAnagram(words[0], words[1]))
}

func makeAnagram(a string, b string) int32 {
	count := 0
	aRunes := make([]rune, 0)
	bRunes := make([]rune, 0)

	for _, r := range a {
		aRunes = append(aRunes, r)
	}

	for _, r := range b {
		bRunes = append(bRunes, r)
	}

	if len(aRunes) >= len(bRunes) {
		for i, aRune := range aRunes {
			for j, bRune := range bRunes {
				if aRune == bRune {
					aRunes[i] = '!'
					bRunes[j] = '!'
					break
				}
			}
		}
	} else {
		for i, bRune := range bRunes {
			for j, aRune := range aRunes {
				if bRune == aRune {
					bRunes[i] = '!'
					aRunes[j] = '!'
					break
				}
			}
		}
	}

	for _, aRune := range aRunes {
		if aRune != '!' {
			count++
		}
	}

	for _, bRune := range bRunes {
		if bRune != '!' {
			count++
		}
	}

	return int32(count)

}
