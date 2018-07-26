package main

import "fmt"

func main() {

	ch := make(chan []int)
	for i := 1; i <= 10; i++ {
		go factorial(i, ch)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}

}

func factorial(n int, ch chan []int) {
	var res int = 1
	for i := 1; i <= n; i++ {
		res = res * i
	}

	ch <- []int{n, res}
}
