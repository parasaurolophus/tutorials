// Copyright Kirk Rader 2024

package main

import "fmt"

// A struct with an int field named Value.
type MyStruct struct {
	Value int
}

// A function that receives a copy of whatever MyStruct was passed by value.
func AddOne(s MyStruct) int {
	s.Value += 1
	return s.Value
}

// A function that receives a pointer to the original MyStruct (ersatz pass by
// reference using pointers).
func Increment(s *MyStruct) int {

	// Note that Go provides a potentially confusing "convenience" syntax
	// whereby "." notation implicitly dereferences pointers. So the bodies of
	// AddOne() and Increment() look identical, but one is modifying a copy that
	// gets thrown away as soon as it returns while the other is modifying the
	// original value through a pointer. Passing by pointer can thus be a
	// significant performance optimization as well as implementing pass by
	// reference, but caution must then be used not to modify a value
	// accidentally due to Go's lack of support for anything like C++'s "pointer
	// to const" and "const reference" constructs.
	s.Value += 1
	return s.Value
}

// Prints:
//
//	43
//	42
//	43
//	43
//
// to stdout.
func main() {

	// Unlike many languages with garbage collectors, parameters in Go are
	// passed by value, even when an implicit copy of a struct is required.
	s := MyStruct{42}

	// Prints 43.
	fmt.Println(AddOne(s))

	// Prints 42, because a copy of s was passed to AddOne(), not a reference.
	fmt.Println(s.Value)

	// Prints 43. The parameter to Increment() isn't a MyStruct but, rather, a
	// *MyStruct, i.e. a "pointer to MyStruct".
	fmt.Println(Increment(&s))

	// Prints 43, because the previous call to Increment() modified s.Value
	// through a pointer.
	fmt.Println(s.Value)
}
