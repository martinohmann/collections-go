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
	"All": Signature{
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
	"Any": Signature{
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
	"Append": Signature{
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
	"Cap": Signature{
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
	"Collect": Signature{
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
	"Contains": Signature{
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
	"Copy": Signature{
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
	"Cut": Signature{
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
	"Each": Signature{
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
	"EachIndex": Signature{
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
	"Filter": Signature{
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
	"Find": Signature{
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
	"First": Signature{
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
	"FirstN": Signature{
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
	"Get": Signature{
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
	"IndexOf": Signature{
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
	"InsertItem": Signature{
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
	"IsSorted": Signature{
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
	"Items": Signature{
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
	"Last": Signature{
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
	"LastN": Signature{
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
	"Len": Signature{
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
	"Map": Signature{
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
	"MapIndex": Signature{
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
	"Nth": Signature{
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
	"Partition": Signature{
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
	"Prepend": Signature{
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
	"Reduce": Signature{
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
	"Reject": Signature{
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
	"Remove": Signature{
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
	"RemoveItem": Signature{
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
	"Reverse": Signature{
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
	"Slice": Signature{
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
	"Sort": Signature{
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
