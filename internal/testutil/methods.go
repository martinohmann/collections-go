package testutil

import (
	"testing"

	"github.com/martinohmann/collections-go/internal/validation"
)

func EnsureCollectionMethods(t *testing.T, obj, slice interface{}) {
	err := validation.ValidateCollection(obj, slice)
	if err != nil {
		t.Fatal(err)
	}
}
