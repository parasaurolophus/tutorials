// Copyright Kirk Rader 2024

package main

import "fmt"

type (

	// Function signatures are first-class types, e.g. than be given names.
	Getter      func() int
	Setter      func(newValue int)
	Transformer func(value int) int
)

// Closures are first-class data values that can be passed as parameters.
//
// Note that the technique shown here could be combined with the approach
// demonstrated by the Sum() example for generic functions to turn this into a
// more general-purpose implementation of a specific form of Continuation
// Passing Style (CPS) in Go.
func Compose(value int, transformers ...Transformer) int {

	for _, transformer := range transformers {
		value = transformer(value)
	}

	return value
}

// Closures can also be returned as values.
//
// Note that the term "closure" is used rather than "function" for such data
// values because closures like getter and setter created in the body of this
// function combine both data and executable code. Each such "closure" encloses
// all of the bindings for variables defined within the lexical scope in which
// it is created. When the closure's function is invoked, it is invoked with
// that enclosed environment's bindings in effect, not those of the lexical
// environment in which the invocation occurred. When more than one closure is
// created within a given environment, they all enclose the same set of
// bindings.
//
// Closures thus provide the functionality of both the "objects" and "methods"
// of object-oriented programming systems. The "fields" of such objects are the
// enclosed lexical bindings. The "methods" are the various closures' function
// bodies. (In fact, the world's first commercially significant object-oriented
// programming language, Flavors on the Symbolics Lisp Machine, was implemented
// using closures in exactly the way demonstrated here by MakeInstance() -- but
// we digress.)
func MakeInstance(value int) (Getter, Setter) {

	getter := func() int { return value }
	setter := func(newValue int) { value = newValue }
	return getter, setter
}

// Prints:
//
//	 2
//	 0, 42
//	-1, 42
//	-1, 43
//
// to stdout.
func main() {

	addOne := func(value int) int { return value + 1 }
	subtractOne := func(value int) int { return value - 1 }
	fmt.Printf("%2d\n", Compose(0, addOne, addOne, subtractOne, addOne))

	getter1, setter1 := MakeInstance(0)
	getter2, setter2 := MakeInstance(42)

	fmt.Printf("%2d, %2d\n", getter1(), getter2())

	setter1(-1)
	fmt.Printf("%2d, %2d\n", getter1(), getter2())

	setter2(getter2() + 1)
	fmt.Printf("%2d, %2d\n", getter1(), getter2())
}
