package validation

import (
	"reflect"
)

// MethodMap contains the names of all expected collection methods together
// with their function signature.
var MethodMap = map[string]Signature{
	"All": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Kind: reflect.Bool,
			},
		},
	},
	"Any": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Kind: reflect.Bool,
			},
		},
	},
	"Append": {
		Variadic: true,
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Slice,
				Name: "items",
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Cap": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
		},
		Out: Parameters{
			{
				Kind: reflect.Int,
			},
		},
	},
	"Collect": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Contains": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Type: ElementType,
				Name: "item",
			},
		},
		Out: Parameters{
			{
				Kind: reflect.Bool,
			},
		},
	},
	"Copy": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Cut": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Int,
				Name: "i",
			},
			{
				Kind: reflect.Int,
				Name: "j",
			},
		},
		Out: Parameters{
			{
				Type: SliceType,
			},
		},
	},
	"Each": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
				},
			},
		},
	},
	"EachIndex": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
						{
							Kind: reflect.Int,
							Name: "i",
						},
					},
				},
			},
		},
	},
	"Filter": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Find": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Type: ElementType,
			},
		},
	},
	"First": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
		},
		Out: Parameters{
			{
				Type: ElementType,
			},
		},
	},
	"FirstN": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Int,
				Name: "n",
			},
		},
		Out: Parameters{
			{
				Type: SliceType,
			},
		},
	},
	"Get": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Int,
			},
		},
		Out: Parameters{
			{
				Type: ElementType,
			},
		},
	},
	"IndexOf": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Type: ElementType,
				Name: "item",
			},
		},
		Out: Parameters{
			{
				Kind: reflect.Int,
			},
		},
	},
	"InsertItem": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Type: ElementType,
				Name: "item",
			},
			{
				Kind: reflect.Int,
				Name: "pos",
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Interface": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
		},
		Out: Parameters{
			{
				Kind: reflect.Interface,
			},
		},
	},
	"IsSorted": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "a",
						},
						{
							Name: "b",
							Type: ElementType,
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Kind: reflect.Bool,
			},
		},
	},
	"Items": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
		},
		Out: Parameters{
			{
				Type: SliceType,
			},
		},
	},
	"Last": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
		},
		Out: Parameters{
			{
				Type: ElementType,
			},
		},
	},
	"LastN": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Int,
				Name: "n",
			},
		},
		Out: Parameters{
			{
				Type: SliceType,
			},
		},
	},
	"Len": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
		},
		Out: Parameters{
			{
				Kind: reflect.Int,
			},
		},
	},
	"Map": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Type: ElementType,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"MapIndex": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
						{
							Kind: reflect.Int,
							Name: "i",
						},
					},
					Out: Parameters{
						{
							Type: ElementType,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Nth": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Int,
				Name: "n",
			},
		},
		Out: Parameters{
			{
				Type: ElementType,
			},
		},
	},
	"Partition": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
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
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Slice,
				Name: "items",
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Reduce": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "reducer",
						},
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Type: ElementType,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Type: ElementType,
			},
		},
	},
	"Reject": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "item",
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Remove": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Int,
				Name: "i",
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"RemoveItem": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Type: ElementType,
				Name: "item",
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Reverse": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
	"Slice": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Int,
				Name: "i",
			},
			{
				Kind: reflect.Int,
				Name: "j",
			},
		},
		Out: Parameters{
			{
				Type: SliceType,
			},
		},
	},
	"Sort": {
		In: Parameters{
			{
				Type: CollectionType,
				Name: "c",
			},
			{
				Kind: reflect.Func,
				Name: "fn",
				FuncSignature: Signature{
					In: Parameters{
						{
							Type: ElementType,
							Name: "a",
						},
						{
							Type: ElementType,
							Name: "b",
						},
					},
					Out: Parameters{
						{
							Kind: reflect.Bool,
						},
					},
				},
			},
		},
		Out: Parameters{
			{
				Type: CollectionType,
			},
		},
	},
}
