// Code generated by collections-gen. DO NOT EDIT.

package customtype

import (
	"sort"
)

// Collection is a collection of *Type values.
type Collection struct {
	items []*Type
}

// NewCollection creates a new collection from a slice of *Type.
func NewCollection(items []*Type) *Collection {
	return &Collection{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *Collection) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of *Type values used by the
// collection.
func (c *Collection) Items() []*Type {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *Collection) EachIndex(fn func(*Type, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *Collection) Each(fn func(*Type)) {
	c.EachIndex(func(item *Type, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *Collection) IndexOf(el *Type) int {
	for i, item := range c.items {
		if Equal(item, el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *Collection) First() *Type {
	return c.Nth(0)
}

// FirstN returns the first n *Type items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *Collection) FirstN(n int) []*Type {
	if n > c.Len() {
		return c.Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *Collection) Last() *Type {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n *Type items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *Collection) LastN(n int) []*Type {
	if c.Len()-n < 0 {
		return c.Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *Collection) Get(pos int) *Type {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *Collection) Nth(pos int) *Type {
	return c.items[pos]
}

// Len returns the length of the underlying *Type slice.
func (c *Collection) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying *Type slice.
func (c *Collection) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.
func (c *Collection) Append(items ...*Type) *Collection {
	c.items = append(c.items, items...)
	return c
}

// Prepend prepends items and returns the collection.
func (c *Collection) Prepend(items ...*Type) *Collection {
	c.items = append(items, c.items...)
	return c
}

// Copy creates a copy of the collection and the underlying *Type slice.
func (c *Collection) Copy() *Collection {
	s := make([]*Type, c.Len(), c.Len())
	copy(s, c.items)

	return NewCollection(s)
}

// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *Collection) Filter(fn func(*Type) bool) *Collection {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue *Type

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = zeroValue
	}

	c.items = s

	return c
}

// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *Collection) Collect(fn func(*Type) bool) *Collection {
	return c.Filter(fn)
}

// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
func (c *Collection) Reject(fn func(*Type) bool) *Collection {
	return c.Filter(func(v *Type) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *Collection) Partition(fn func(*Type) bool) (*Collection, *Collection) {
	lhs := make([]*Type, 0, c.Len())
	rhs := make([]*Type, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewCollection(lhs), NewCollection(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *Collection) Map(fn func(*Type) *Type) *Collection {
	for i, item := range c.items {
		c.items[i] = fn(item)

	}

	return c
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero *Type value on the first invocation.
func (c *Collection) Reduce(fn func(reducer *Type, item *Type) *Type) *Type {
	var reducer *Type

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// *Type value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *Collection) Find(fn func(*Type) bool) *Type {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// *Type value. The second return value denotes whether the condition
// matched any item or not.
func (c *Collection) FindOk(fn func(*Type) bool) (*Type, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue *Type
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *Collection) Any(fn func(*Type) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *Collection) All(fn func(*Type) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *Collection) Contains(el *Type) bool {
	for _, item := range c.items {
		if Equal(item, el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
func (c *Collection) Sort(fn func(*Type, *Type) bool) *Collection {
	sort.Slice(c.items, c.lessFunc(fn))
	return c
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *Collection) IsSorted(fn func(*Type, *Type) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *Collection) lessFunc(fn func(*Type, *Type) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse reverses the order of the collection items in place and returns c.
func (c *Collection) Reverse() *Collection {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
func (c *Collection) Remove(pos int) *Collection {
	c.items = append(c.items[:pos], c.items[pos+1:]...)
	return c
}

// RemoveItem removes all instances of item from the collection and returns it.
func (c *Collection) RemoveItem(item *Type) *Collection {
	for i, el := range c.items {
		if Equal(el, item) {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
func (c *Collection) InsertItem(item *Type, pos int) *Collection {
	var zeroValue *Type
	c.items = append(c.items, zeroValue)
	copy(c.items[pos+1:], c.items[pos:])
	c.items[pos] = item
	return c
}

// Cut returns a copy of the underlying *Type slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *Collection) Cut(i, j int) []*Type {
	s := make([]*Type, 0, c.Cap())
	s = append(s, c.items[:i]...)
	return append(s, c.items[j:]...)
}

// Slice returns the *Type items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *Collection) Slice(i, j int) []*Type {
	return c.items[i:j]
}
