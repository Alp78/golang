package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	fmt.Println("send London")
	ch <- "London"
	fmt.Println("send New York")
	ch <- "New York"
	fmt.Println("send Paris")
	ch <- "Paris"
	fmt.Println("send Prague")
	ch <- "Prague"
	fmt.Println("send Tokio")
	ch <- "Tokio"

}

func getData(ch chan string) {
	var input string
	// for loop breaks when channel is empty
	for {
		input = <-ch
		fmt.Println(input)
	}
}
