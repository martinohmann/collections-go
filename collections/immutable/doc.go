package immutable

//go:generate collections-gen --immutable --item-type=int int.go
//go:generate collections-gen --immutable --item-type=int64 int64.go
//go:generate collections-gen --immutable --item-type=bool bool.go
//go:generate collections-gen --immutable --item-type=float32 float32.go
//go:generate collections-gen --immutable --item-type=float64 float64.go
//go:generate collections-gen --immutable --item-type=string string.go
//go:generate collections-gen --immutable --item-type=[]byte --name=ByteSliceCollection --equality-func=bytes.Equal --imports=bytes byte_slice.go
