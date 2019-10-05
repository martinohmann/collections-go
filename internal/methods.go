package internal

import (
	"fmt"
	"reflect"

	"github.com/martinohmann/collections-go/internal/method"
	"github.com/pkg/errors"
)

type collectionTypes struct {
	obj   reflect.Type
	slice reflect.Type
	elem  reflect.Type
}

func EnsureCollectionMethods(objType, sliceType reflect.Type) error {
	if sliceType == nil {
		sliceType = reflect.TypeOf((*interface{})(nil)).Elem()
	}

	elemType := getElemType(sliceType)
	if elemType == nil {
		return errors.Errorf("invalid slice type %s", sliceType)
	}

	types := collectionTypes{
		obj:   objType,
		slice: sliceType,
		elem:  elemType,
	}

	for methodName, signature := range method.CollectionMethods {
		m, ok := objType.MethodByName(methodName)
		if !ok {
			return errors.Errorf("%s does not have expected method %s", objType, methodName)
		}

		err := ensureSignature(methodName, types, m.Type, signature)
		if err != nil {
			return errors.Wrapf(err, "while validating type %s", objType)
		}
	}

	return nil
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

func ensureSignature(methodName string, types collectionTypes, methodType reflect.Type, sig method.Signature) error {
	if sig.Variadic && !methodType.IsVariadic() {
		return errors.Errorf("expected method %s to be variadic but it is not", methodName)
	} else if !sig.Variadic && methodType.IsVariadic() {
		return errors.Errorf("expected method %s not to be variadic but it is", methodName)
	}

	if methodType.NumIn() != len(sig.In) {
		return errors.Errorf("expected %s to have %d input parameters but it has %d", methodName, len(sig.In), methodType.NumIn())
	}

	for i := 0; i < methodType.NumIn(); i++ {
		err := ensureParameter("input", methodName, i, types, methodType.In(i), sig.In[i])
		if err != nil {
			return err
		}
	}

	if methodType.NumOut() != len(sig.Out) {
		return errors.Errorf("expected %s to have %d output parameters but it has %d", methodName, len(sig.Out), methodType.NumOut())
	}

	for i := 0; i < methodType.NumOut(); i++ {
		err := ensureParameter("output", methodName, i, types, methodType.Out(i), sig.Out[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func ensureParameter(paramType, methodName string, index int, types collectionTypes, p reflect.Type, s method.Parameter) error {
	expectedType := getExpectedType(s, types)

	if expectedType != nil {
		if p != expectedType {
			return errors.Errorf("expected %s parameter #%d of method %s to be of type %s but it is %s", paramType, index, methodName, expectedType, p)
		}

		if p.Kind() != expectedType.Kind() {
			return errors.Errorf("expected %s parameter #%d of method %s to be of kind %s but it is %s", paramType, index, methodName, expectedType.Kind(), p.Kind())
		}
	} else if p.Kind() != s.Kind {
		return errors.Errorf("expected %s parameter #%d of method %s to be of kind %s but it is %s", paramType, index, methodName, s.Kind, p.Kind())
	}

	if p.Kind() == reflect.Func {
		return ensureSignature(fmt.Sprintf("%s.#%d.func", methodName, index), types, p, s.FuncSignature)
	}

	return nil
}
