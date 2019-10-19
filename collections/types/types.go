// Package types provides type definitions that collection-gen should generate
// collections for.
package types

//go:generate collection-gen -i github.com/martinohmann/collections-go/collections/types -p github.com/martinohmann/collections-go/collections

//nolint:deadcode
type (
	// +collection-gen=true
	// +collection-gen:options=mutable,nosuffix
	// +collection-gen:options=immutable,nosuffix
	b = byte

	// +collection-gen=true
	// +collection-gen:options=name=ByteSlice,equality-func=bytes.Equal
	// +collection-gen:options=immutable,name=ImmutableByteSlice,equality-func=bytes.Equal
	bs = []byte

	// +collection-gen=true
	// +collection-gen:options=mutable,nosuffix,out-name=float32
	// +collection-gen:options=immutable,nosuffix,out-name=immutable_float32
	f32 = float32

	// +collection-gen=true
	// +collection-gen:options=mutable,nosuffix,out-name=float64
	// +collection-gen:options=immutable,nosuffix,out-name=immutable_float64
	f64 = float64

	// +collection-gen=true
	// +collection-gen:options=mutable,nosuffix
	// +collection-gen:options=immutable,nosuffix
	i = int

	// +collection-gen=true
	// +collection-gen:options=mutable,nosuffix,out-name=int32
	// +collection-gen:options=immutable,nosuffix,out-name=immutable_int32
	i32 = int32

	// +collection-gen=true
	// +collection-gen:options=mutable,nosuffix,out-name=int64
	// +collection-gen:options=immutable,nosuffix,out-name=immutable_int64
	i64 = int64

	// +collection-gen=true
	// +collection-gen:options=mutable,nosuffix
	// +collection-gen:options=immutable,nosuffix
	s = string
)
