// Copyright Kirk Rader 2024

package main

import "fmt"

// You can only define methods in the same package as the definition of the
// receiver's type. `MyInt` is an application-specific alias that allows us to
// define methods on what is in most ways really just an `int`.
type MyInt int

// A method is a function with a receiver. Methods are invoked using "."
// notation rather than using normal function-calling syntax. I.e. Go's
// "receiver" syntax is its equivalent of the implicit "this" variable defined
// by default for methods in languages like C++, C#, Java etc.
func (mi MyInt) AddOne() MyInt {

	// Go's strict type enforcement is not actually all that strict. Numeric
	// constants like 0, 1, 2, 2.0, 3.0 can be used interchangeably for any
	// compatible numeric type as a "convenience" for not always having to write
	// stuff like mi + MyInt(1). It therefore can be confusing when, exactly,
	// you do and do not need an explicit type cast and the compiler diagnostics
	// are not always as clearly written as one might hope.
	return mi + 1
}

// A method with a pointer receiver can alter its receiver's state.
func (mi *MyInt) Increment() {

	// Pointers must (usually) be derefernced in order to access their values.
	*mi += 1
}

// Prints:
//
//	1
//	0
//	1
//
// to stdout.
func main() {

	// Explicit cast so that the type of `mi` will be inferred to be `MyInt`
	// rather than `int`. This could have been written:
	//
	//  var mi MyInt = 0
	//
	// which form is "better" is a matter of taste.
	mi := MyInt(0)

	fmt.Println(mi.AddOne())
	fmt.Println(mi)

	// Invoking a method with "." notation implicitly adds the necessary
	// "pointer to" operation (another case of a convenience that can lead to
	// rather tortured syntax which we will get to when discussing interfaces).
	mi.Increment()
	fmt.Println(mi)
}
