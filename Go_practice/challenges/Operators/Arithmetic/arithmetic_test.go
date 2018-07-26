package main

import "testing"

func TestSolve(t *testing.T) {
	expected := "The total meal cost is 15 dollars."
	var mealCost = 12.00
	var tipPercent int32 = 20
	var taxPercent int32 = 8
	if observed := solve(mealCost, tipPercent, taxPercent); expected != observed {
		t.Fatalf("Got %s want %s", observed, expected)
	}
}
