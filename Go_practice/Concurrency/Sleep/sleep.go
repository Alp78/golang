package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")

	go longWait()
	go shortWait()

	fmt.Println("Sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("End of main()")
}

func longWait() {
	fmt.Println("Begin longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("End longWait()")
}

func shortWait() {
	fmt.Println("Begin shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("End shortWait()")
}
