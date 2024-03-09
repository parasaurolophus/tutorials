// Copyright Kirk Rader 2024

package main

import (
	"fmt"
	"os"
)

type (

	// Closures' signatures are types, e.g. they can be aliased.
	Handler func(int) int
)

// Closures are values, e.g. they can be passed as parameters.
func Compose(v int, handlers ...Handler) int {

	for _, handler := range handlers {

		// Functional composition is simply one closure calling another.
		v = handler(v)
	}

	return v
}

// Closures can be returned as values.
func Adder(increment int) Handler {

	// Closures can be stored in variables.
	c := func(n int) int {
		return n + increment
	}

	return c
}

// Prints 11 to stdout.
func main() {

	// Closures "remember" the lexical environment in which they were created.
	add1 := Adder(1)
	subtract1 := Adder(-1)
	add10 := Adder(10)
	n := Compose(0, add1, add1, subtract1, add10)

	if n != 11 {
		fmt.Fprintf(os.Stderr, "expected 11, got %d\n", n)
	}

	fmt.Println(n)
}
