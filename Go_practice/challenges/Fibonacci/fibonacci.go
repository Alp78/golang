package main

import "fmt"

func main() {

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// func fibonacci() func() int {

// 	n := 0

// 	a := 0
// 	b := 1
// 	c := 1

// 	return func() int {
// 		var res int
// 		switch n {
// 		case 0:
// 			n++
// 			res = 0
// 		case 1:
// 			n++
// 			res = 1
// 		default:
// 			res = c
// 			a = b
// 			b = c
// 			c = a + b
// 		}
// 		return res
// 	}
// }

// func fibonacci() func() int {
// 	n := 0
// 	return func() int {
// 		a, b := 0, 1
// 		for i := 0; i < n; i++ {
// 			a, b = b, a+b
// 		}
// 		n++
// 		return a
// 	}
// }
func fibonacci() func() int {
	a := 0
	b := 1

	return func() int {
		c := a
		a = b
		b += c

		return c
	}
}
