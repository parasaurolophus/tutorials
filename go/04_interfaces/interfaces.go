// Copyright Kirk Rader 2024

package main

import "fmt"

type (

	// On the surface, Go's interface types resemble those in object-oriented
	// languages like C# and Java(but see below for the ways that Go's interface
	// differ from those in object-oriented programming systems).
	Counter interface {

		// Return the current value.
		Value() int

		// Add 1 to the current value.
		Increment()

		// Subtrace 1 from the current value.
		Decrement()
	}

	// Declare a scalar type that will implement the Counter interface.
	MyInt int

	// Declare a struct type that will also implement the Counter interface.
	MyStruct struct {
		value *int
	}
)

// Implement Counter.Value() for MyInt.
func (c MyInt) Value() int {
	return int(c)
}

// Implement Counter.Increment() for *MyInt.
func (c *MyInt) Increment() {
	*c += 1
}

// Implement Counter.Decrement() for *MyInt.
func (c *MyInt) Decrement() {
	*c -= 1
}

// Implement Counter.Value() for MyStruct.
func (c MyStruct) Value() int {
	return *c.value
}

// Implement Counter.Increment() for MyStruct.
func (c MyStruct) Increment() {
	*c.value += 1
}

// Implement Counter.Decrement() for MyStruct.
func (c MyStruct) Decrement() {
	*c.value -= 1
}

// Define a constructor for Counter that uses MyStruct as its underlying type.
//
// No explicit declaration that MyStruct is a Counter is needed or supported in
// Go. The compiler implicitly infers that MyStruct is convertible to Counter
// simply due to MyStruct having implemented all of Counter's methods, and
// nothing more nor less.
//
// This lack of "is a" relationships is pervasive in Go and intrinsic to its
// design. This is why Go is better suited to a functional paradigm than
// object-oriented approaches. In other words, interface types in Go are well
// suited for simple cross-cutting concerns like the fmt.Stringer and
// fmt.Scanner interfaces (which are designed as "mix-ins" rather than
// first-class types) but cannot be used directly for cases in which
// general-purpose polymorphism would be a useful feature.
func MakeCounter() Counter {
	v := new(int)
	s := MyStruct{v}

	// Note that unlike MyInt, a value of type MyStruct implements Counter
	// without any pointer receivers, so we do not use & here.
	return s
}

// Prints:
//
//	42
//	43
//	0
//	1
//
// to stdout.
func main() {

	var i MyInt = 42

	// Note that *MyInt is implicitly convertible to Counter because MyInt
	// implements all of Counter's methods and a pointer is required because
	// MyInt's implementation includes methods with pointer receivers. Compare
	// this to the MyStruct implementation which uses a wrapped *int to achieve
	// the same effect.
	var c Counter = &i

	fmt.Println(c.Value())
	c.Increment()
	fmt.Println(c.Value())

	var s Counter = MakeCounter()

	fmt.Println(s.Value())
	s.Increment()
	fmt.Println(s.Value())
}
