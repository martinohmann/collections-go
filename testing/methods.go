package testing

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/martinohmann/collections-go/testing/method"
)

type collectionTypes struct {
	obj   reflect.Type
	slice reflect.Type
	elem  reflect.Type
}

func EnsureCollectionMethods(t *testing.T, objType, sliceType reflect.Type) {
	elemType := getElemType(sliceType)
	if elemType == nil {
		t.Fatalf("invalid slice type %s", sliceType)
	}

	types := collectionTypes{
		obj:   objType,
		slice: sliceType,
		elem:  elemType,
	}

	for methodName, signature := range method.CollectionMethods {
		m, ok := types.obj.MethodByName(methodName)
		if !ok {
			t.Fatalf("%s does not have expected method %s", types.obj, methodName)
		}

		ensureSignature(t, methodName, types, m.Type, signature)
	}
}

func getElemType(sliceType reflect.Type) reflect.Type {
	switch sliceType.Kind() {
	case reflect.Interface:
		return sliceType
	case reflect.Slice:
		return sliceType.Elem()
	}

	return nil
}

func getExpectedType(param method.Parameter, types collectionTypes) reflect.Type {
	switch {
	case param.ElementType:
		return types.elem
	case param.SliceType:
		return types.slice
	case param.CollectionType:
		return types.obj
	}

	return nil
}

func ensureSignature(t *testing.T, methodName string, types collectionTypes, methodType reflect.Type, sig method.Signature) {
	if sig.Variadic && !methodType.IsVariadic() {
		t.Fatalf("expected method %s to be variadic but it is not", methodName)
	} else if !sig.Variadic && methodType.IsVariadic() {
		t.Fatalf("expected method %s not to be variadic but it is", methodName)
	}

	if methodType.NumIn() != len(sig.In) {
		t.Fatalf("expected %s to have %d input parameters but it has %d", methodName, len(sig.In), methodType.NumIn())
	}

	for i := 0; i < methodType.NumIn(); i++ {
		ensureParameter(t, "input", methodName, i, types, methodType.In(i), sig.In[i])
	}

	if methodType.NumOut() != len(sig.Out) {
		t.Fatalf("expected %s to have %d output parameters but it has %d", methodName, len(sig.Out), methodType.NumOut())
	}

	for i := 0; i < methodType.NumOut(); i++ {
		ensureParameter(t, "output", methodName, i, types, methodType.Out(i), sig.Out[i])
	}
}

func ensureParameter(t *testing.T, paramType, methodName string, index int, types collectionTypes, p reflect.Type, s method.Parameter) {
	expectedType := getExpectedType(s, types)

	if expectedType != nil {
		if p != expectedType {
			t.Fatalf("expected %s parameter #%d of method %s to be of type %s but it is %s", paramType, index, methodName, expectedType, p)
		}

		if p.Kind() != expectedType.Kind() {
			t.Fatalf("expected %s parameter #%d of method %s to be of kind %s but it is %s", paramType, index, methodName, expectedType.Kind(), p.Kind())
		}
	} else if p.Kind() != s.Kind {
		t.Fatalf("expected %s parameter #%d of method %s to be of kind %s but it is %s", paramType, index, methodName, s.Kind, p.Kind())
	}

	if p.Kind() == reflect.Func {
		ensureSignature(t, fmt.Sprintf("%s.#%d.func", methodName, index), types, p, s.FuncSignature)
	}
}
