package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://concur.com",
	}

	// declaration of a string chan
	// to communicate between go routines and Main
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// main receives the value from channel and prints is
	// as this is a blocking call
	// -> Main is put to sleep and waits an input in the channel
	// once it got the first input from the first go routine
	// it is satisifed and then exits the program
	// while the other go routines are up and not completed
	// fmt.Println(<-c)

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

// takes the channel as arg
func checkLink(link string, c chan string) {
	// GET is a blocking call
	_, err := http.Get(link)
	if err != nil {
		// send the formatted string down the channel
		c <- fmt.Sprintf("%s might be down!\n", link)
		return
	}
	// send the formatted string down the channel
	c <- fmt.Sprintf("%s is up!\n", link)

}
