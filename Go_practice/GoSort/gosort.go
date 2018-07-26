package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"./algorithms"
)

var slices []algorithms.SortSlice

type sortChan chan string

func main() {
	s1 := ParseFileToSortSlice("input00.txt")
	s2 := ParseFileToSortSlice("input01.txt")
	s3 := ParseFileToSortSlice("input08.txt")
	s4 := ParseFileToSortSlice("input14.txt")

	slices = append(slices, s1, s2, s3, s4)

	rCh := make(sortChan)

	go RoutineManager(slices, rCh)

	for {
		s, ok := <-rCh
		if ok == false {
			break
		}
		fmt.Println(s)
	}

}

// RoutineManager spawns functions and writes the returns in the channel
func RoutineManager(slices []algorithms.SortSlice, sCh sortChan) {
	for _, slice := range slices {
		sCh <- CalculateExecutionTime(slice.BubbleSort)
		sCh <- CalculateExecutionTime(slice.InsertionSort)
		sCh <- CalculateExecutionTime(slice.SelectionSort)
		sCh <- CalculateExecutionTime(slice.QuickSort)
	}
	close(sCh)
}

// CalculateExecutionTime prints a formatted string
// rendering the exxecution time of the function passed as argument
func CalculateExecutionTime(f func() []int) string {
	start := time.Now()
	time.Sleep(1 * time.Millisecond)
	sorted := f()
	elapsed := time.Since(start) - (1 * time.Millisecond)
	nameFull := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	functionName := strings.TrimSuffix(strings.TrimPrefix(filepath.Ext(nameFull), "."), "-fm")
	return fmt.Sprintf("Slice of length %v took %v to be sorted by %v\n", len(sorted), elapsed, functionName)
}

// ParseFileToSortSlice opens a file of space-separated integers and stores it as a SortSlice type
func ParseFileToSortSlice(filename string) algorithms.SortSlice {
	var sortSlice algorithms.SortSlice

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Enter filename: ")
	// scanner.Scan()
	// filename := scanner.Text()
	// if scanner.Err() != nil {
	// 	fmt.Printf("Error scanning the filename: %v\n", scanner.Err())
	// }

	file, err := os.Open("./input/" + filename)
	defer file.Close()
	if err != nil {
		fmt.Printf("Error openning the file: %v", err)
		os.Exit(1)
	}

	fileReader := bufio.NewReader(file)
	for {
		line, err := fileReader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading the numbers: %v", err)
			os.Exit(1)
		}

		strSlice := strings.Fields(line)

		intSlice := make([]int, 0, len(strSlice))

		for _, num := range strSlice {
			tempInt, _ := strconv.Atoi(num)
			intSlice = append(intSlice, tempInt)
		}

		sortSlice = make(algorithms.SortSlice, len(intSlice))

		copy(sortSlice, intSlice)

		if err == io.EOF {
			break
		}
	}

	return sortSlice
}
