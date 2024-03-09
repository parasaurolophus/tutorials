_Copyright &copy; Kirk Rader 2024_

# Interfaces in Go

The [../enums/enums.go](../enums/enums.go) an example of adding methods to an
application-specific type to implement methods defined by existing interfaces
(`fmt.Stringer` and `fmt.Scanner`, in that example). The code in
[./interfaces.go](./interfaces.go) demonstrates how to declare your own
interfaces and implement them for more than one type. The features of Go's
version of interfaces (remember: "a feature is a bug as described by the
marketing department") to note:

- There is no way to declare that a given type implements a given set of
  interfaces.

- A corollary is that in Go, a given type implements a given interface if and
  only it implements all of that interface's methods.

- Another corollary is that the compiler often provides less help than one might
  hope or expect when it comes to knowing exactly what is missing or incorrect
  when it complains that a given type can't be used as an instance of a given
  interface.

The preceding issues arise because Go eschews any notion of "is a" relationships
between types, i.e. it does everything possible to minimize support for
polymorphism. Since polymorphism is the entire _raison d'etre_ for interfaces,
Go's interfaces are "special" in many ways. This is especially true for
interfaces whose semantics include methods that must modify the state of the
receiver. See the inline comments in the example code for more information.

For all these reasons, you may find that Go is better suited to a
[../functional/](../functional/) programming paradigm for things like dependency
injection and other inversions of control than relying on interfaces. Interfaces
in Go are mainly useful to implement cross-cutting concerns like the `fmt`
package's interface methods demonstrated for [enumerated values](../enums/).

## Type Identity and the `any` Type

All of that said, Go implements interface types in a particular way that many of
its core idioms exploit. Because Go does not support true polymorphism,
assignment of a value to a variable whose type is some interface always implies
an implicit type conversion. Inside the Go runtime, interface values are
indirect references to the values that provide the implementation. Those
references can be uninitialized or erased. That is why interface values are
always nilable even when the underlying value type is not:

```go
package main

import (
	"fmt"
	"os"
)

type (
	MyInt int

	MyInterface interface {
		MyMethod() int
	}
)

func (mi MyInt) MyMethod() int {
	return int(mi)
}

func MyFunc() (MyInterface, error) {
	mi := MyInt(0)
	return mi, nil
}

func main() {

	n := MyInt(0)
	fmt.Printf("n is %d\n", n)

	// Syntax error because n is not nilable
	//      |
	//      V
	// if n == nil {
	//   ....
	// }

	i, err := MyFunc()

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	if i == nil {
		// i is nilable even though MyFunc() returns a MyInt
		fmt.Println("i is nil")
	}
}
```

The preceding "feature" of interfaces is what enables the ubiquitous idiom for
error reporting and handling in Go.

```go
if err := someFunc(); err != nil {
  // handle the error...
}
```

In addition, Go provides an escape hatch for its enthusiasticly strict type
enforcement: the `any` type. Inside the Go runtime, `any` is actually just an
alias for the type more properly known as `interface{}`, i.e. "the empty
interface."

> The reason the empty interface can be used as a type for values of any other
> type has to do with the "features" of interfaces described previously. A given
> type, `T`, will be defined as implementing an interface, `I`, so long as `I`
> has no methods which `T` fails to implement. Since `interface{}` has no
> methods in the first place, it has no methods which _any_ type fails to
> implement. So every type implements `any`.

Values of every type can be cast to `any`. Go's general antipathy for
polymorphism makes the `any` type less useful than you might imagine at first.
But it can be exploited like other interface types for cross-cutting concerns
like logging and other kinds of I/O.
