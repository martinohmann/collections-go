package collections

import (
	"reflect"
	"sort"

	"github.com/pkg/errors"
)

// Generic is a generalized implementation of an immutable collection
// that works with slices of arbitrary type. If performance matters consider
// using a collection specialized for your type as Generic makes have
// use of reflection to access the underlying slice which is very slow.
type Generic struct {
	sliceType reflect.Type
	zeroValue interface{}
	rval      reflect.Value
}

// New creates a new generalized immutable collection from items. Will
// panic if items is not a slice type. Consider using SafeNew instead
// if you cannot guarantee items to always be a slice.
func New(items interface{}) *Generic {
	c, err := SafeNew(items)
	if err != nil {
		panic(err)
	}

	return c
}

// SafeNew creates a new generalized immutable collection from items.
// Will return an error if items is not a slice type. Consider using
// SafeNew instead.
func SafeNew(items interface{}) (*Generic, error) {
	if items == nil {
		return nil, errors.Errorf("cannot create *Generic for nil interface{}")
	}

	t := reflect.TypeOf(items)

	switch t.Kind() {
	case reflect.Slice:
		return newGeneric(t, reflect.ValueOf(items)), nil
	default:
		return nil, errors.Errorf("expected slice type, got %T", items)
	}
}

func newGeneric(t reflect.Type, rval reflect.Value) *Generic {
	return &Generic{
		rval:      rval,
		sliceType: t,
		zeroValue: reflect.Zero(t.Elem()).Interface(),
	}
}

func (c *Generic) makeSlice() reflect.Value {
	return reflect.MakeSlice(c.sliceType, 0, c.rval.Len())
}

func (c *Generic) copySlice() reflect.Value {
	s := reflect.MakeSlice(c.sliceType, c.rval.Len(), c.rval.Len())
	reflect.Copy(s, c.rval)
	return s
}

func (c *Generic) valueAt(idx int) interface{} {
	return c.rval.Index(idx).Interface()
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *Generic) Interface() interface{} {
	return c.rval.Interface()
}

// Items returns the underlying slice used by the collection.
func (c *Generic) Items() interface{} {
	return c.rval.Interface()
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *Generic) EachIndex(fn func(interface{}, int)) {
	for i := 0; i < c.rval.Len(); i++ {
		fn(c.valueAt(i), i)
	}
}

// Each calls fn for every item in the collection.
func (c *Generic) Each(fn func(interface{})) {
	c.EachIndex(func(val interface{}, _ int) {
		fn(val)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *Generic) IndexOf(el interface{}) int {
	for i := 0; i < c.rval.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *Generic) First() interface{} {
	return c.Nth(0)
}

// FirstN returns the first n items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *Generic) FirstN(n int) interface{} {
	if n > c.Len() {
		return c.Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *Generic) Last() interface{} {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n string items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *Generic) LastN(n int) interface{} {
	if c.Len()-n < 0 {
		return c.Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at idx from the collection. Will panic if the
// underlying slice is shorter than idx+1.
func (c *Generic) Get(idx int) interface{} {
	return c.Nth(idx)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than idx+1.
func (c *Generic) Nth(idx int) interface{} {
	return c.valueAt(idx)
}

// Len returns the length of the underlying slice.
func (c *Generic) Len() int {
	return c.rval.Len()
}

// Cap returns the capacity of the underlying slice.
func (c *Generic) Cap() int {
	return c.rval.Cap()
}

// Append appends items and returns the collection. Will panic if items are not
// of the slices element type.
func (c *Generic) Append(items ...interface{}) *Generic {
	for _, item := range items {
		c.rval = reflect.Append(c.rval, reflect.ValueOf(item))
	}

	return c
}

// Prepend prepends items and returns the collection. Will panic if items are
// not of the slices element type.
func (c *Generic) Prepend(items ...interface{}) *Generic {
	s := c.copySlice()

	for i := 0; i < len(items); i++ {
		c.rval = reflect.Append(c.rval, reflect.ValueOf(c.zeroValue))
	}

	n := 0
	for _, item := range items {
		c.rval.Index(n).Set(reflect.ValueOf(item))
		n++
	}

	for i := 0; i < s.Len(); i++ {
		c.rval.Index(n + i).Set(s.Index(i))
	}

	return c
}

// Copy creates a copy of the collection and the underlying slice.
func (c *Generic) Copy() *Generic {
	s := c.copySlice()
	return newGeneric(c.sliceType, s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection.
func (c *Generic) Filter(fn func(interface{}) bool) *Generic {
	n := 0
	for i := 0; i < c.rval.Len(); i++ {
		v := c.rval.Index(i)

		if fn(v.Interface()) {
			c.rval.Index(n).Set(v)
			n++
		}
	}

	c.rval = c.rval.Slice(0, n)

	return c
}

// Collect collects all items for which fn evaluates to true into a new
// collection.
func (c *Generic) Collect(fn func(interface{}) bool) *Generic {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection.
func (c *Generic) Reject(fn func(interface{}) bool) *Generic {
	return c.Filter(func(v interface{}) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *Generic) Partition(fn func(interface{}) bool) (*Generic, *Generic) {
	lhs, rhs := c.makeSlice(), c.makeSlice()

	for i := 0; i < c.rval.Len(); i++ {
		v := c.rval.Index(i)

		if fn(v.Interface()) {
			lhs = reflect.Append(lhs, v)
		} else {
			rhs = reflect.Append(rhs, v)
		}
	}

	return newGeneric(c.sliceType, lhs), newGeneric(c.sliceType, rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. Will panic if the value
// returned by fn is not of the slices element type.
func (c *Generic) Map(fn func(interface{}) interface{}) *Generic {
	for i := 0; i < c.Len(); i++ {
		c.rval.Index(i).Set(reflect.ValueOf(fn(c.valueAt(i))))
	}

	return c
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero value of the slice's element type on the first
// invocation. Will panic if the value returned by fn is not of the slices
// element type.
func (c *Generic) Reduce(fn func(reducer interface{}, item interface{}) interface{}) interface{} {
	reducer := c.zeroValue

	for i := 0; i < c.Len(); i++ {
		reducer = fn(reducer, c.valueAt(i))
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero value
// of the slice's element type. If you need to distinguish zero values from a
// condition that did not match any item consider using FindOk instead.
func (c *Generic) Find(fn func(interface{}) bool) interface{} {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// value of the slice's element type. The second return value denotes whether
// the condition matched any item or not.
func (c *Generic) FindOk(fn func(interface{}) bool) (interface{}, bool) {
	for i := 0; i < c.rval.Len(); i++ {
		item := c.valueAt(i)

		if fn(item) {
			return item, true
		}
	}

	return c.zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *Generic) Any(fn func(interface{}) bool) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if fn(c.valueAt(i)) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *Generic) All(fn func(interface{}) bool) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if !fn(c.valueAt(i)) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *Generic) Contains(el interface{}) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
func (c *Generic) Sort(fn func(interface{}, interface{}) bool) *Generic {
	sort.Slice(c.rval.Interface(), c.lessFunc(fn))
	return c
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *Generic) IsSorted(fn func(interface{}, interface{}) bool) bool {
	return sort.SliceIsSorted(c.rval.Interface(), c.lessFunc(fn))
}

func (c *Generic) lessFunc(fn func(interface{}, interface{}) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.valueAt(i), c.valueAt(j))
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *Generic) Reverse() *Generic {
	for l, r := 0, c.Len()-1; l < r; l, r = l+1, r-1 {
		v := c.rval.Index(l).Interface()
		c.rval.Index(l).Set(c.rval.Index(r))
		c.rval.Index(r).Set(reflect.ValueOf(v))
	}

	return c
}

// Remove removes the collection item at position idx. Will panic if idx is out
// of bounds.
func (c *Generic) Remove(idx int) *Generic {
	c.rval = reflect.AppendSlice(c.rval.Slice(0, idx), c.rval.Slice(idx+1, c.rval.Len()))
	return c
}

// RemoveItem removes all instances of item from the collection and returns it.
func (c *Generic) RemoveItem(item interface{}) *Generic {
	for i := 0; i < c.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), item) {
			c.rval = reflect.AppendSlice(c.rval.Slice(0, i), c.rval.Slice(i+1, c.rval.Len()))
		}
	}

	return c
}

// InsertItem inserts item into the collection at position idx. Will panic if
// idx is out of bounds or if item is not of the slices element type.
func (c *Generic) InsertItem(item interface{}, idx int) *Generic {
	c.rval = reflect.Append(c.rval, reflect.ValueOf(c.zeroValue))
	reflect.Copy(c.rval.Slice(idx+1, c.rval.Len()), c.rval.Slice(idx, c.rval.Len()-1))
	c.rval.Index(idx).Set(reflect.ValueOf(item))
	return c
}

// Cut returns a copy of the underlying string slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *Generic) Cut(i, j int) interface{} {
	s := c.makeSlice()
	s = reflect.AppendSlice(s, c.rval.Slice(0, i))
	s = reflect.AppendSlice(s, c.rval.Slice(j, c.rval.Len()))
	return s.Interface()
}

// Slice returns the items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *Generic) Slice(i, j int) interface{} {
	return c.rval.Slice(i, j).Interface()
}
