// +build !ignore_autogenerated

// Code generated by collections-gen. DO NOT EDIT.

package collections

import (
	"bytes"
	"sort"
)

// ImmutableByteSlice is an immutable collection of []byte values.
type ImmutableByteSlice struct {
	items [][]byte
}

// NewImmutableByteSlice creates a new immutable collection from a slice of []byte.
func NewImmutableByteSlice(items [][]byte) *ImmutableByteSlice {
	return &ImmutableByteSlice{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *ImmutableByteSlice) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of []byte values used by the
// collection.
func (c *ImmutableByteSlice) Items() [][]byte {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *ImmutableByteSlice) EachIndex(fn func([]byte, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *ImmutableByteSlice) Each(fn func([]byte)) {
	c.EachIndex(func(item []byte, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *ImmutableByteSlice) IndexOf(el []byte) int {
	for i, item := range c.items {
		if bytes.Equal(item, el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *ImmutableByteSlice) First() []byte {
	return c.Nth(0)
}

// FirstN returns the first n []byte items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *ImmutableByteSlice) FirstN(n int) [][]byte {
	if n > c.Len() {
		return c.Copy().Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *ImmutableByteSlice) Last() []byte {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n []byte items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *ImmutableByteSlice) LastN(n int) [][]byte {
	if c.Len()-n < 0 {
		return c.Copy().Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *ImmutableByteSlice) Get(pos int) []byte {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *ImmutableByteSlice) Nth(pos int) []byte {
	return c.items[pos]
}

// Len returns the length of the underlying []byte slice.
func (c *ImmutableByteSlice) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying []byte slice.
func (c *ImmutableByteSlice) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableByteSlice) Append(items ...[]byte) *ImmutableByteSlice {
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
}

// Prepend prepends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableByteSlice) Prepend(items ...[]byte) *ImmutableByteSlice {
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
}

// Copy creates a copy of the collection and the underlying []byte slice.
func (c *ImmutableByteSlice) Copy() *ImmutableByteSlice {
	s := make([][]byte, c.Len(), c.Len())
	copy(s, c.items)

	return NewImmutableByteSlice(s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableByteSlice) Filter(fn func([]byte) bool) *ImmutableByteSlice {
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue []byte

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = zeroValue
	}

	d.items = s

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableByteSlice) Collect(fn func([]byte) bool) *ImmutableByteSlice {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The original collection is not altered.
func (c *ImmutableByteSlice) Reject(fn func([]byte) bool) *ImmutableByteSlice {
	return c.Filter(func(v []byte) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *ImmutableByteSlice) Partition(fn func([]byte) bool) (*ImmutableByteSlice, *ImmutableByteSlice) {
	lhs := make([][]byte, 0, c.Len())
	rhs := make([][]byte, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewImmutableByteSlice(lhs), NewImmutableByteSlice(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableByteSlice) Map(fn func([]byte) []byte) *ImmutableByteSlice {
	return c.MapIndex(func(item []byte, _ int) []byte {
		return fn(item)
	})
}

// MapIndex calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableByteSlice) MapIndex(fn func([]byte, int) []byte) *ImmutableByteSlice {
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item, i)

	}

	return d
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero []byte value on the first invocation.
func (c *ImmutableByteSlice) Reduce(fn func(reducer []byte, item []byte) []byte) []byte {
	var reducer []byte

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// []byte value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *ImmutableByteSlice) Find(fn func([]byte) bool) []byte {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// []byte value. The second return value denotes whether the condition
// matched any item or not.
func (c *ImmutableByteSlice) FindOk(fn func([]byte) bool) ([]byte, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue []byte
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *ImmutableByteSlice) Any(fn func([]byte) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *ImmutableByteSlice) All(fn func([]byte) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *ImmutableByteSlice) Contains(el []byte) bool {
	for _, item := range c.items {
		if bytes.Equal(item, el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
// The result will be a copy of c which is sorted, the original collection is
// not altered.
func (c *ImmutableByteSlice) Sort(fn func([]byte, []byte) bool) *ImmutableByteSlice {
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *ImmutableByteSlice) IsSorted(fn func([]byte, []byte) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *ImmutableByteSlice) lessFunc(fn func([]byte, []byte) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *ImmutableByteSlice) Reverse() *ImmutableByteSlice {
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableByteSlice) Remove(pos int) *ImmutableByteSlice {
	d := c.Copy()
	d.items = append(d.items[:pos], d.items[pos+1:]...)
	return d
}

// RemoveItem removes all instances of item from the collection and returns it.
// The result is a new collection, the original is not modified.
func (c *ImmutableByteSlice) RemoveItem(item []byte) *ImmutableByteSlice {
	d := c.Copy()

	for i, el := range d.items {
		if bytes.Equal(el, item) {
			d.items = append(d.items[:i], d.items[i+1:]...)
		}
	}

	return d
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableByteSlice) InsertItem(item []byte, pos int) *ImmutableByteSlice {
	var zeroValue []byte
	d := c.Copy()
	d.items = append(d.items, zeroValue)
	copy(d.items[pos+1:], d.items[pos:])
	d.items[pos] = item
	return d
}

// Cut returns a copy of the underlying []byte slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *ImmutableByteSlice) Cut(i, j int) [][]byte {
	d := c.Copy()
	return append(d.items[:i], d.items[j:]...)
}

// Slice returns the []byte items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *ImmutableByteSlice) Slice(i, j int) [][]byte {
	return c.Copy().items[i:j]
}
