// Copyright Kirk Rader 2024

package interfaces

type (

	// Define a simple interface.
	MutableInt interface {
		Get() int
		Set(n int)
	}

	// Define an alias for int capable of implementing MutableInt.
	mutableInt int

	// Define a private data type capable of implementing MutableInt.
	mutableIntStruct struct {
		n int
	}
)

// Implement MutableInt.Get() interface method by mutableIntStruct.
func (mi mutableIntStruct) Get() int {
	return mi.n
}

// Implement MutableInt.Set() interface method by mutableIntStruct.
//
// Note that a pointer is required in order to support methods which can modify
// the receiver's state.
// .     |
// .     V
func (mi *mutableIntStruct) Set(n int) {
	mi.n = n
}

// Implement MutableInt.Get() interface method by mutableInt.
func (mi mutableInt) Get() int {
	return int(mi)
}

// Implement MutableInt.Set() interface method by mutableInt
//
// Note that a pointer is required in order to support methods which can modify
// the receiver's state.
// .     |
// .     V
func (mi *mutableInt) Set(n int) {
	*mi = mutableInt(n)
}

// Constructor for a mutableIntStruct.
func MakeMutableIntStruct(n int) MutableInt {

	mi := new(mutableIntStruct)
	mi.n = n
	return mi
}

// Constructor for a mutableInt.
func MakeMutableInt(n int) MutableInt {

	mi := mutableInt(n)

	// Note that an address is required here because there is a method with a
	// pointer receiver.
	//     |
	//     V
	return &mi
}
