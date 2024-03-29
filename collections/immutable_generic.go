package collections

import (
	"reflect"
	"sort"

	"github.com/pkg/errors"
)

// ImmutableGeneric is a generalized implementation of an immutable collection
// that works with slices of arbitrary type. If performance matters consider
// using a collection specialized for your type as ImmutableGeneric makes have
// use of reflection to access the underlying slice which is very slow.
type ImmutableGeneric struct {
	sliceType reflect.Type
	zeroValue reflect.Value
	items     reflect.Value
}

// NewImmutable creates a new generalized immutable collection from items. Will
// panic if items is not a slice type. Consider using SafeNewImmutable instead
// if you cannot guarantee items to always be a slice.
func NewImmutable(items interface{}) *ImmutableGeneric {
	c, err := SafeNewImmutable(items)
	if err != nil {
		panic(err)
	}

	return c
}

// SafeNewImmutable creates a new generalized immutable collection from items.
// Will return an error if items is not a slice type. Consider using
// SafeNewImmutable instead.
func SafeNewImmutable(items interface{}) (*ImmutableGeneric, error) {
	if items == nil {
		return nil, errors.Errorf("cannot create *ImmutableGeneric for nil interface{}")
	}

	t := reflect.TypeOf(items)

	switch t.Kind() {
	case reflect.Slice:
		return newImmutableGeneric(t, reflect.ValueOf(items)), nil
	default:
		return nil, errors.Errorf("expected slice type, got %T", items)
	}
}

func newImmutableGeneric(t reflect.Type, items reflect.Value) *ImmutableGeneric {
	return &ImmutableGeneric{
		items:     items,
		sliceType: t,
		zeroValue: reflect.Zero(t.Elem()),
	}
}

func (c *ImmutableGeneric) makeSlice() reflect.Value {
	return reflect.MakeSlice(c.sliceType, 0, c.items.Len())
}

func (c *ImmutableGeneric) copySlice() reflect.Value {
	s := reflect.MakeSlice(c.sliceType, c.items.Len(), c.items.Len())
	reflect.Copy(s, c.items)
	return s
}

func (c *ImmutableGeneric) valueAt(pos int) interface{} {
	return c.items.Index(pos).Interface()
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *ImmutableGeneric) Interface() interface{} {
	return c.items.Interface()
}

// Items returns the underlying slice used by the collection.
func (c *ImmutableGeneric) Items() interface{} {
	return c.items.Interface()
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *ImmutableGeneric) EachIndex(fn func(interface{}, int)) {
	for i := 0; i < c.items.Len(); i++ {
		fn(c.valueAt(i), i)
	}
}

// Each calls fn for every item in the collection.
func (c *ImmutableGeneric) Each(fn func(interface{})) {
	c.EachIndex(func(item interface{}, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *ImmutableGeneric) IndexOf(el interface{}) int {
	for i := 0; i < c.items.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *ImmutableGeneric) First() interface{} {
	return c.Nth(0)
}

// FirstN returns the first n items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *ImmutableGeneric) FirstN(n int) interface{} {
	if n > c.Len() {
		return c.Copy().Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *ImmutableGeneric) Last() interface{} {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n string items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *ImmutableGeneric) LastN(n int) interface{} {
	if c.Len()-n < 0 {
		return c.Copy().Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *ImmutableGeneric) Get(pos int) interface{} {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *ImmutableGeneric) Nth(pos int) interface{} {
	return c.valueAt(pos)
}

// Len returns the length of the underlying slice.
func (c *ImmutableGeneric) Len() int {
	return c.items.Len()
}

// Cap returns the capacity of the underlying slice.
func (c *ImmutableGeneric) Cap() int {
	return c.items.Cap()
}

// Append appends items and returns the collection. The
// original collection will not be modified. Will panic if items are not of the
// slices element type.
func (c *ImmutableGeneric) Append(items ...interface{}) *ImmutableGeneric {
	s := c.copySlice()

	for _, item := range items {
		s = reflect.Append(s, reflect.ValueOf(item))
	}

	return newImmutableGeneric(c.sliceType, s)
}

// Prepend prepends items and returns the collection. The original collection
// will not be modified. Will panic if items are not of the slices element
// type.
func (c *ImmutableGeneric) Prepend(items ...interface{}) *ImmutableGeneric {
	l := c.items.Len() + len(items)
	s := reflect.MakeSlice(c.sliceType, l, l)
	reflect.Copy(s.Slice(len(items), s.Len()), c.items)

	for i, item := range items {
		s.Index(i).Set(reflect.ValueOf(item))
	}

	return newImmutableGeneric(c.sliceType, s)
}

// Copy creates a copy of the collection and the underlying slice.
func (c *ImmutableGeneric) Copy() *ImmutableGeneric {
	s := c.copySlice()
	return newImmutableGeneric(c.sliceType, s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableGeneric) Filter(fn func(interface{}) bool) *ImmutableGeneric {
	d := c.Copy()

	n := 0
	for i := 0; i < d.items.Len(); i++ {
		v := d.items.Index(i)

		if fn(v.Interface()) {
			d.items.Index(n).Set(v)
			n++
		}
	}

	d.items = d.items.Slice(0, n)

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableGeneric) Collect(fn func(interface{}) bool) *ImmutableGeneric {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The original collection is not altered.
func (c *ImmutableGeneric) Reject(fn func(interface{}) bool) *ImmutableGeneric {
	return c.Filter(func(v interface{}) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *ImmutableGeneric) Partition(fn func(interface{}) bool) (*ImmutableGeneric, *ImmutableGeneric) {
	lhs, rhs := c.makeSlice(), c.makeSlice()

	for i := 0; i < c.items.Len(); i++ {
		v := c.items.Index(i)

		if fn(v.Interface()) {
			lhs = reflect.Append(lhs, v)
		} else {
			rhs = reflect.Append(rhs, v)
		}
	}

	return newImmutableGeneric(c.sliceType, lhs), newImmutableGeneric(c.sliceType, rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original collection is not
// modified. Will panic if the value returned by fn is not of the slices
// element type.
func (c *ImmutableGeneric) Map(fn func(interface{}) interface{}) *ImmutableGeneric {
	return c.MapIndex(func(item interface{}, _ int) interface{} {
		return fn(item)
	})
}

// MapIndex calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original collection is not
// modified. Will panic if the value returned by fn is not of the slices
// element type.
func (c *ImmutableGeneric) MapIndex(fn func(interface{}, int) interface{}) *ImmutableGeneric {
	s := c.copySlice()

	for i := 0; i < c.Len(); i++ {
		s.Index(i).Set(reflect.ValueOf(fn(s.Index(i).Interface(), i)))
	}

	return newImmutableGeneric(c.sliceType, s)
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero value of the slice's element type on the first
// invocation. Will panic if the value returned by fn is not of the slices
// element type.
func (c *ImmutableGeneric) Reduce(fn func(reducer interface{}, item interface{}) interface{}) interface{} {
	reducer := c.zeroValue.Interface()

	for i := 0; i < c.Len(); i++ {
		reducer = fn(reducer, c.valueAt(i))
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero value
// of the slice's element type. If you need to distinguish zero values from a
// condition that did not match any item consider using FindOk instead.
func (c *ImmutableGeneric) Find(fn func(interface{}) bool) interface{} {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// value of the slice's element type. The second return value denotes whether
// the condition matched any item or not.
func (c *ImmutableGeneric) FindOk(fn func(interface{}) bool) (interface{}, bool) {
	for i := 0; i < c.items.Len(); i++ {
		item := c.valueAt(i)

		if fn(item) {
			return item, true
		}
	}

	return c.zeroValue.Interface(), false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *ImmutableGeneric) Any(fn func(interface{}) bool) bool {
	for i := 0; i < c.items.Len(); i++ {
		if fn(c.valueAt(i)) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *ImmutableGeneric) All(fn func(interface{}) bool) bool {
	for i := 0; i < c.items.Len(); i++ {
		if !fn(c.valueAt(i)) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *ImmutableGeneric) Contains(el interface{}) bool {
	for i := 0; i < c.items.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
// The result will be a copy of c which is sorted, the original collection is
// not altered.
func (c *ImmutableGeneric) Sort(fn func(interface{}, interface{}) bool) *ImmutableGeneric {
	d := c.Copy()
	sort.Slice(d.items.Interface(), d.lessFunc(fn))
	return d
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *ImmutableGeneric) IsSorted(fn func(interface{}, interface{}) bool) bool {
	return sort.SliceIsSorted(c.items.Interface(), c.lessFunc(fn))
}

func (c *ImmutableGeneric) lessFunc(fn func(interface{}, interface{}) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.valueAt(i), c.valueAt(j))
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *ImmutableGeneric) Reverse() *ImmutableGeneric {
	s := reflect.MakeSlice(c.sliceType, c.Len(), c.Len())
	for l, r := 0, c.Len()-1; l < c.Len(); l, r = l+1, r-1 {
		s.Index(l).Set(c.items.Index(r))
	}

	return newImmutableGeneric(c.sliceType, s)
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableGeneric) Remove(pos int) *ImmutableGeneric {
	s := c.copySlice()
	s = reflect.AppendSlice(s.Slice(0, pos), s.Slice(pos+1, s.Len()))

	return newImmutableGeneric(c.sliceType, s)
}

// RemoveItem removes all instances of item from the collection and returns it.
// The result is a new collection, the original is not modified.
func (c *ImmutableGeneric) RemoveItem(item interface{}) *ImmutableGeneric {
	s := c.copySlice()

	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), item) {
			s = reflect.AppendSlice(s.Slice(0, i), s.Slice(i+1, s.Len()))
		}
	}

	return newImmutableGeneric(c.sliceType, s)
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds or if item is not of the slices element type.
// The result is a new collection, the original is not modified.
func (c *ImmutableGeneric) InsertItem(item interface{}, pos int) *ImmutableGeneric {
	s := c.copySlice()
	s = reflect.Append(s, c.zeroValue)
	reflect.Copy(s.Slice(pos+1, s.Len()), s.Slice(pos, s.Len()-1))
	s.Index(pos).Set(reflect.ValueOf(item))
	return newImmutableGeneric(c.sliceType, s)
}

// Cut returns a copy of the underlying string slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *ImmutableGeneric) Cut(i, j int) interface{} {
	s := c.copySlice()
	s = reflect.AppendSlice(s.Slice(0, i), s.Slice(j, s.Len()))
	return s.Interface()
}

// Slice returns the items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *ImmutableGeneric) Slice(i, j int) interface{} {
	return c.copySlice().Slice(i, j).Interface()
}
