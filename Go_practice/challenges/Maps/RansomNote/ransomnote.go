package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {

	file, err := os.Open("./input/input11.txt")
	if err != nil {
		fmt.Printf("There was an error openning the file: %v", err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	tempSlice := make([]string, 0)

	for {
		line, err = reader.ReadString('\n')
		tempSlice = append(tempSlice, line)

		if err == io.EOF {
			break
		}
	}

	tempSlice = append(tempSlice[:0], tempSlice[1:]...)
	magazine := strings.Fields(tempSlice[0])
	note := strings.Fields(tempSlice[1])

	start := time.Now()
	checkMagazine(magazine, note)
	fmt.Printf("Done in: %v\n", time.Since(start))

}

// func checkMagazine(magazine []string, note []string) {
// 	magMap := make(map[int]string)
// 	noteMap := make(map[int]string)

// 	for i, word := range magazine {
// 		magMap[i] = word
// 	}

// 	for i, word := range note {
// 		noteMap[i] = word
// 	}

// 	for i := range noteMap {
// 		for j := range magMap {
// 			if noteMap[i] == magMap[j] {
// 				delete(magMap, j)
// 				delete(noteMap, i)
// 			}
// 		}
// 	}

// 	if len(noteMap) == 0 {
// 		fmt.Println("Yes")
// 	} else {
// 		fmt.Println("No")
// 	}

// }
func checkMagazine(magazine []string, note []string) {
	count := 0
	for i := range note {
		for j := 0; j < (len(magazine) - count); j++ {
			if note[i] == magazine[j] {
				magazine[j], magazine[(len(magazine)-1-count)] = magazine[(len(magazine)-1-count)], magazine[j]
				count++
				break
			}
		}
	}
	if len(note) == count {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
