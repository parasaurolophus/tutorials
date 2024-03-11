// Copyright Kirk Rader 2024

package main

import (
	"fmt"
	"math"
)

// Worker goroutine.
//
// Send 0, 1, 2, 3, ... continuously on the values channel, terminating as soon
// as a value is received on the quit channel.
func worker(values chan int, quit chan bool) {

	// When terminating, close the channel on which the worker goroutine sends
	// values.
	defer close(values)

	n := 0

	// Note that this will be an infinite loop until and unless some other
	// goroutine sends a message on the quit channel.
	for {
		select {
		case values <- n:
			if n == math.MaxInt {
				n = 0
			} else {
				n += 1
			}
		case <-quit:
			goto terminate
		}
	}
terminate:
}

// Prints:
//
//	0
//	1
//	2
//	3
//	4
//
// to stdout.
func main() {

	// Create the channels that will be used to communicate between the main and
	// worker goroutines.
	values := make(chan int)
	quit := make(chan bool)

	// When terminating, close the channel on which the main goroutine sends
	// messages.
	defer close(quit)

	// Launch the worker goroutine.
	go worker(values, quit)

	// Consume all values until the worker closes that channel; note that this
	// would be an infinite loop if the worker were to fail to close the
	// channel.
	for value := range values {

		fmt.Println(value)

		// At some point, send a message to the worker using the quit channel
		// that will cause it to terminate and, using a deffered call to
		// close(calues), close that channel and so cause this goroutine to
		// exit the for loop.
		if value >= 4 {
			quit <- true
		}
	}
}
