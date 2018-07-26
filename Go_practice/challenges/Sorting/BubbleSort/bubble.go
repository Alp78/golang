package main

import "fmt"

func main() {
	s := []int32{3, 2, 1}
	countSwaps(s)
}

func countSwaps(a []int32) {

	count := 0

	// Bubble Sort
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				count++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", count)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[len(a)-1])
}
