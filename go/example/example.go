// Copyright Kirk Rader 2024

package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

type (

	// Type constraint for the set of non-negative numbers that can be
	// represented as binary values of up to 64 bits.
	Natural interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	// Type constraint for the set of numbers that can be represented as
	// twos-complement binary values of up to 64 bits.
	Integer interface {
		int | int8 | int16 | int32 | int64
	}

	// Type constraint for single- or double-precision IEEE-754 floating-point
	// values.
	Float interface {
		float32 | float64
	}

	// Type constraint for Integer or Float numbers.
	Real interface {
		Integer | Float
	}

	// Type constraint for 64- or 128-bit complex numbers.
	//
	// Note that these are single, first-class numeric values even though the
	// syntax for a complex number appears as a sum between a given pair of real
	// and imaginary numbers (where 0 is assumed as the real component when an
	// imaginary constant is specified without a real component).
	Complex interface {
		complex64 | complex128
	}

	// Type constraint for Real, Natural or Complex numbers; i.e. types that are
	// intended for use in mathematical calculations.
	Number interface {
		Real | Natural | Complex
	}

	// Type constraint for a value that may be a Number, byte or rune.
	//
	// While bytes and runes are binary values with number-like properties, they
	// are not intended to be used in ordinary mathematical calculations.
	NumberLike interface {
		byte | rune | Number
	}
)

var seed = time.Now().Unix()

// Worker goroutine function.
//
// Repeatedly sends the result of calling `next` to `valuesChan` until a value
// is received on `quitChan`.
//
// Keeps a count of numbers sent to `valuesChan` and sends it to `countChan`
// before terminating.
func worker[R Real](next func(R, *rand.Rand) R, valuesChan chan R, countChan chan int, quitChan chan bool) {

	// Count of numbers sent to `valuesChan`.
	count := 0

	// Send `count` to `countChan` when this worker terminates.
	//
	// Close `countChan` after sending.
	//
	// Do not close `valuesChan` because more than one goroutine sends to it.
	defer func() {
		countChan <- count
		close(countChan)
	}()

	// Each goroutine needs its own random number source.
	source := rand.NewSource(seed)
	generator := rand.New(source)

	// Current number to send to `valueChan`.
	var n R

done: // Loop until `break done`.

	for {

		// Increase the odds for all worker goroutines to run by initially
		// yielding the CPU in each of them at every iteration.
		//
		// Note this is not a blocking operation nor does it guarantee which
		// eligible goroutine is scheduled, so you will often still see some
		// goroutines run more frequently than others. But  yielding here
		// generally results in more even distribution of run time across all
		// the worker goroutines than if the first one to run just keeps sending
		// to `valuesChan` without hitting any blocking system calls. You
		// probably don't have to worry about this in more realistic programs
		// where the workers are likely to make blocking calls more frequently
		// than in this very contrived example.
		runtime.Gosched()

		// Wait for `valuesChan` to be available for sending or a value is
		// recieved on `quitChan`.
		select {

		// Send `n`, increment `count` and invoke `next(...)`.
		case valuesChan <- n:
			count += 1
			n = next(n, generator)

		// Terminate the loop when a value is received on `quitChan`.
		case <-quitChan:
			break done
		}
	}

}

// Function passed as first parameter to `go worker(...)` in `main()`.
func next[N Real](_ N, r *rand.Rand) N {
	return N(-5 + (10 * r.Float64()))
}

func main() {

	w := runtime.GOMAXPROCS(0)
	values := make(chan float64)
	quits := make([]chan bool, w)
	counts := make([]chan int, w)

	for i := 0; i < runtime.GOMAXPROCS(0); i += 1 {
		counts[i] = make(chan int)
		quits[i] = make(chan bool)
		go worker(
			next,
			values,
			counts[i],
			quits[i])
	}

	count := 0

	for n := range values {

		fmt.Printf("%3d: %f\n", count, n)
		count += 1

		if math.Abs(n) > 4 {
			break
		}
	}

	for _, q := range quits {
		q <- true
		close(q)
	}

	for _, c := range counts {
		fmt.Printf("%d ", <-c)
	}

	fmt.Println()
}
