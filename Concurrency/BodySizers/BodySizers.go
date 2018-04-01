/*
program that sends a number of requests to urls and prints out their size in byte.

the program launches a number of concurrent go routines (to demonstrate loadbalancing)
each calling the same "worker" function, that itself calls internally another "getPage" function
which sends a request to a url, calculates its size (bytes) and returns this size to "worker".

Once "worker" gets the size of the url from "getPage", it transmits the size of the webpage,
and the ID of the routine (used for that that specific url request) down a string channel
which is received and printed out by the program.

This "worker" function uses 2 string channels:
	string a) receiving the urls from the program
	string b) transmitting the formatted size of urls and routine ID back to the program

*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// function "getPage" taking a string as argument, and returning an integer (the size) and an error.
// it records the information of the webpage in a response (and returns an error if any) with http.Get
// then reads the content of the body of the page with Body.ReadAll
// and returns the size (number of bytes) of that body
// closing the body once done
func getPage(url string) (int, error) {

	// Get: url as param, returns a response (byte slice) and error
	resp, err := http.Get(url)
	// if error, return 0
	if err != nil {
		return 0, err
	}

	// close the body after each request (specified)
	defer resp.Body.Close()

	// read the body of the webpage
	body, err := ioutil.ReadAll(resp.Body)
	// if error, return 0
	if err != nil {
		return 0, err
	}
	// returns the number of bytes of the body and no error
	// (if there was an error, the above conditions would have been triggered)
	return len(body), nil
}

// function "worker" taking 2 string channels and 1 integers as arguments.
// listens to channel urlCh on which the urls will be received from the program
// calls for getPage() to get the size of each url
// then transmits back the string formatted sizes of the urls down the sizeCh channel
// indicating which routine is used (id)
func worker(urlCh chan string, sizeCh chan string, id int) {
	// assigns the url received from urlCh channel
	url := <-urlCh

	// invokes getPage function with the received url
	length, err := getPage(url)
	if err == nil {
		// if no error, transmits down the sizeCh channel
		//Sprintf: returns a string without printing
		sizeCh <- fmt.Sprintf("The size of %s is %d bytes (%d)", url, length, id)
	} else {
		sizeCh <- fmt.Sprintf("%s could not be reached %s ", url, err)
	}
}

func main() {

	// defining a slice of urls
	urls := []string{"http://www.google.com", "http://www.yahoo.com",
		"http://www.bing.com", "http://www.seznam.cz", "http://www.concur.com"}

	// creation of a transmitting channel to send the urls
	urlCh := make(chan string)

	// creation of a receiving channel for the url sizes
	sizeCh := make(chan string)

	// launching a number of concurrent go routines (for showing loadbalancing)
	// calling the "worker" function
	// each routine will have a unique ID assigned to it (i)
	for i := 0; i < 10; i++ {
		go worker(urlCh, sizeCh, i)
	}

	// sending the urls through urlCh channel
	for _, url := range urls {
		// sending urls to urlCh channel
		urlCh <- url
	}

	// printing size transmitted through sizeCh channel
	for i := 0; i < len(urls); i++ {

		fmt.Printf("%s\n", <-sizeCh)
	}

	// sending a couple more urls
	// demonstrating that both channels are still open
	// and loadbalancing
	urls2 := []string{"http://www.soundcloud.com", "http://www.facebook.com",
		"http://www.github.com", "http://hub.docker.com"}
	// sending the urls through urlCh channel
	for _, url := range urls2 {
		// sending urls to urlCh channel
		urlCh <- url
	}

	// printing formatted strings transmitted through sizeCh channel
	for i := 0; i < len(urls2); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}
}
