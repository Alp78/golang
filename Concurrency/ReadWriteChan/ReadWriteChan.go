/*
Program that launches two go routines on one same channel, communicating with each other:

a) writer: generates and transmits random integers down the channel for a set period of time
pauses the process for another set period of time, and then restarts the transmission indefinitely

b) reader: listens to the same channel, and prints out the recieved integers until a timeout is reached

"nil" value is affected to the channel in each respective functions, telling these function to either
stop transmitting (writer) or receiving (reader) on that channel

the program sets a duration in which this mechanism keeps going and thus prints out
the integers through "reader" that are generated and transmitted in "writer".

the prgram exits once its timeout is reached
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// "reader" function taking an integer channel as argument.
// keeps listening to its channel and prints out the integers transmitted up the channel
// each time it receives one (from any source but here it will be from "writer" function)
// stops listening after a set period of time (timeout)

func reader(ch chan int) {
	// setting the timer to 5 seconds
	t := time.NewTimer(5 * time.Second)

	// infinite loop with select:
	// unless the timeout is reached, it keeps listening to the channel
	// stores the received integers in "i" and prints them out
	for {
		select {
		// if the channel receives, print the transmitted integer
		case i := <-ch:
			fmt.Printf("%d\n", i)

		// after 5 seconds receiving, set channel to nil -> "close" the channel
		case <-t.C:
			ch = nil
		}
	}
}

// "writer" function taking an integer channel as argument.
// a select statement listens to 3 cases concurrently:
// 	- case1: generates and send random integers down the channel each time the channel is triggered
//	- case2: stops the routine when the "stopper" timer has reached its limit
//	- case3: restart the routine when the "restarter" timer has reached its limit
func writer(ch chan int) {

	// defining both timers
	stopper := time.NewTimer(1 * time.Second)
	restarter := time.NewTimer(3 * time.Second)

	// affecting the channel to a variable than can be reused to call itself
	savedCh := ch

	// infinite for loop with select statement
	// keeps generating and transmitting random integers (between 0-50) on the same channel as "reader"
	// unless any other case tells it to stop
	for {
		select {
		// first case keep transmitting random numbers on "ch" channel
		// as it transmits on the same channel as "reader", then "reader" receives them all
		// case without instruction = execute itself
		case ch <- rand.Intn(50):
			// stop transmission on "ch" channel when "stopper" reaches its timeout (1s)
		case <-stopper.C:
			ch = nil
			// restart transmission on "ch" channel when "restarter" reaches its timeout (3s)
			// by affecting the channel to itself
		case <-restarter.C:
			ch = savedCh
		}
	}
}

func main() {
	// creating one channel for both read/write operations
	ch := make(chan int)

	// launching both routines with the same "ch" channel
	// therby connecting the two functions' I/O's
	go reader(ch)
	go writer(ch)

	// both routines will keep working until this timeout timer terminates the program
	// obviously must be set to the same value as the "reader" timeout, or longer...
	// (otherwise the program terminates before the function reached its internal timeout)
	time.Sleep(6 * time.Second)
}
