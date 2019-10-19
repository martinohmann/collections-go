package customtype

//go:generate collection-gen -i github.com/martinohmann/collections-go/examples/custom-type -p github.com/martinohmann/collections-go/examples/custom-type

// +collection-gen=true
// +collection-gen:options=pointer,equality-func=Equal
// +collection-gen:options=pointer,immutable,equality-func=Equal

// Type is just a dummy type for demonstrating a custom type example.
type Type struct {
	Name string
}

// Equal is just a dummy func for demonstrating a custom equals func.
func Equal(a, b *Type) bool {
	return a == b
}
