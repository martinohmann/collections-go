package collections

// Interface is the smallest common interface shared by all collection types.
type Interface interface {
	// Len returns the length of the collection's underlying slice.
	Len() int

	// Cap returns the capacity of the collection's underlying slice.
	Cap() int

	// Interface returns the collections's underlying slice as an interface{}.
	Interface() interface{}
}
