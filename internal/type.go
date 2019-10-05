package internal

//go:generate go run ../cmd/collections-gen/main.go -t *Type -n Collection collection.go
//go:generate go run ../cmd/collections-gen/main.go -t *Type -n ImmutableCollection -i immutable_collection.go

// Type is just a dummy type for validating the code generator
type Type struct {
	Name string
}
