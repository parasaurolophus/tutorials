_Copyright &copy; Kirk Rader 2024_

# Concurrency in Go

That Go's designers had concurrency in mind from the beginning is revealed by
the languages name. It is taken from the `go` keyword that is used to launch
what Go's documentation refers to as "goroutines" (like "co-routined created
using go;" don't forget to tip the wait staff!)

See [./concurrency.go](./concurrency.go) for an example.

It says on Go's packaging that goroutines are lighter weight than the threads
implemented natively by other languages like Java. Go's marketing makes much of
the fact that another of its built-in mechanisms, _channels_, can be used for
very low-overhead yet still safe communication between goroutines. Go's
creators' slogan, "Don't communicate by sharing memory, share memory by
communicating," is a reference to goroutines and channels.

And Go mostly, kind of, sort of lives up to that slogan. No language could, even
in principle, implement that slogan literally. Not only do Go's
[interfaces](../interfaces/), [closures](../functional/) and pointer types allow
for all-too easy subversion of that principle, but look at the standard idiom
for using goroutines, below. Note that the main goroutine allocates `ch`, passes
it to the worker goroutine and then synchronizes the two goroutines by waiting
for the worker to manage that shared channel's state. Busted!

In other words, even using channels according to the "official" idiom, Go
programmers need to code carefully when incorporating concurrency into their
design. Go's channels may provide a mechanism that is lighter weight than
traditional semaphores, mutexes and conditions, but channels give programmers
every bit as much rope by which to hang themselves (there's never a rim-shot
around when you need one!)

```go
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
```

As called out in the example code, you must take stack unwinding into account
when closing a channel. If you were to move the call to `close(ch)` to where it
"naturally" belongs -- after the `for` loop has completed -- your code would be
susceptible to deadly embrace if an injected dependency caused a panic.

At a higher level, you must think carefully about which goroutine closes a
channel at all. Naively, one might assume that the main goroutine should "own"
`ch` in the given example since it was the one that created it. But that is not
the standard pattern. Generally, the goroutine which sends values over a channel
should be the one to close it, without regard to which goroutine created it.
Conversely, any goroutine that consumes values from a channel usually should
wait for the channel to be closed before terminating. That way it knows that the
sender goroutine is actually finished performing its operations. But that relies
on well-behaved senders, as discussed above in regards to `defer` and `close()`.

For cases where the ultralight channel mechanism for synchronization of
goroutines is insufficient, Go's standard library also provides a repertoire of
other synchronization primitives modeled on those in other languages which, in
turn, are inspired by those of Linux' native _pthread_ library.
