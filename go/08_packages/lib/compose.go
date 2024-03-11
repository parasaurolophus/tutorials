// Copyright Kirk Rader 2024

// Note that the package name is determined by the module name if there is a
// go.mod file in the same directory or else the directory name. In the latter
// case, there must be a go.mod file in some ancestor directory and this
// package's fully-qualified name will start with the module name and include
// intermediate names corresponding to the directory path from the module's base
// directory to this package's directory.
package lib

import (
	"fmt"
)

// Type of function that can be passed to Compose(any, ...Transfomer).
type Transformer func(value any) any

// Turn a specific type of function into a Transformer. The returned Transformer
// will panic if passed an argument of the wrong type.
func MakeTransformer[T any](trans func(t T) T) Transformer {

	return func(a any) any {
		return trans(a.(T))
	}
}

// Return the result of applying each of the given Transformer functions to the
// result of invoking the previous one, starting by passing the given initial
// value to the first function.
//
// Returns the final result and nil, or nil and an error value if any of the
// given Transformer functions cause a panic.
func Compose(value any, transformers ...Transformer) (any, error) {

	var recovered any = nil

	invoke := func(transformer Transformer) any {
		defer func() { recovered = recover() }()
		return transformer(value)
	}

	for _, transformer := range transformers {
		value = invoke(transformer)
		if recovered != nil {
			break
		}
	}

	if recovered == nil {
		return value, nil
	}

	return nil,
		fmt.Errorf(
			"Compose() recovered from a panic in a Transformer: %v",
			recovered)
}
