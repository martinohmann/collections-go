package testutil

import (
	"testing"

	"github.com/martinohmann/collections-go/internal/validation"
)

// EnsureCollectionMethods tests if an object is a valid collection type with
// slice as its underlying storage type. An object is a valid collection if it
// has all methods defined in the internal validation package.
func EnsureCollectionMethods(t *testing.T, obj, slice interface{}) {
	err := validation.ValidateCollection(obj, slice)
	if err != nil {
		t.Fatal(err)
	}
}
