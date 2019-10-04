package internal

//go:generate go run ../cmd/collections-gen/main.go --item-type=*Type --name=Collection collection.go
//go:generate go run ../cmd/collections-gen/main.go --item-type=*Type --name=ImmutableCollection --immutable immutable_collection.go

// Type is just a dummy type for validating the code generator
type Type struct {
	Name string
}
