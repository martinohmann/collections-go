package collections

import (
	"reflect"
	"testing"

	"github.com/martinohmann/collections-go/generic/immutable"
	"github.com/martinohmann/collections-go/internal"
	collectiontesting "github.com/martinohmann/collections-go/testing"
)

func TestEnsureCollectionMethods(t *testing.T) {
	collectiontesting.EnsureCollectionMethods(
		t,
		reflect.TypeOf((*immutable.Collection)(nil)),
		reflect.TypeOf((*interface{})(nil)).Elem(),
	)
	collectiontesting.EnsureCollectionMethods(
		t,
		reflect.TypeOf(&internal.Collection{}),
		reflect.TypeOf([]*internal.Type{}),
	)
	collectiontesting.EnsureCollectionMethods(
		t,
		reflect.TypeOf(&internal.ImmutableCollection{}),
		reflect.TypeOf([]*internal.Type{}),
	)
}
