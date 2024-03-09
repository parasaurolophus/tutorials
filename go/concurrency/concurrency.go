// Copyright Kirk Rader 2024

package main

import (
	"fmt"
	"os"
)

func helper(n int) int {

	return n + 1
}

// Function intended to be invoked as a goroutine.
func worker(ch chan int, h func(int) int) {

	// Use defer to ensure that panics in injected dependencies do no cause
	// crashes and that the channel is closed on exit.
	defer func() {

		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "recovered: %v", r)
		}

		close(ch)
	}()

	// Write some values to the channel, then exit.
	for n := 0; n < 10; n += 1 {

		ch <- h(n)
	}
}

// Prints
//
//	1 2 3 4 5 6 7 8 9 10
//
// to stdout.
func main() {

	// Make a channel that transmits int values.
	ch := make(chan int)

	// Launch a worker goroutine, passing it ch.
	go worker(ch, helper)

	// Consume values from ch until the worker closes it.
	for n := range ch {

		fmt.Printf(" %d", n)
	}
}
