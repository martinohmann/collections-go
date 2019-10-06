package collections

import (
	"testing"

	"github.com/martinohmann/collections-go/internal/testutil"
)

func TestEnsureCollectionMethods(t *testing.T) {
	testutil.EnsureCollectionMethods(t, &ImmutableGeneric{}, nil)
}
