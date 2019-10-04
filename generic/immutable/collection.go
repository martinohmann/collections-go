package immutable

import (
	"reflect"

	"github.com/pkg/errors"
)

type Collection struct {
	sliceType reflect.Type
	elemType  reflect.Type
	zeroValue interface{}
	rval      reflect.Value
	items     interface{}
}

func NewCollection(items interface{}) *Collection {
	c, err := SafeNewCollection(items)
	if err != nil {
		panic(err)
	}

	return c
}

func SafeNewCollection(items interface{}) (*Collection, error) {
	if items == nil {
		return nil, errors.Errorf("cannot create *Collection for nil interface{}")
	}

	t := reflect.TypeOf(items)

	switch t.Kind() {
	case reflect.Slice:
		return newCollection(t, reflect.ValueOf(items), items), nil
	default:
		return nil, errors.Errorf("expected slice type, got %T", items)
	}
}

func newCollection(t reflect.Type, rval reflect.Value, items interface{}) *Collection {
	elemType := t.Elem()
	return &Collection{
		items:     items,
		rval:      rval,
		sliceType: t,
		elemType:  elemType,
		zeroValue: reflect.Zero(elemType).Interface(),
	}
}

func (c *Collection) Items() interface{} {
	return c.items
}

func (c *Collection) makeSlice() reflect.Value {
	return reflect.MakeSlice(c.sliceType, 0, c.rval.Len())
}

func (c *Collection) copySlice() reflect.Value {
	s := reflect.MakeSlice(c.sliceType, c.rval.Len(), c.rval.Len())

	reflect.Copy(s, c.rval)

	return s
}

func (c *Collection) valueAt(idx int) interface{} {
	return c.rval.Index(idx).Interface()
}
