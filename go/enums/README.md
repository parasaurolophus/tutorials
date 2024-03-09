_Copyright &copy; Kirk Rader 2024_

# Enumerated Values (Not Types) in Go

Go does not support enumerated types comparable to `enum` in C++, C# or Java.
The code in [./enums.go](./enums.go) contains an example of the closest
approxiumation Go offers using named constants of application-specific types:

- Define an alias for `int`.
- Define constants of that type using Go's `iota` idiom.
- Implement the `fmt.Stringer` and `fmt.Scanner` interfaces to handle I/O using
  strings corresponding to the named constants.

While the preceding offers much of what is desired of a an `enum`, Go offers no
means of constraining your alias type to only contain values corresponding to
your named constants. A simple type cast from `int` can be used to create values
of your alias type that do not match any of your named constants and thus
violate the essential contract of `enum`.
