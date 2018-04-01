/*
Program that shows the loadbalancing of go routines at work calling for a same function
through a buffered channel, printing their unique ID along with "[" when they start
and "]" when they stop, and replacing each ended routine with a new one.

Th program launches a large number of go routines through a buffered channel
which limits the number of routines drastically, acting as a semaphore.

The routines call for the same function that waits for an input in its boolean channel,
then prints out the start symbol and routine ID, and waiting during a random period of time
before ending the routines and print the stop symbol.
After a routine has ended, a new one is launched to replace it.
This would go on indefinitely unless the program exits.

--> results show that only the last routine is replaced, while the first 9 ones keep running
*/

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	// declaring an int64 which is the format required by atomic.AddInt64
	// which will be referred to with a pointer in "semaWorker" function
	// no need to assign 0, go does it by default
	running int64
)

// function that takes 1 boolean channel as argument.
// it listens to the channel, once a signal is received
// an atomic iterator is increased by 1 and prints out a "[" symbol and that ID number
// then waits for random period of time (between 1-10 seconds)
// before decreasing the iterator by 1 and printing a "]" symbol.
// finally transmits "true" on the channel, which will trigger another round
// this would go on indefinitey unless the program exits
func semaWorker(sema chan bool) {
	// wait to execute instructions until the channel receives a signal
	<-sema

	// adding an atomic counter -> adds the value set as second param
	// to the variable (must be int64) set as first param as an address
	// (& is the go symbol for pointers) -> the variable is modified at its
	// original address, not as a new instance of it -> dereferencing / indirecting
	// it would be equivalent to i++ with a regular variable
	atomic.AddInt64(&running, 1)

	// printing "[" and the atomic value to show when and which routine starts
	fmt.Printf("[%d", running)

	// allowing a random time (between 1-10 seconds) during which
	// the other routines are launched and before closing a routine
	// and replace it by a new one
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	// printing "]" and the same ID to show when that routine stops
	fmt.Printf(" %d]", running)

	// decrementing the atomic counter once the routine is done
	// equivalent to i--
	atomic.AddInt64(&running, -1)

	// once a routine is completed, sends a signal down the channel
	// this will trigger another routine, replacing the closed one
	// until timeout is reached
	sema <- true
}

func main() {
	// buffered channel: give it an x int value
	// --> the channel can store up to x values/routines
	// once it gets filled up to its capacity, requests sent ont that channel
	// will be on hold until there's an available slot (if one of the x routines stops)
	// -> control concurrency = semaphore = how many routines can run simultaneously

	// allowing only 10 routines on this channel
	sema := make(chan bool, 10)

	// launching 1000 routines -> only 10 will be allowed to work on the channel
	for i := 0; i < 1000; i++ {
		go semaWorker(sema)
	}

	// filling the channel up to its capacity
	for i := 0; i < cap(sema); i++ {
		// transmitting 10 signals will run 10 concurrent functions
		sema <- true
	}

	// allocating some time during which the function has time to process the routines
	// before terminating the program
	time.Sleep(15 * time.Second)

	fmt.Printf("\n")
}
