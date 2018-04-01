/* program that runs a number of identical functions continuously and in parallel
through go routines calling the same function
- showing how go does the loadbalancing
- showing how many go routines can be coordinated by a single channel (here that just stops the routines)

select: listens to both function eninputs: channel and variable concurrently
and execute the set of instructions in each "case" whichever is satisfied
*/

package main

import (
	"fmt"
	"time"
)

// function taking 1 string variable, and 1 boolean channel
// prints the string received in msg variable indefinitely
// unless the routine receives a signal on goCh to stop
func printer(msg string, stopCh chan bool) {

	for {
		select {
		// in case a signal is sent on stopCh, exit the function
		case <-stopCh:
			return
		// by default print the content of msg variable indefinitely
		// as there are 10 routines running the same function -> go laodbalances them
		default:
			fmt.Printf("%s\n", msg)
		}
	}
}

func main() {
	// creating the stop channel
	stopCh := make(chan bool)

	// starting 10 go routines
	for i := 0; i < 10; i++ {
		// routines started and ran concurrently, corrdinated by the stopCh (initially empty)
		// i indiactes which routine does the print job -> shows loadbalancing between routines
		// Sprintf sends a string to the routine without printing
		// -> the printing is done inside the function
		go printer(fmt.Sprintf("printer: %d", i), stopCh)
	}

	// program waits 2 seconds and sends a signal on stopCh -> all go routines stop
	time.Sleep(2 * time.Second)
	stopCh <- true

	// program waits 2 seconds and terminates
	time.Sleep(2 * time.Second)
}
