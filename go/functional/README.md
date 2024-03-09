_Copyright &copy; Kirk Rader 2024_

# Functional Programming in Go

Go's deliberate lack of support for polymorphism coupled with its implementation
of _lexical closures_ makes it far better suited for the functional programming
paradigm than, for example, object orientation. See
[./functional.go](./functional.go) for a "hello world" level of example of what
that can look like.

## Closures

Lexical closures (or just "closures," for brevity) are a core feature of
functional programming. The difference between a closure and, for example, a
function pointer in C is that closures represent both executable code and data.

> Not only is a closure itself a first-class data value, it also stores data in
> the form of the lexical environment it encloses (hence the name).

To understand what that means, here is some terminology used when discussing the
functional programming paradigm:

- Data _values_ exist independently of any given set of _variable bindings_.

- A _variable_ represents a potential _binding_ between an identifier and a value.

- A given variable may be _free_ or _bound_ from the point of view of any given
  _lexical scope_, e.g. some function's "body."

- _Variable binding operators_ create new bindings that are only visible within
  the lexical scope defined by that operator.

- A _lexical environment_ (or just "environment") is a set of bindings between a
  particular set of variables and values that is accessible within the scope of
  a given variable binding operator.

The `func` keyword is Go's variable binding operator. It is functionally (pun
intended) identical to Lisp's `lambda` special form, Ruby's `{|x, y, ...| ...}`
construct and similar features of countless other languages. Not only can `func`
be used to define functions to be called as top-level forms, it can also be used
to define closures. Closures are both data values in themselves but also enclose
the environment in which they were created, i.e. they retain the particular set
of bindings for all the lexical variables active at the moment of their
creation. In other words, when a closure is applied to some set of parameters,
the code it executes has access to all of the variables that were bound from the
point of view of the `func` definition that was used to create that closure,
each time the closure is called.

Such lexical scopes can be nested. I.e. a variable that is free in the body of a
`func` might be bound in the environment that is active where the `func` special
form is evaluated. Such "outer" variable bindings are included in the
environments enclosed by each closure created within the environment of whatever
variable binding operator caused them to be bound in the first place. All
closures created within a given lexical environment share the same set of
bindings for variables that are free from the point of the `func` body but bound
in an enclosing environment -- a feature of closures that makes the particuarly
powerful and enables them to be not just as an alternative object-orientation,
but as an underlying implementation of object-oriented programming constructs.

```go
package main

import (
	"fmt"
	"os"
)

// Return two closures which function like the "getter" and "setter" methods
// for some object with an int "field."
func MakeClosures(a int) (func() int, func(int)) {

    // Return the current value of `a` each time `getter` is called.
	getter := func() int { return a }

	// Update the currently bound value of `a` to `b` each time `setter`
	// is called.
	setter := func(b int) { a = b }

	// Because `a` is a free variable from the point of view of both
	// `getter` and `setter`, they both include the same binding for `a`
	// in their enclosed environments.
	return getter, setter
}

// Prints
//
//	100  42
//
// to stdout
func main() {

	getter1, setter1 := MakeClosures(0)
	getter2, _ := MakeClosures(42)
	n1 := getter1()
	n2 := getter2()

	if n1 != 0 {
		fmt.Fprintf(os.Stderr, "expected 0, got %d\n", n1)
	}

	if n2 != 42 {
		fmt.Fprintf(os.Stderr, "expected 42, got %d\n", n1)
	}

	setter1(100)
	n1 = getter1()
	n2 = getter2()

	if n1 != 100 {
		fmt.Fprintf(os.Stderr, "expected 100, got %d\n", n1)
	}

	if n2 != 42 {
		fmt.Fprintf(os.Stderr, "expected 42, got %d\n", n1)
	}

	fmt.Printf("%3d %3d\n", n1, n2)
}
```

Looking at the implementation and output for the preceding example, we can see that:

- All closures created by `MakeClosures(n)` share the single binding for `n` that
  is active when they are created.

- The set of closures created by `MakeClosures(n)` retains a distinct binding
  for `n` each time it is invoked.

All of this enables some incredibly powerful idioms. Anything you can do with
objects can be done with closures as "methods" for "instances" whose state is
represented by the lexical environments they enclose.

> It is an ancient Lisp programmers' maxim that "objects are a poor man's
> closures." Evidence for this can be found in the fact that the world's first
> commercially significant object-oriented programming system --
> [Flavors](https://en.wikipedia.org/wiki/Flavors_(programming_language)) on the
> [Symbolics Lisp Machine](https://en.wikipedia.org/wiki/Symbolics) -- was based
> on closures as its underlying implementation mechanism in a style that
> survives to this day in the form of
> [CLOS](https://en.wikipedia.org/wiki/Common_Lisp_Object_System), the Common
> Lisp Object System.

While you can use closures to support an object-oriented paradigm, e.g. as a
work-around for Go's lack of native support for polymorphism, they are more
general purpose and flexible than the objects in any C++, C# or Java style
object-oriented language. Go falls short of supporting a true
[Continuation-Passing
Style](https://en.wikipedia.org/wiki/Continuation-passing_style) (CPS) due to
its lack of tail-call optimization and first-class continuations, but closures
by themselves are incredibly powerful as building blocks for algorithms. As
always, with great power comes even greater responsibility.

> You can, of course, use Go's closures to implement specific CPS-inspired
> idioms, as demonstrated by the `Compose()` function in
> [./functional.go](./functional.go).
