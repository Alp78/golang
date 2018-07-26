package main

import "testing"

func TestMinimumBribes(t *testing.T) {
	expected := "3"
	q := []int32{2, 1, 5, 3, 4}
	if observed := minimumBribes(q); observed != expected {
		t.Fatalf("minimumBribes() = %v, want %v", observed, expected)
	}
}
