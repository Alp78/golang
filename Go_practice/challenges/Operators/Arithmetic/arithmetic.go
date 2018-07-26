package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/input00.txt")
	if err != nil {
		fmt.Printf("Error openning the file: %s\n", err)
	}

	scanner := bufio.NewScanner(file)
	strSlice := make([]string, 0, 3)
	var (
		mealCost   float64
		tipPercent int32
		taxPercent int32
	)

	for scanner.Scan() {
		strSlice = append(strSlice, strings.Fields(scanner.Text())[0])
		if err == io.EOF {
			break
		}
	}

	mealCost, _ = strconv.ParseFloat(strSlice[0], 64)

	tempInt64, _ := strconv.ParseInt(strSlice[1], 10, 32)
	tipPercent = int32(tempInt64)

	tempInt64, _ = strconv.ParseInt(strSlice[2], 10, 32)
	taxPercent = int32(tempInt64)

	fmt.Println(solve(mealCost, tipPercent, taxPercent))
}

func solve(mealCost float64, tipPercent int32, taxPercent int32) string {
	var tip, tipPercent64, tax, taxPercent64, totalCost float64

	tipPercent64 = float64(tipPercent)
	taxPercent64 = float64(taxPercent)

	tip = (tipPercent64 * mealCost) / 100
	tax = (taxPercent64 * mealCost) / 100

	totalCost = math.Round(mealCost + tip + tax)
	return fmt.Sprintf("The total meal cost is %.0f dollars.", totalCost)
}
