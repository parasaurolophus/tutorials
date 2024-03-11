// Copyright Kirk Rader 2024

package main

import (
	"fmt"
	"os"
)

// A function that returns the result of adding 1 to its parameter.
func AnOrdinaryFunction(n int) int {
	return n + 1
}

// A function that returns two values, the second of which implements the error
// interface.
func MultipleValues(n int) (int, error) {

	if n%2 == 1 {
		return 0, fmt.Errorf("only even numbers are supported: %d", n)
	}

	return AnOrdinaryFunction(n), nil
}

// Prints:
//
//	Hello, world!
//	Another way to say "Hello, world!"
//	0
//	1
//	2
//	10.0
//
// to stdout and:
//
//	only even numbers are supported: 41
//
// to stderr.
func main() {

	fmt.Println("Hello, world!")
	fmt.Printf("Another way to say \"%s, %s!\"\n", "Hello", "world")

	// Implicit variable declaration and type inference using := is equivalent,
	// in this case, to:
	//
	//	var n int = 42
	//
	// Which form is "better" is largely a matter of taste, but := enables
	// features such as composite expressions, as shown below.
	n := 41

	// Note that:
	//
	//  n = 4.2
	//
	// would produce a syntax error at this point because n's type was inferred
	// to be int and 4.2 is not assignable to an int variable.

	// Composite expression declares, binds, and uses err all in the test
	// expression of a single if statement:
	if i, err := MultipleValues(n); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		// note that the scope of variables like i and err, bound by
		// composite expressions, covers the entire statement in which they
		// occur.
		fmt.Printf("MultipleValues(%d) returned %d\n", n, i)
	}

	// Such composite expressions are also used in for loop syntax:
	for i := 0; i < 3; i += 1 {
		fmt.Println(i)
	}

	// Go borrows C's "pointer" syntax and a small subset of the corresponding
	// semantics. Here, the type of f is float64:
	f := 4.2

	// The type of p is *float64, i.e. "pointer to float64":
	p := &f

	// Pointers can be used to access the underlying value to which they point:
	*p += 5.8
	fmt.Printf("%.1f", f)
}
