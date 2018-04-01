/*
// program that shows how to stop go routines safely by closing the channels and demonstrates the coordination between go routines
with "select" statement

- two ways of terminating cleanly a go routine:
1) through a channel dedicated to receiving a signal that will trigger the end of the routine
--> case: <-done
2) with a timer inside the function that will trigger the end of the routine after a set period of time
--> case: <-t.C

- select: listens or transmits on different channels concurrently in the same function
- case: each executing a specific set of instructions if the case is satisfied
if one of the "case" is triggered, "select" will run the concerned case's isntructions

Nb. for demonstrating the effect of both channels:
the code in main is active for one or the other case scenario at a time
keeping them both will show only the effect of one of them or generate a fatal error:
- if timer is first: done channel would be close after the timeout -> deadlock
- if done is first: timeout will not be triggered and show its effect
*/

package main

import (
	"fmt"
	"time"
)

// function taking 1 string channel and 1 integer channel as arguments.
// sets of slice of strings and waits a set period of time
// before lauching an infinite "for" loop that listens to both channels
// through a "select" statement which will execute the instructions
// of each "case" triggered by one of the two channels:
// 	- case a) loops the injection of the string slice in the string channel indefinitely
//			 --> it will run forever unless the timer goes out
//	- case b) if the boolean channel is triggered, the routine is ended
func emitSelect(wordChannel chan string, done chan bool) {
	// defer the closure of both channels once the routine is terminated
	defer close(wordChannel)
	defer close(done)

	// defining a string slice
	words := []string{"The", "quick", "brown", "fox"}

	//initializing the iterator
	i := 0

	// defining a timer - timeout of 3 seconds
	t := time.NewTimer(3 * time.Second)

	// infinite loop - keeps running until a signal is sent or received on either channel
	// which would instruct the loop to terminate
	for {
		// select: selects the channel that is called - either receive or send
		// this happens concurrently
		select {

		// the iteration transmits the words[i] element down the wordChannel
		// as it is triggered, "select" execute the case's instruction
		// unless another case instructs to terminate
		case wordChannel <- words[i]:
			i++
			// looping around the slice if the length of words[] is reached
			if i == len(words) {
				i = 0
			}
		// if there is send/receive signal on "done" channel, then return (terminate the function)
		// here the channel receives the value true, so it terminates and cleanly close the channel
		// this will stop the function to transmit on wordChannel as well
		// as the function is terminated: no more calls on wordChannel
		// defer closing will occur then, so both channels are closed when the routine quits
		case <-done:
			fmt.Printf("\nDone!\n")
			// send back true value - can be false, whatever
			// just to show the case is working
			done <- true
			return

		// in case the routine runs for more than 3 seconds, terminate
		// C: time channel
		case <-t.C:
			// to show the case is working
			fmt.Printf("\nTimeout!\n")
			return
		}
	}
}

func main() {
	wordChan := make(chan string)
	doneChan := make(chan bool)

	go emitSelect(wordChan, doneChan)

	/*
		// case done: the routine will run for 30 rounds then will wait for further signal
		for i := 0; i < 30; i++ {
			fmt.Printf("%s ", <-wordChan)
		}
		// a value is sent on doneChan - it can be true or false, whatever
		// what triggers the case in the function is the call itself
		// as case "done" is triggered in the function, the routine terminates
		doneChan <- true

		// waits a signal through doneChan to keep the program going
		// here we send back "true" value
		<-doneChan

		// receives and prints the value sent on doneChan: true here but also could be false
		fmt.Printf("%t\n", <-doneChan)
	*/

	// case function timer:
	// channel receives and prints out the elements until it is closed
	// here until the timer has reached its limit: 3s
	for word := range wordChan {
		fmt.Printf("%s ", word)
	}
}
