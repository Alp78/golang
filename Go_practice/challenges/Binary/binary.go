package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var nb int64 = 1266
	bin := strconv.FormatInt(nb, 2)
	strPat := "1"
	for {
		if strings.Contains(bin, strPat) {
			strPat += "1"
		} else {
			break
		}
	}
	fmt.Println(len(strPat) - 1)

	fmt.Println(fib(40))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
