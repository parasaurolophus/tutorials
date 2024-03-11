// Copyright Kirk Rader 2024

package main

import (
	"fmt"
	"os"
)

// The built-in type named any is a synonym for interface{}, i.e. the empty
// interface.
//
// Its name implies the common use case of "a type of variable that can be bound
// to a value of any type." But that is not entirely accurate. Under the hood,
// values whose types are interfaces are implemented as references to the
// interface type and to the actual underlying value which implements it. So
// when you pass an int as the parameter, as in fn(42), value is bound not to an
// int whose value is 42, but to an interface value which points to an int whose
// value is 42. As with pointers and "." notation for methods, Go's syntax and
// runtime attempt to smooth over as many of these distinctions as they can but
// sometimes with somewhat inconsistent and confusing results.
//
// [As an aside, the officially stated reason that the type named any, which is
// a synonym for interface{}, can be bound to values of any type is due to the
// fact that Go does not support true polymorphism in the first place. Any given
// type, T, implements an interface, I, if and only if interface I declares no
// methods which type T fails to implement. Since interface{} declares no
// methods, it has none which any type fails to implement and so every type
// implements it. No, really.]
func fn(value any) {

	// Use a "type switch" statement to access the underlying value.
	//
	// You could also use an unconditional coercion such as value.(int) at the
	// risk of causing the runtime to panic if value is bound to the wrong type.
	switch v := value.(type) {

	case int:
		// If execution reaches here, v's type has been inferred to be int.
		fmt.Printf("%d (%T)\n", v, v)

	case nil:
		// If execution reaches here, value was bound to nil (which must be
		// accounted for as a possibility even when value is generally expected
		// to be bound to a type which is not, itself, nilable).
		fmt.Fprintf(os.Stderr, "expected an int, got nil\n")

	default:
		// Execution reaches here if no case clause matched. In that case we are no more
		fmt.Fprintf(os.Stderr, "expected an int, got %v of type %T\n", v, v)
	}
}

// Prints:
//
//	42 (int)
//
// to stdout.
func main() {

	fn(42)
}
