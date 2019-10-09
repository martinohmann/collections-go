package validation

import (
	"reflect"
)

// ParameterType defines the type of function parameter in the MethodMap.
// Usually types are defined using reflection kinds, but for types that are
// specific to each collection (e.g. the type of the collection itself, its
// slice and element type) this is not possible.
type ParameterType int

const (
	// ReflectionType denotes that a parameter should just be checked using its
	// reflection kind.
	ReflectionType ParameterType = iota

	// CollectionType denotes that the parameter must match the type of the
	// collection.
	CollectionType

	// SliceType denotes that the parameter must match the type of the
	// collection's underlying slice.
	SliceType

	// ElementType denotes that the parameter must match the type of the
	// collection's element type.
	ElementType
)

// Parameter describes a function parameter.
type Parameter struct {
	Type          ParameterType
	Kind          reflect.Kind
	FuncSignature Signature
}

// Signature describes the signature of a function.
type Signature struct {
	NumIn    int
	NumOut   int
	In       []Parameter
	Out      []Parameter
	Variadic bool
}

// MethodMap contains the names of all expected collection methods together
// with their function signature.
var MethodMap = map[string]Signature{
	"All": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Kind: reflect.Bool,
			},
		},
	},
	"Any": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Kind: reflect.Bool,
			},
		},
	},
	"Append": {
		Variadic: true,
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Slice,
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Cap": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
		},
		Out: []Parameter{
			{
				Kind: reflect.Int,
			},
		},
	},
	"Collect": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Contains": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Type: ElementType,
			},
		},
		Out: []Parameter{
			{
				Kind: reflect.Bool,
			},
		},
	},
	"Copy": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Cut": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Int,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				Type: SliceType,
			},
		},
	},
	"Each": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
				},
			},
		},
	},
	"EachIndex": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
						{
							Kind: reflect.Int,
						},
					},
				},
			},
		},
	},
	"Filter": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Find": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: ElementType,
			},
		},
	},
	"First": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
		},
		Out: []Parameter{
			{
				Type: ElementType,
			},
		},
	},
	"FirstN": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				Type: SliceType,
			},
		},
	},
	"Get": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				Type: ElementType,
			},
		},
	},
	"IndexOf": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Type: ElementType,
			},
		},
		Out: []Parameter{
			{
				Kind: reflect.Int,
			},
		},
	},
	"InsertItem": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Type: ElementType,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"IsSorted": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Kind: reflect.Bool,
			},
		},
	},
	"Items": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
		},
		Out: []Parameter{
			{
				Type: SliceType,
			},
		},
	},
	"Last": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
		},
		Out: []Parameter{
			{
				Type: ElementType,
			},
		},
	},
	"LastN": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				Type: SliceType,
			},
		},
	},
	"Len": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
		},
		Out: []Parameter{
			{
				Kind: reflect.Int,
			},
		},
	},
	"Map": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Type: ElementType,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"MapIndex": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
						{
							Kind: reflect.Int,
						},
					},
					Out: []Parameter{
						{
							Type: ElementType,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Nth": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				Type: ElementType,
			},
		},
	},
	"Partition": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Type: CollectionType,
			},
		},
	},
	"Prepend": {
		Variadic: true,
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Slice,
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Reduce": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Type: ElementType,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: ElementType,
			},
		},
	},
	"Reject": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Remove": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"RemoveItem": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Type: ElementType,
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Reverse": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
	"Slice": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Int,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				Type: SliceType,
			},
		},
	},
	"Sort": {
		In: []Parameter{
			{
				Type: CollectionType,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							Type: ElementType,
						},
						{
							Type: ElementType,
						},
					},
					Out: []Parameter{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				Type: CollectionType,
			},
		},
	},
}
