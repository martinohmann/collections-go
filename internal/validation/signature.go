package validation

import (
	"fmt"
	"reflect"
	"strings"
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
	Name          string
}

func (p Parameter) String() string {
	if p.Name == "" {
		return p.TypeString()
	}

	return p.Name + " " + p.TypeString()
}

func (p Parameter) TypeString() string {
	switch p.Type {
	case ReflectionType:
		switch p.Kind {
		case reflect.Func:
			return "func" + p.FuncSignature.String()
		case reflect.Interface:
			return "interface{}"
		case reflect.Slice:
			return "ElemType"
		default:
			return p.Kind.String()
		}
	case CollectionType:
		return "*Collection"
	case SliceType:
		return "[]ElemType"
	case ElementType:
		return "ElemType"
	default:
		return "interface{}"
	}
}

type Parameters []Parameter

func (p Parameters) String() string {
	params := make([]string, len(p))

	for i, param := range p {
		if param.Name != "" && i < len(p)-1 && p[i+1].TypeString() == param.TypeString() {
			params[i] = param.Name
		} else {
			params[i] = param.String()
		}
	}

	return strings.Join(params, ", ")
}

// Signature describes the signature of a function.
type Signature struct {
	NumIn    int
	NumOut   int
	In       Parameters
	Out      Parameters
	Variadic bool
}

func (s Signature) String() string {
	sin := s.In
	if s.Variadic {
		sin = s.In[0 : len(s.In)-1]
	}

	in := sin.String()
	if s.Variadic {
		if len(in) > 0 {
			in += ", "
		}

		vp := s.In[len(s.In)-1]
		if vp.Name != "" {
			in += vp.Name + " "
		}

		in += "..." + vp.TypeString()
	}

	out := s.Out.String()
	if strings.Index(out, " ") != -1 {
		out = "(" + out + ")"
	}

	if len(out) > 0 {
		out = " " + out
	}

	return fmt.Sprintf("(%s)%s", in, out)
}
