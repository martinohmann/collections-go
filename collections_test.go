package collections

import (
	"testing"

	"github.com/martinohmann/collections-go/generic/immutable"
	"github.com/martinohmann/collections-go/internal"
	"github.com/martinohmann/collections-go/internal/testutil"
)

func TestEnsureCollectionMethods(t *testing.T) {
	testutil.EnsureCollectionMethods(t, &immutable.Collection{}, nil)
	testutil.EnsureCollectionMethods(t, &internal.Collection{}, []*internal.Type{})
	testutil.EnsureCollectionMethods(t, &internal.ImmutableCollection{}, []*internal.Type{})
}
