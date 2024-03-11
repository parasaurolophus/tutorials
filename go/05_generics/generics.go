// Copyright Kirk Rader 2024

package main

import "fmt"

type (

	// Type constraint for any non-negative number that can be represented as a
	// binary value of up to 64 bits, i.e. Go's approximation of the set of
	// natural numbers.
	Natural interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	// Type constraint for any number that can be represented as a
	// twos-complement binary value of up to 64 bits, i.e. Go's approximation of
	// the set of integers.
	Integer interface {
		int | int8 | int16 | int32 | int64
	}

	// Type constraint for single- and double-precision IEEE-754 floating-point
	// numbers.
	Float interface {
		float32 | float64
	}

	// Type constraint for Go's approximation of the set of real numbers.
	Real interface {
		Natural | Integer | Float
	}

	// Type constraint for 64- and 128-bit complex numbers.
	Complex interface {
		complex64 | complex128
	}

	// Type constraint for all of Go's types intended for use in numeric
	// calculations. (Note that this deliberately excludes the "number-like"
	// types byte and rune, since they are not intended for use in ordinary
	// mathematical calculations.)
	Number interface {
		Real | Complex
	}
)

// Define a variadic generic function that calculates the sum of a given set of
// Numbers.
//
// Note that a limitation of Go's generic functions is that even though Sum() is
// defined for any type that satisfies the Number constraint, all of the
// parameters in a given invocation must be of the same underlying type.
//
// Another limitation of Go's generics is that while you can define generic
// functions, you cannot define generic interface methods.
//
// Finally, note that type constraints like Integer, Real, Number etc. are not
// actual types, despite the overloading of the interface syntax. That means you
// cannot use them as types when defining variables or return values outside of
// the definitions of generic functions.
func Sum[N Number](numbers ...N) N {

	var result N

	for _, n := range numbers {
		result += n
	}

	return result
}

// Prints:
//
//	6
//	6.5
//	(0+6i)
//
// to stdout.
func main() {

	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(0.9, 2.1, 3.5))
	fmt.Println(Sum(1i, 2i, 3i))
}
