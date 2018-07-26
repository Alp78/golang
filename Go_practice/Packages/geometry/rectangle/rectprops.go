package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle package initialized")
}

// Area returns the area of the rectangle
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal returns the diagonal of the rectangle
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
