package collections

import (
	"testing"

	"github.com/martinohmann/collections-go/internal/testutil"
)

func TestEnsureCollectionMethods(t *testing.T) {
	// Mutable collections
	testutil.EnsureCollectionMethods(t, &Generic{}, nil)
	testutil.EnsureCollectionMethods(t, &Byte{}, ([]byte)(nil))
	testutil.EnsureCollectionMethods(t, &ByteSlice{}, ([][]byte)(nil))
	testutil.EnsureCollectionMethods(t, &Float32{}, ([]float32)(nil))
	testutil.EnsureCollectionMethods(t, &Float64{}, ([]float64)(nil))
	testutil.EnsureCollectionMethods(t, &Int{}, ([]int)(nil))
	testutil.EnsureCollectionMethods(t, &Int32{}, ([]int32)(nil))
	testutil.EnsureCollectionMethods(t, &Int64{}, ([]int64)(nil))
	testutil.EnsureCollectionMethods(t, &String{}, ([]string)(nil))

	// Immutable collections
	testutil.EnsureCollectionMethods(t, &ImmutableGeneric{}, nil)
	testutil.EnsureCollectionMethods(t, &ImmutableByte{}, ([]byte)(nil))
	testutil.EnsureCollectionMethods(t, &ImmutableByteSlice{}, ([][]byte)(nil))
	testutil.EnsureCollectionMethods(t, &ImmutableFloat32{}, ([]float32)(nil))
	testutil.EnsureCollectionMethods(t, &ImmutableFloat64{}, ([]float64)(nil))
	testutil.EnsureCollectionMethods(t, &ImmutableInt{}, ([]int)(nil))
	testutil.EnsureCollectionMethods(t, &ImmutableInt32{}, ([]int32)(nil))
	testutil.EnsureCollectionMethods(t, &ImmutableInt64{}, ([]int64)(nil))
	testutil.EnsureCollectionMethods(t, &ImmutableString{}, ([]string)(nil))
}
