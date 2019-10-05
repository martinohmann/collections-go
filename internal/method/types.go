package method

import "reflect"

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
