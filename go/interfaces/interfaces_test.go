// Copyright Kirk Rader 2024

package interfaces

import "testing"

func TestStructGet(t *testing.T) {

	mi := MakeMutableIntStruct(42)
	n := mi.Get()

	if n != 42 {
		t.Errorf("expected 42, got %d", n)
	}
}

func TestStructSet(t *testing.T) {

	mi := MakeMutableIntStruct(0)
	mi.Set(42)
	n := mi.Get()

	if n != 42 {
		t.Errorf("expected 42, got %d", n)
	}
}

func TestIntGet(t *testing.T) {

	mi := MakeMutableInt(42)
	n := mi.Get()

	if n != 42 {
		t.Errorf("expected 42, got %d", n)
	}
}

func TestIntSet(t *testing.T) {

	mi := MakeMutableInt(0)
	mi.Set(42)
	n := mi.Get()

	if n != 42 {
		t.Errorf("expected 42, got %d", n)
	}
}
