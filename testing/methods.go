package testing

import (
	"reflect"
	"testing"

	"github.com/martinohmann/collections-go/internal"
)

func EnsureCollectionMethods(t *testing.T, obj, slice interface{}) {
	objType := reflect.TypeOf(obj)
	sliceType := reflect.TypeOf(slice)

	err := internal.EnsureCollectionMethods(objType, sliceType)
	if err != nil {
		t.Fatal(err)
	}
}
