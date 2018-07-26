package algorithms

import (
	"math/rand"
	"time"
)

// SortSlice has sorting functions
type SortSlice []int

// BubbleSort returns a slice sorted with Bubble sort algorithm
func (a SortSlice) BubbleSort() []int {
	s := make([]int, len(a))
	copy(s, a)

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

// InsertionSort returns a slice sorted with Insertion sort algorithm
func (a SortSlice) InsertionSort() []int {
	s := make([]int, len(a))
	copy(s, a)

	b := make([]int, 0, len(s))
	min := 0
	index := 0
	for len(b) < len(s) {

		for _, c := range s {
			if c != 0 {
				min = c
				break
			}
		}

		for i := 0; i < len(s); i++ {
			if s[i] <= min && s[i] != 0 {
				min = s[i]
				index = i
			}
		}

		b = append(b, min)
		s[index] = 0
	}
	return b
}

// SelectionSort returns a slice sorted with Selection sort algorithm
func (a SortSlice) SelectionSort() []int {
	s := make([]int, len(a))
	copy(s, a)

	for i := 0; i < len(s)-1; i++ {
		/* set current element as minimum*/
		minIndex := i
		/* check remaining elements:
		if any smaller than min, swap */
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}

		/* swap the minimum element with the current element*/
		if minIndex != i {
			s[minIndex], s[i] = s[i], s[minIndex]
		}
	}

	return s
}

// QuickSort returns a slice sorted with Quicksort algorithm
func (a SortSlice) QuickSort() []int {

	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	source := int64(time.Now().Nanosecond())
	rand.Seed(source)
	pivotIndex := rand.Intn((len(a) - 1))

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	// sort the left part of the slice
	(a[:left]).QuickSort()
	// sort the right slice
	(a[left+1:]).QuickSort()

	// return the sorted slice
	return a
}
