// Copyright Kirk Rader 2024

package lib

import (
	"testing"
)

func TestCompose(t *testing.T) {

	add := MakeTransformer(func(value int) int { return value + 1 })
	sub := MakeTransformer(func(value int) int { return value - 1 })

	result, err := Compose(0, add, add, sub, add)

	if err != nil {
		t.Errorf(err.Error())
	}

	switch v := result.(type) {

	case int:

		if v != 2 {
			t.Errorf("expected 2, got %d", v)
		}

	default:
		t.Errorf("expected an int, got %v of type %T", v, v)
	}
}

func TestComposePanic(t *testing.T) {

	count := 0

	add := MakeTransformer(func(value int) int {
		count += 1
		return value + 1
	})

	sub := MakeTransformer(func(value float64) float64 {
		count += 1
		return value - 1.0
	})

	result, err := Compose(0, add, add, sub, add)

	if err == nil {
		t.Errorf("expected an error")
	}

	if result != nil {
		t.Errorf("expected result to be nil, got %v of type %T", result, result)
	}

	if count != 2 {
		t.Errorf("expected transformer1 to have been called twice, got %d", count)
	}
}
