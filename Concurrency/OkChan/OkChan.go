/*
Program that launches a go routine which transmits string elements of a slice down a channel
until all elements of the collection have been transmitted then closes the channel

"ok" boolean statement: once the channel is closed, trying to receive from it normally would
trigger a fatal error (deadlock), but "ok" handles it and returns "false" instead
*/

package main

import (
	"fmt"
)

func emit(c chan string) {

	// defining a string slice
	words := []string{"The", "quick", "brown", "fox"}

	// transmits each element of the slice down the channel
	for _, w := range words {

		c <- w
	}
	// close the channel once done
	close(c)

}

func main() {
	// create a string channel
	wordChannel := make(chan string)

	// emit() runs concurrently with main() --> go
	go emit(wordChannel)

	// receives elements transmitted from the function through the channel and prints them out
	// once all elements are passed, the channel closes

	for word1 := range wordChannel {
		fmt.Printf("%s ", word1)
	}

	// if the channel is closed (here after passing all 4 elements)
	// it returns an error (ok = false)
	word1, ok := <-wordChannel
	fmt.Printf("%s %t\n", word1, ok)

	fmt.Printf("\n")

	/*
		// looping is equivalent to explicitely receiving each element separately from the channel
		// in this case the routine will wait between each statement until it transmits all 4 elements
		// before closing the channel

		// receive and print the first element..
		word2 := <-wordChannel
		fmt.Printf("%s ", word2)

		// second..
		word2 = <-wordChannel
		fmt.Printf("%s ", word2)

		// third..
		word2 = <-wordChannel
		fmt.Printf("%s ", word2)

		// fourth..
		word2 = <-wordChannel
		fmt.Printf("%s ", word2)

	*/

}
