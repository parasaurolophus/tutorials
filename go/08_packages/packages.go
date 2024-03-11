// Copyright Kirk Rader 2024

// The package named main is treated specially by the Go build tools and
// runtime. It is used as the package for the main() function, i.e. the entry
// point function for stand-alone executables.
package main

import (

	// These imports reference the standard library.
	"fmt"
	"os"

	// Import types and functions from the package in the lib subdirectory.
	//
	// In particular, the fully qualified package name starts with
	// parasaurolophus/tutorial because that is the module name in ../go.mod and
	// ends with 08_packages/lib because that is the subdirectory path from the
	// directory containing go.mod to the one containing the .go files being
	// referenced. Each of them must declare that they are in a package name lib
	// because that is the name of the directory in which they reside.
	//
	// The same would apply to packages imported from an entirely different
	// module, but then those packages' contents would have to be made available
	// in a manner compatible with Go's package publishing mechanisms.
	"parasaurolophus/tutorial/08_packages/lib"
)

// Prints 2 to stdout.
//
// All of the other examples in this repository operate exclusively in package
// main. This example includes type and function definitions from a package in a
// subdirectory named lib.
func main() {

	// Invoke functions imported from the lib package. Note that while the
	// import statement uses the fully qualified package name, only the leaf
	// package name is used to reference symbols. I.e. while Go's packages exist
	// in a hierarchical namespace, the way that they are actually used causes
	// them to be flattened into only their leaf package names. Go's import
	// statement supports a package aliasing feature to allow for resulting name
	// collisions (i.e. cases where more than one package ends in the same leaf
	// package name).
	add := lib.MakeTransformer(func(value int) int { return value + 1 })
	sub := lib.MakeTransformer(func(value int) int { return value - 1 })
	result, err := lib.Compose(0, add, add, sub, add)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Println(result)
}
