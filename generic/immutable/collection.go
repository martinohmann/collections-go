package immutable

import (
	"reflect"
	"sort"

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

func (c *Collection) EachIndex(fn func(interface{}, int)) {
	for i := 0; i < c.rval.Len(); i++ {
		fn(c.valueAt(i), i)
	}
}

func (c *Collection) Each(fn func(interface{})) {
	c.EachIndex(func(val interface{}, _ int) {
		fn(val)
	})
}

func (c *Collection) IndexOf(el interface{}) int {
	for i := 0; i < c.rval.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return i
		}
	}

	return -1
}

func (c *Collection) First() interface{} {
	return c.Nth(0)
}

func (c *Collection) FirstN(n int) *Collection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

func (c *Collection) Last() interface{} {
	return c.Nth(c.Len() - 1)
}

func (c *Collection) LastN(n int) *Collection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

func (c *Collection) Get(idx int) interface{} {
	return c.Nth(idx)
}

func (c *Collection) Nth(idx int) interface{} {
	return c.valueAt(idx)
}

func (c *Collection) Remove(idx int) *Collection {
	s := c.copySlice()

	s = reflect.AppendSlice(s.Slice(0, idx), s.Slice(idx+1, s.Len()))

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) RemoveItem(item interface{}) *Collection {
	s := c.copySlice()

	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), item) {
			s = reflect.AppendSlice(s.Slice(0, i), s.Slice(i+1, s.Len()))
		}
	}

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) InsertItem(item interface{}, idx int) *Collection {
	s := c.copySlice()
	s = reflect.Append(s, reflect.ValueOf(c.zeroValue))
	reflect.Copy(s.Slice(idx+1, s.Len()), s.Slice(idx, s.Len()-1))
	s.Index(idx).Set(reflect.ValueOf(item))

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Cut(i, j int) *Collection {
	s := c.copySlice()
	s = reflect.AppendSlice(s.Slice(0, i), s.Slice(j, s.Len()))

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Slice(i, j int) *Collection {
	s := c.copySlice()
	s = c.rval.Slice(i, j)

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Filter(fn func(interface{}) bool) *Collection {
	s := c.makeSlice()

	for i := 0; i < c.rval.Len(); i++ {
		v := c.rval.Index(i)

		if fn(v.Interface()) {
			s = reflect.Append(s, v)
		}
	}

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Collect(fn func(interface{}) bool) *Collection {
	return c.Filter(fn)
}

func (c *Collection) Reject(fn func(interface{}) bool) *Collection {
	return c.Filter(func(v interface{}) bool {
		return !fn(v)
	})
}

func (c *Collection) Partition(fn func(interface{}) bool) (*Collection, *Collection) {
	lhs, rhs := c.makeSlice(), c.makeSlice()

	for i := 0; i < c.rval.Len(); i++ {
		v := c.rval.Index(i)

		if fn(v.Interface()) {
			lhs = reflect.Append(lhs, v)
		} else {
			rhs = reflect.Append(rhs, v)
		}
	}

	return newCollection(c.sliceType, lhs, lhs.Interface()),
		newCollection(c.sliceType, rhs, rhs.Interface())
}

func (c *Collection) Map(fn func(interface{}) interface{}) *Collection {
	s := c.makeSlice()

	for i := 0; i < c.Len(); i++ {
		s = reflect.Append(s, reflect.ValueOf(fn(c.valueAt(i))))
	}

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Reduce(fn func(reducer interface{}, item interface{}) interface{}) interface{} {
	reducer := c.zeroValue

	for i := 0; i < c.Len(); i++ {
		reducer = fn(reducer, c.valueAt(i))
	}

	return reducer
}

func (c *Collection) Find(fn func(interface{}) bool) interface{} {
	item, _ := c.FindOk(fn)

	return item
}

func (c *Collection) FindOk(fn func(interface{}) bool) (interface{}, bool) {
	for i := 0; i < c.rval.Len(); i++ {
		item := c.valueAt(i)

		if fn(item) {
			return item, true
		}
	}

	return reflect.Zero(c.elemType).Interface(), false
}

func (c *Collection) Any(fn func(interface{}) bool) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if fn(c.valueAt(i)) {
			return true
		}
	}

	return false
}

func (c *Collection) All(fn func(interface{}) bool) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if !fn(c.valueAt(i)) {
			return false
		}
	}

	return true
}

func (c *Collection) Contains(el interface{}) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return true
		}
	}

	return false
}

func (c *Collection) Len() int {
	return c.rval.Len()
}

func (c *Collection) Cap() int {
	return c.rval.Cap()
}

func (c *Collection) Append(items ...interface{}) *Collection {
	s := c.copySlice()

	for _, item := range items {
		s = reflect.Append(s, reflect.ValueOf(item))
	}

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Prepend(items ...interface{}) *Collection {
	s := reflect.MakeSlice(c.sliceType, 0, c.rval.Len()+len(items))

	for _, item := range items {
		s = reflect.Append(s, reflect.ValueOf(item))
	}

	s = reflect.AppendSlice(s, c.rval)

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Copy() *Collection {
	s := c.copySlice()

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Sort(fn func(interface{}, interface{}) bool) *Collection {
	d := c.Copy()

	sort.Slice(d.items, d.lessFunc(fn))

	return d
}

func (c *Collection) IsSorted(fn func(interface{}, interface{}) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *Collection) lessFunc(fn func(interface{}, interface{}) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.valueAt(i), c.valueAt(j))
	}
}

func (c *Collection) Reverse() *Collection {
	newSlice := c.makeSlice()

	for i := c.rval.Len() - 1; i >= 0; i-- {
		newSlice = reflect.Append(newSlice, c.rval.Index(i))
	}

	return newCollection(c.sliceType, newSlice, newSlice.Interface())
}
