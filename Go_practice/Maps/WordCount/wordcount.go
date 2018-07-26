package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordCount("I ate a donut. Then I ate another donut."))
}

func wordCount(s string) map[string]int {
	sSlice := make([]string, 0)
	sMap := make(map[string]int)

	sSlice = strings.Fields(s)

	for i := 0; i < len(sSlice); i++ {
		count := 1
		for j := 0; j < len(sSlice); j++ {
			if sSlice[i] == sSlice[j] && i != j {
				count++
			}
		}
		sMap[sSlice[i]] = count
	}

	return sMap
}
