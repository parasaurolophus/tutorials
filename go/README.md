_Copyright &copy; Kirk Rader 2024_

# Examples and Tutorials for the Go Programming Language

- [Basics](./01_basics/)
- [Pointers](./02_pointers/)
- [Methods](./03_methods/)
- [Interfaces](./04_interfaces/)
- [Generics](./05_generics/)
- [Closures](./06_closures/)
- [`any`](./07_any/)
- [Packages](./08_packages/)
- [Enums (More or Less)](./09_enums/)
- [Concurrency](./10_concurrency/)

## `parasaurolophus/tutorial` Module Structure

The files and directory structure in this repository represents an example of
organizing packages within a module.

```
- tuorials/
  |
  +- go/
  |
  +- go.mod (defines module named "parasaurolophus/tutorial")
  |
  +- 01_basics/
  |  |
  |  +- basics.go (standalone program with a `main()` in `main` package)
  |
  +- 02_pointers/
  |  |
  |  +- pointers.go  (standalone program with a `main()` in `main` package)
  |
  +- 03_methods/
  |  |
  |  +- methods.go (standalone program with a `main()` in `main` package)
  |
  +- 04_interfaces/
  |  |
  |  +- interfaces.go (standalone program with a `main()` in `main` package)
  |
  +- 05_generics/
  |  |
  |  +- generics.go (standalone program with a `main()` in `main` package)
  |
  +- 06_closures/
  |  |
  |  +- closures.go (standalone program with a `main()` in `main` package)
  |
  +- 07_any/
  |  |
  |  +- any.go (standalone program with a `main()` in `main` package)
  |
  +- 08_packages/
  |  |
  |  +- packages.go (standalone program with a `main()` in `main` package)
  |  |
  |  +- lib/
  |     |
  |     +- compose.go (library code in package `parasaurolophus/tutorial/08_packages/lib`)
  |     |
  |     +- compose_test.go (unit tests for the library code)
  |
  +- 09_enums/
  |  |
  |  +- enums.go (standalone program with a `main()` in `main` package)
  |
  +- 10_concurrency/
     |
     +- concurrency.go (standalone program with a `main()` in `main` package)
```
