/*
Program that uses 2 different routines triggering 2 different functions each generating random numbers
and transmitting them down an int channel. This is to show how a routine can remain open and called
on demand versus a routine that stops after a countdown.

1) emitRand: transmits down 10 random numbers and closes the channel once done
2) makeID: generates and transmits a random number down the channel
each time a request is sent on that channel indefinitely

in case of 2) the program executes two sets of instructions:
	a) calls the function a number of times and directly prints out the random numbers
	b) calls the function a number of times and stores the random integers in a map,
	generating a unique key for each of them (randoms can randomly be identical....)
	then prints out the map in random order - no sorting
*/

package main

import (
	"fmt"
	"math/rand"
)

// function taking a int channel as argument.
// generates and transmits 10 random number between 1-1000 down the channel
// and then close the channel
func emitRand(c chan int) {
	i := 10
	for i > 0 {
		// math/rand method generating random positive integers
		// within a range (here 1000)
		c <- rand.Intn(1000)
		i--
	}
	close(c)
}

// function taking an int channel as argument.
// generates and transmits a random number on demand - infinite "for" and no "close()"
func makeID(idChan chan int) {
	for {
		// always be ready to transmit on idChan
		// as for range is not defined
		idChan <- rand.Intn(1000)
	}
}

func main() {
	// creating a channel to receive the numbers generated in emitRand
	randChannel := make(chan int)

	// launching a go routine using randChannel channel
	go emitRand(randChannel)

	// print each element in randChannel
	// once all elements have been transmitted through the channel
	// the channel closes as per instruction in the function
	for n := range randChannel {
		fmt.Printf("%d ", n)
	}
	fmt.Printf("\n\n")

	// create a new channel for makeID function
	newID := make(chan int)

	// launching a go routine using newID channel
	// -> channel remains open and ready to transmit the integers generated in makeID function
	go makeID(newID)

	// receives and prints out a new generated random integer
	// through newID channel each time makeID function is called
	fmt.Printf("New ID: %d\n", <-newID)
	fmt.Printf("New ID: %d\n", <-newID)
	fmt.Printf("New ID: %d\n", <-newID)
	fmt.Printf("New ID: %d\n", <-newID)

	// create a map to store ID's with unique keys
	idMap := make(map[int]int)

	for i := 0; i < 10; i++ {
		idMap[i] = <-newID
	}
	fmt.Printf("\n\n")
	fmt.Printf("Map of unique keys and id's:\n")

	// a map is never sorted - random order
	for k, v := range idMap {
		fmt.Printf("Key: %d ID: %d\n", k, v)
	}
}
