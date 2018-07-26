package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./input/input08.txt")
	if err != nil {
		fmt.Printf("Error openning the file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var line string
	intArr := make([][]int32, 0, 2)

	for {
		stringSlice := make([]string, 0)
		intSlice := make([]int32, 0)

		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading file: %v\n", err)
		}

		stringSlice = strings.Fields(line)

		var number int

		for _, strNum := range stringSlice {
			number, _ = strconv.Atoi(strNum)
			intSlice = append(intSlice, int32(number))
		}

		intArr = append(intArr, intSlice)
		if err == io.EOF {
			break
		}
	}

	prices := intArr[1]
	budget := intArr[0][1]

	fmt.Println(maximumToys(prices, int32(budget)))
}

func maximumToys(prices []int32, k int32) int32 {

	// phase 1: eliminate all prices above the budget
	for i := 0; i < len(prices); i++ {
		if prices[i] > k {
			prices[i], prices[len(prices)-1] = prices[len(prices)-1], prices[i]
			prices = prices[:len(prices)-1]
		}
		if i == len(prices)-1 && prices[i] > k {
			prices = prices[:len(prices)-1]
		}
	}

	// phase 2: sort the slice (Selection Sort)

	var f func(s []int32) []int32
	f = func(prices []int32) []int32 {

		if len(prices) < 2 {
			return prices
		}

		left, right := 0, len(prices)-1

		// Pick a pivot
		source := int64(time.Now().Nanosecond())
		rand.Seed(source)
		pivotIndex := rand.Intn((len(prices) - 1))

		// Move the pivot to the right
		prices[pivotIndex], prices[right] = prices[right], prices[pivotIndex]

		// Pile elements smaller than the pivot on the left
		for i := range prices {
			if prices[i] < prices[right] {
				prices[i], prices[left] = prices[left], prices[i]
				left++
			}
		}

		// Place the pivot after the last smaller element
		prices[left], prices[right] = prices[right], prices[left]

		// Go down the rabbit hole
		// sort the left part of the slice
		f(prices[:left])

		// sort the right slice
		f(prices[left+1:])

		// return the sorted slice
		return prices
	}

	prices = f(prices)

	// phase 3: combinations from minimum price
	var sum int32
	i := 0
	for {
		sum += prices[i]

		if i == len(prices)-1 {
			if sum > k {
				i--
				break
			}
			break
		}
		if sum >= k {
			i--
			break
		}
		i++

	}

	return int32(i + 1)
}
