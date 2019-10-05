package collections

import (
	"testing"

	"github.com/martinohmann/collections-go/generic/immutable"
	"github.com/martinohmann/collections-go/internal"
	ctesting "github.com/martinohmann/collections-go/testing"
)

func TestEnsureCollectionMethods(t *testing.T) {
	ctesting.EnsureCollectionMethods(t, &immutable.Collection{}, nil)
	ctesting.EnsureCollectionMethods(t, &internal.Collection{}, []*internal.Type{})
	ctesting.EnsureCollectionMethods(t, &internal.ImmutableCollection{}, []*internal.Type{})
}
