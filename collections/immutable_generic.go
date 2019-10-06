package collections

import (
	"reflect"
	"sort"

	"github.com/pkg/errors"
)

type ImmutableGeneric struct {
	sliceType reflect.Type
	elemType  reflect.Type
	zeroValue interface{}
	rval      reflect.Value
	items     interface{}
}

func NewImmutable(items interface{}) *ImmutableGeneric {
	c, err := SafeNewImmutable(items)
	if err != nil {
		panic(err)
	}

	return c
}

func SafeNewImmutable(items interface{}) (*ImmutableGeneric, error) {
	if items == nil {
		return nil, errors.Errorf("cannot create *ImmutableGeneric for nil interface{}")
	}

	t := reflect.TypeOf(items)

	switch t.Kind() {
	case reflect.Slice:
		return newImmutableGeneric(t, reflect.ValueOf(items), items), nil
	default:
		return nil, errors.Errorf("expected slice type, got %T", items)
	}
}

func newImmutableGeneric(t reflect.Type, rval reflect.Value, items interface{}) *ImmutableGeneric {
	elemType := t.Elem()
	return &ImmutableGeneric{
		items:     items,
		rval:      rval,
		sliceType: t,
		elemType:  elemType,
		zeroValue: reflect.Zero(elemType).Interface(),
	}
}

func (c *ImmutableGeneric) Items() interface{} {
	return c.items
}

func (c *ImmutableGeneric) makeSlice() reflect.Value {
	return reflect.MakeSlice(c.sliceType, 0, c.rval.Len())
}

func (c *ImmutableGeneric) copySlice() reflect.Value {
	s := reflect.MakeSlice(c.sliceType, c.rval.Len(), c.rval.Len())

	reflect.Copy(s, c.rval)

	return s
}

func (c *ImmutableGeneric) valueAt(idx int) interface{} {
	return c.rval.Index(idx).Interface()
}

func (c *ImmutableGeneric) EachIndex(fn func(interface{}, int)) {
	for i := 0; i < c.rval.Len(); i++ {
		fn(c.valueAt(i), i)
	}
}

func (c *ImmutableGeneric) Each(fn func(interface{})) {
	c.EachIndex(func(val interface{}, _ int) {
		fn(val)
	})
}

func (c *ImmutableGeneric) IndexOf(el interface{}) int {
	for i := 0; i < c.rval.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return i
		}
	}

	return -1
}

func (c *ImmutableGeneric) First() interface{} {
	return c.Nth(0)
}

func (c *ImmutableGeneric) FirstN(n int) *ImmutableGeneric {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

func (c *ImmutableGeneric) Last() interface{} {
	return c.Nth(c.Len() - 1)
}

func (c *ImmutableGeneric) LastN(n int) *ImmutableGeneric {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

func (c *ImmutableGeneric) Get(idx int) interface{} {
	return c.Nth(idx)
}

func (c *ImmutableGeneric) Nth(idx int) interface{} {
	return c.valueAt(idx)
}

func (c *ImmutableGeneric) Remove(idx int) *ImmutableGeneric {
	s := c.copySlice()

	s = reflect.AppendSlice(s.Slice(0, idx), s.Slice(idx+1, s.Len()))

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) RemoveItem(item interface{}) *ImmutableGeneric {
	s := c.copySlice()

	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), item) {
			s = reflect.AppendSlice(s.Slice(0, i), s.Slice(i+1, s.Len()))
		}
	}

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) InsertItem(item interface{}, idx int) *ImmutableGeneric {
	s := c.copySlice()
	s = reflect.Append(s, reflect.ValueOf(c.zeroValue))
	reflect.Copy(s.Slice(idx+1, s.Len()), s.Slice(idx, s.Len()-1))
	s.Index(idx).Set(reflect.ValueOf(item))

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) Cut(i, j int) *ImmutableGeneric {
	s := c.copySlice()
	s = reflect.AppendSlice(s.Slice(0, i), s.Slice(j, s.Len()))

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) Slice(i, j int) *ImmutableGeneric {
	s := c.copySlice()
	s = c.rval.Slice(i, j)

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) Filter(fn func(interface{}) bool) *ImmutableGeneric {
	s := c.makeSlice()

	for i := 0; i < c.rval.Len(); i++ {
		v := c.rval.Index(i)

		if fn(v.Interface()) {
			s = reflect.Append(s, v)
		}
	}

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) Collect(fn func(interface{}) bool) *ImmutableGeneric {
	return c.Filter(fn)
}

func (c *ImmutableGeneric) Reject(fn func(interface{}) bool) *ImmutableGeneric {
	return c.Filter(func(v interface{}) bool {
		return !fn(v)
	})
}

func (c *ImmutableGeneric) Partition(fn func(interface{}) bool) (*ImmutableGeneric, *ImmutableGeneric) {
	lhs, rhs := c.makeSlice(), c.makeSlice()

	for i := 0; i < c.rval.Len(); i++ {
		v := c.rval.Index(i)

		if fn(v.Interface()) {
			lhs = reflect.Append(lhs, v)
		} else {
			rhs = reflect.Append(rhs, v)
		}
	}

	return newImmutableGeneric(c.sliceType, lhs, lhs.Interface()),
		newImmutableGeneric(c.sliceType, rhs, rhs.Interface())
}

func (c *ImmutableGeneric) Map(fn func(interface{}) interface{}) *ImmutableGeneric {
	s := c.makeSlice()

	for i := 0; i < c.Len(); i++ {
		s = reflect.Append(s, reflect.ValueOf(fn(c.valueAt(i))))
	}

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) Reduce(fn func(reducer interface{}, item interface{}) interface{}) interface{} {
	reducer := c.zeroValue

	for i := 0; i < c.Len(); i++ {
		reducer = fn(reducer, c.valueAt(i))
	}

	return reducer
}

func (c *ImmutableGeneric) Find(fn func(interface{}) bool) interface{} {
	item, _ := c.FindOk(fn)

	return item
}

func (c *ImmutableGeneric) FindOk(fn func(interface{}) bool) (interface{}, bool) {
	for i := 0; i < c.rval.Len(); i++ {
		item := c.valueAt(i)

		if fn(item) {
			return item, true
		}
	}

	return reflect.Zero(c.elemType).Interface(), false
}

func (c *ImmutableGeneric) Any(fn func(interface{}) bool) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if fn(c.valueAt(i)) {
			return true
		}
	}

	return false
}

func (c *ImmutableGeneric) All(fn func(interface{}) bool) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if !fn(c.valueAt(i)) {
			return false
		}
	}

	return true
}

func (c *ImmutableGeneric) Contains(el interface{}) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return true
		}
	}

	return false
}

func (c *ImmutableGeneric) Len() int {
	return c.rval.Len()
}

func (c *ImmutableGeneric) Cap() int {
	return c.rval.Cap()
}

func (c *ImmutableGeneric) Append(items ...interface{}) *ImmutableGeneric {
	s := c.copySlice()

	for _, item := range items {
		s = reflect.Append(s, reflect.ValueOf(item))
	}

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) Prepend(items ...interface{}) *ImmutableGeneric {
	s := reflect.MakeSlice(c.sliceType, 0, c.rval.Len()+len(items))

	for _, item := range items {
		s = reflect.Append(s, reflect.ValueOf(item))
	}

	s = reflect.AppendSlice(s, c.rval)

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) Copy() *ImmutableGeneric {
	s := c.copySlice()

	return newImmutableGeneric(c.sliceType, s, s.Interface())
}

func (c *ImmutableGeneric) Sort(fn func(interface{}, interface{}) bool) *ImmutableGeneric {
	d := c.Copy()

	sort.Slice(d.items, d.lessFunc(fn))

	return d
}

func (c *ImmutableGeneric) IsSorted(fn func(interface{}, interface{}) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *ImmutableGeneric) lessFunc(fn func(interface{}, interface{}) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.valueAt(i), c.valueAt(j))
	}
}

func (c *ImmutableGeneric) Reverse() *ImmutableGeneric {
	newSlice := c.makeSlice()

	for i := c.rval.Len() - 1; i >= 0; i-- {
		newSlice = reflect.Append(newSlice, c.rval.Index(i))
	}

	return newImmutableGeneric(c.sliceType, newSlice, newSlice.Interface())
}
