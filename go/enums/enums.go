// Copyright Kirk Rader 2024

package enums

import (
	"fmt"
	"unicode"
)

// Int alias for enum type
type MyEnum int

// Define three enumerated values for MyEnum using iota
const (
	One MyEnum = iota
	Two
	Three
)

// Implement fmt.Stringer interface
func (e MyEnum) String() string {

	switch e {

	case One:
		return "One"

	case Two:
		return "Two"

	case Three:
		return "Three"

	default:
		return fmt.Sprintf("<unrecognized MyEnum: %d>", int(e))
	}
}

// Implement the fmt.Scanner interface
func (e *MyEnum) Scan(state fmt.ScanState, _ rune) error {

	b, err := state.Token(true, unicode.IsLetter)

	if err != nil {
		return err
	}

	token := string(b)
	switch token {

	case One.String():
		*e = One

	case Two.String():
		*e = Two

	case Three.String():
		*e = Three

	default:
		return fmt.Errorf("unsupported MyEnum token: \"%s\"", token)
	}

	return nil
}
