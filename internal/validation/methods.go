package validation

import (
	"reflect"
)

type Parameter struct {
	Type           reflect.Type
	Kind           reflect.Kind
	FuncSignature  Signature
	ElementType    bool
	CollectionType bool
	SliceType      bool
}

type Signature struct {
	NumIn    int
	NumOut   int
	In       []Parameter
	Out      []Parameter
	Variadic bool
}

var MethodMap = map[string]Signature{
	"All": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
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
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
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
				CollectionType: true,
			},
			{
				Kind: reflect.Slice,
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"Cap": {
		In: []Parameter{
			{
				CollectionType: true,
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
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
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
				CollectionType: true,
			},
		},
	},
	"Contains": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				ElementType: true,
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
				CollectionType: true,
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"Cut": {
		In: []Parameter{
			{
				CollectionType: true,
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
				SliceType: true,
			},
		},
	},
	"Each": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
						},
					},
				},
			},
		},
	},
	"EachIndex": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
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
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
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
				CollectionType: true,
			},
		},
	},
	"Find": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
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
				ElementType: true,
			},
		},
	},
	"First": {
		In: []Parameter{
			{
				CollectionType: true,
			},
		},
		Out: []Parameter{
			{
				ElementType: true,
			},
		},
	},
	"FirstN": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				SliceType: true,
			},
		},
	},
	"Get": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				ElementType: true,
			},
		},
	},
	"IndexOf": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				ElementType: true,
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
				CollectionType: true,
			},
			{
				ElementType: true,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"IsSorted": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
						},
						{
							ElementType: true,
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
				CollectionType: true,
			},
		},
		Out: []Parameter{
			{
				SliceType: true,
			},
		},
	},
	"Last": {
		In: []Parameter{
			{
				CollectionType: true,
			},
		},
		Out: []Parameter{
			{
				ElementType: true,
			},
		},
	},
	"LastN": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				SliceType: true,
			},
		},
	},
	"Len": {
		In: []Parameter{
			{
				CollectionType: true,
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
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
						},
					},
					Out: []Parameter{
						{
							ElementType: true,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"MapIndex": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
						},
						{
							Kind: reflect.Int,
						},
					},
					Out: []Parameter{
						{
							ElementType: true,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"Nth": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				ElementType: true,
			},
		},
	},
	"Partition": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
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
				CollectionType: true,
			},
			{
				CollectionType: true,
			},
		},
	},
	"Prepend": {
		Variadic: true,
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Slice,
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"Reduce": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
						},
						{
							ElementType: true,
						},
					},
					Out: []Parameter{
						{
							ElementType: true,
						},
					},
				},
			},
		},
		Out: []Parameter{
			{
				ElementType: true,
			},
		},
	},
	"Reject": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
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
				CollectionType: true,
			},
		},
	},
	"Remove": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"RemoveItem": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				ElementType: true,
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"Reverse": {
		In: []Parameter{
			{
				CollectionType: true,
			},
		},
		Out: []Parameter{
			{
				CollectionType: true,
			},
		},
	},
	"Slice": {
		In: []Parameter{
			{
				CollectionType: true,
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
				SliceType: true,
			},
		},
	},
	"Sort": {
		In: []Parameter{
			{
				CollectionType: true,
			},
			{
				Kind: reflect.Func,
				FuncSignature: Signature{
					In: []Parameter{
						{
							ElementType: true,
						},
						{
							ElementType: true,
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
				CollectionType: true,
			},
		},
	},
}
