// Copyright Kirk Rader 2024

package enums

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {

	actual := One.String()

	if actual != "One" {
		t.Errorf("expected \"One\", got \"%s\"", actual)
	}

	actual = Two.String()

	if actual != "Two" {
		t.Errorf("expected \"Two\", got \"%s\"", actual)
	}

	actual = Three.String()

	if actual != "Three" {
		t.Errorf("expected \"Three\", got \"%s\"", actual)
	}

	actual = MyEnum(100).String()

	if actual != "<unrecognized MyEnum: 100>" {
		t.Errorf("expected \"<unrecognized MyEnum: 100>\", got \"%s\"", actual)
	}
}

func TestScan(t *testing.T) {

	var one, two, three, invalid MyEnum

	n, err := fmt.Sscanf(" One, Two, Three ", " %v, %v, %v ", &one, &two, &three)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if n != 3 {
		t.Fatalf("expected 3, got %d", n)
	}

	if one != One {
		t.Errorf("expected One, got %v", one)
	}

	if two != Two {
		t.Errorf("expected Two, got %v", two)
	}

	if three != Three {
		t.Errorf("expected Three, got %v", three)
	}

	_, err = fmt.Sscanf(" Foo ", " %v ", &invalid)

	if err == nil {
		t.Errorf("expected an error")
	}
}
