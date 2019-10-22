// Package types provides type definitions that collections-gen should generate
// collections for.
package types

//go:generate collections-gen -i github.com/martinohmann/collections-go/collections/types -p github.com/martinohmann/collections-go/collections

//nolint:deadcode
type (
	// +collections-gen=true
	// +collections-gen:options=mutable,nosuffix
	// +collections-gen:options=immutable,nosuffix
	b = byte

	// +collections-gen=true
	// +collections-gen:options=name=ByteSlice,equality-func=bytes.Equal
	// +collections-gen:options=immutable,name=ImmutableByteSlice,equality-func=bytes.Equal
	bs = []byte

	// +collections-gen=true
	// +collections-gen:options=mutable,nosuffix,out-name=float32
	// +collections-gen:options=immutable,nosuffix,out-name=immutable_float32
	f32 = float32

	// +collections-gen=true
	// +collections-gen:options=mutable,nosuffix,out-name=float64
	// +collections-gen:options=immutable,nosuffix,out-name=immutable_float64
	f64 = float64

	// +collections-gen=true
	// +collections-gen:options=mutable,nosuffix
	// +collections-gen:options=immutable,nosuffix
	i = int

	// +collections-gen=true
	// +collections-gen:options=mutable,nosuffix,out-name=int32
	// +collections-gen:options=immutable,nosuffix,out-name=immutable_int32
	i32 = int32

	// +collections-gen=true
	// +collections-gen:options=mutable,nosuffix,out-name=int64
	// +collections-gen:options=immutable,nosuffix,out-name=immutable_int64
	i64 = int64

	// +collections-gen=true
	// +collections-gen:options=mutable,nosuffix
	// +collections-gen:options=immutable,nosuffix
	s = string
)
