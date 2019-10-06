package customtype

//go:generate collections-gen -t *Type -n Collection -e Equal collection.go
//go:generate collections-gen -t *Type -n ImmutableCollection -e Equal -i immutable_collection.go

// Type is just a dummy type for demonstrating a custom type example.
type Type struct {
	Name string
}

// Equal is just a dummy func for demonstrating a custom equals func.
func Equal(a, b *Type) bool {
	return a == b
}
