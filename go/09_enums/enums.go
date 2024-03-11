// Copyright Kirk Rader 2024

package main

import (
	"fmt"
	"os"
	"unicode"
)

// Go does not have a mechanism that corresponds exactly to enum types in
// languages like C++, C#, Java etc. but you can create aliases for scalar types
// like int.
type MyEnum int

// There is a special syntax involving the iota keyword for treating a const
// block as the definition of a set of monotoniclly increasing values of a
// scalar type starting at zero.
//
// In other words, while Go does not supported enumerated types, it does support
// enumerating particular values of any type. But it does not guarantee that
// such named values of a given type are unique nor numerically sequential. It
// is up to programmers to enforce such conventions in how constants are
// declared and used when the desire is to treat a given type as if it were an
// enum.
const (
	One MyEnum = iota
	Two
	Three
)

// Implement the fmt.Stringer interface to cause fmt.Printf() and the like to
// emit names instead of numbers for known values.
func (e MyEnum) String() string {

	switch e {
	case One:
		return "One"

	case Two:
		return "Two"

	case Three:
		return "Three"

	default:
		// As is implicit in this default case, there is no way to prevent a
		// programmer from coercing an arbitrary int value to be treated as a
		// value of MyEnum.
		return fmt.Sprintf("<MyEnum %d>", e)
	}
}

// Implement the fmt.Scanner interface to allow parsing names for known values
// instead of numbers.
func (e *MyEnum) Scan(state fmt.ScanState, _ rune) error {

	b, err := state.Token(true, unicode.IsLetter)

	if err != nil {
		return err
	}

	s := string(b)

	switch s {
	case One.String():
		*e = One

	case Two.String():
		*e = Two

	case Three.String():
		*e = Three

	default:
		return fmt.Errorf("unrecognized token, \"%s\"; expected a MyEnum", s)
	}

	return nil
}

// Prints:
//
//	One, Two, Three
//	Three, Two, One
//
// to stdout.
func main() {

	e1 := One
	e2 := Two
	e3 := Three

	fmt.Printf("%v, %v, %v\n", e1, e2, e3)

	n, err := fmt.Sscanf(" Three, Two, One  ", " %v, %v, %v ", &e1, &e2, &e3)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if n != 3 {
		fmt.Fprintf(os.Stderr, "expected 3, got %d", n)
	}

	fmt.Printf("%s, %s, %s\n", e1, e2, e3)
}
