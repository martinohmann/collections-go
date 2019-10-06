// Code generated by collections-gen. DO NOT EDIT.

package collections

import (
	"bytes"
	"sort"
)

// ByteSliceCollection is a collection of []byte values.
type ByteSliceCollection struct {
	items [][]byte
}

// NewByteSliceCollection creates a new collection from a slice of []byte.
func NewByteSliceCollection(items [][]byte) *ByteSliceCollection {
	return &ByteSliceCollection{items}
}

// Items returns the underlying slice of []byte values used by the
// collection.
func (c *ByteSliceCollection) Items() [][]byte {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *ByteSliceCollection) EachIndex(fn func([]byte, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *ByteSliceCollection) Each(fn func([]byte)) {
	c.EachIndex(func(item []byte, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *ByteSliceCollection) IndexOf(el []byte) int {
	for i, item := range c.items {
		if bytes.Equal(item, el) {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *ByteSliceCollection) First() []byte {
	return c.Nth(0)
}

// FirstN returns a new collection containing the first n items. Will return
// less than n items if the underlying slice's length is < n.
func (c *ByteSliceCollection) FirstN(n int) *ByteSliceCollection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *ByteSliceCollection) Last() []byte {
	return c.Nth(c.Len() - 1)
}

// LastN returns a new collection containing the last n items. Will return less
// than n items if the underlying slice's length is < n.
func (c *ByteSliceCollection) LastN(n int) *ByteSliceCollection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at idx from the collection. Will panic if the
// underlying slice is shorter than idx+1.
func (c *ByteSliceCollection) Get(idx int) []byte {
	return c.Nth(idx)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than idx+1.
func (c *ByteSliceCollection) Nth(idx int) []byte {
	return c.items[idx]
}

// Len returns the length of the underlying []byte slice.
func (c *ByteSliceCollection) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying []byte slice.
func (c *ByteSliceCollection) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.
func (c *ByteSliceCollection) Append(items ...[]byte) *ByteSliceCollection {
	c.items = append(c.items, items...)
	return c
}

// Prepend prepends items and returns the collection.
func (c *ByteSliceCollection) Prepend(items ...[]byte) *ByteSliceCollection {
	c.items = append(items, c.items...)
	return c
}

// Copy creates a copy of the collection and the underlying []byte slice.
func (c *ByteSliceCollection) Copy() *ByteSliceCollection {
	s := make([][]byte, c.Len(), c.Len())
	copy(s, c.items)

	return NewByteSliceCollection(s)
}

// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *ByteSliceCollection) Filter(fn func([]byte) bool) *ByteSliceCollection {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue []byte

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = zeroValue
	}

	c.items = s

	return c
}

// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *ByteSliceCollection) Collect(fn func([]byte) bool) *ByteSliceCollection {
	return c.Filter(fn)
}

// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
func (c *ByteSliceCollection) Reject(fn func([]byte) bool) *ByteSliceCollection {
	return c.Filter(func(v []byte) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *ByteSliceCollection) Partition(fn func([]byte) bool) (*ByteSliceCollection, *ByteSliceCollection) {
	lhs := make([][]byte, 0, c.Len())
	rhs := make([][]byte, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewByteSliceCollection(lhs), NewByteSliceCollection(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *ByteSliceCollection) Map(fn func([]byte) []byte) *ByteSliceCollection {
	for i, item := range c.items {
		c.items[i] = fn(item)

	}

	return c
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero []byte value on the first invocation.
func (c *ByteSliceCollection) Reduce(fn func(reducer []byte, item []byte) []byte) []byte {
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
func (c *ByteSliceCollection) Find(fn func([]byte) bool) []byte {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// []byte value. The second return value denotes whether the condition
// matched any item or not.
func (c *ByteSliceCollection) FindOk(fn func([]byte) bool) ([]byte, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue []byte
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *ByteSliceCollection) Any(fn func([]byte) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *ByteSliceCollection) All(fn func([]byte) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *ByteSliceCollection) Contains(el []byte) bool {
	for _, item := range c.items {
		if bytes.Equal(item, el) {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
func (c *ByteSliceCollection) Sort(fn func([]byte, []byte) bool) *ByteSliceCollection {
	sort.Slice(c.items, c.lessFunc(fn))
	return c
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *ByteSliceCollection) IsSorted(fn func([]byte, []byte) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *ByteSliceCollection) lessFunc(fn func([]byte, []byte) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse reverses the order of the collection items in place and returns c.
func (c *ByteSliceCollection) Reverse() *ByteSliceCollection {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

// Remove removes the collection item at position idx. Will panic if idx is out
// of bounds.
func (c *ByteSliceCollection) Remove(idx int) *ByteSliceCollection {
	c.items = append(c.items[:idx], c.items[idx+1:]...)
	return c
}

// RemoveItem removes all instances of item from the collection and returns it.
func (c *ByteSliceCollection) RemoveItem(item []byte) *ByteSliceCollection {
	for i, el := range c.items {
		if bytes.Equal(el, item) {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

// InsertItem inserts item into the collection at position idx. Will panic if
// idx is out of bounds.
func (c *ByteSliceCollection) InsertItem(item []byte, idx int) *ByteSliceCollection {
	var zeroValue []byte
	c.items = append(c.items, zeroValue)
	copy(c.items[idx+1:], c.items[idx:])
	c.items[idx] = item
	return c
}

// Cut removes all items between index i and j from the collection and returns
// it. Will panic if i or j is out of bounds of the underlying slice.
func (c *ByteSliceCollection) Cut(i, j int) *ByteSliceCollection {
	c.items = append(c.items[:i], c.items[j:]...)
	return c
}

// Slice replaces the underlying slice of c with the items between i and j and
// returns the collection. Will panic if i or j is out of bounds.
func (c *ByteSliceCollection) Slice(i, j int) *ByteSliceCollection {
	c.items = c.items[i:j]
	return c
}
