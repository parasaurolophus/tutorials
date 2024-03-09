// Copyright Kirk Rader 2024

package main

import "fmt"

type (

	// Type constraint for any twos-complement integer.
	Integer interface {
		int | int8 | int16 | int32 | int64
	}

	// Type constraint for any unsigned integer.
	Natural interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	// Type constraint for single- or double-precision IEEE 754 floating-point
	// number.
	Float interface {
		float32 | float64
	}

	Real interface {
		Integer | Natural | Float
	}

	// Type constraint for any complex nunber.
	Complex interface {
		complex64 | complex128
	}

	// Type constraint for any number.
	Number interface {
		Integer | Natural | Float | Complex
	}
)

// Generic function to calculate the sum of a set of values of any numeric type.
func Sum[N Number](n ...N) N {

	var result N

	for _, v := range n {

		result += v
	}

	return result
}

// Prints
//
//	      15
//	      15.0
//	(15.0+21.0i)
//
// to stdout.
func main() {

	// Syntax error because Number is a constraint, not a type.
	//         |
	//         V
	// var n Number = 0

	fmt.Printf("%8d\n", Sum(0, 1, 2, 3, 4, 5))
	fmt.Printf("%10.1f\n", Sum(0.1, 0.9, 2.5, 2.5, 3.8, 5.2))
	fmt.Printf("%.1f\n", Sum(1i, 1+2i, 2+3i, 3+4i, 4+5i, 5+6i))
}
