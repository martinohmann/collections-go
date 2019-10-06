package collections

// Mutable collections.

//go:generate collections-gen -t int -n Int int.go
//go:generate collections-gen -t int32 -n Int32 int32.go
//go:generate collections-gen -t int64 -n Int64 int64.go
//go:generate collections-gen -t float32 -n Float32 float32.go
//go:generate collections-gen -t float64 -n Float64 float64.go
//go:generate collections-gen -t string -n String string.go
//go:generate collections-gen -t byte -n Byte byte.go
//go:generate collections-gen -t []byte -n ByteSlice -e bytes.Equal -I bytes byte_slice.go

// Immutable collections.

//go:generate collections-gen -i -t int -n ImmutableInt immutable_int.go
//go:generate collections-gen -i -t int32 -n ImmutableInt32 immutable_int32.go
//go:generate collections-gen -i -t int64 -n ImmutableInt64 immutable_int64.go
//go:generate collections-gen -i -t float32 -n ImmutableFloat32 immutable_float32.go
//go:generate collections-gen -i -t float64 -n ImmutableFloat64 immutable_float64.go
//go:generate collections-gen -i -t string -n ImmutableString immutable_string.go
//go:generate collections-gen -i -t byte -n ImmutableByte immutable_byte.go
//go:generate collections-gen -i -t []byte -n ImmutableByteSlice -e bytes.Equal -I bytes immutable_byte_slice.go
