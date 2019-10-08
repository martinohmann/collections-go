// Code generated by collections-gen. DO NOT EDIT.

package collections

import (
	"sort"
)

// Int64 is a collection of int64 values.
type Int64 struct {
	items []int64
}

// NewInt64 creates a new collection from a slice of int64.
func NewInt64(items []int64) *Int64 {
	return &Int64{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *Int64) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of int64 values used by the
// collection.
func (c *Int64) Items() []int64 {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *Int64) EachIndex(fn func(int64, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *Int64) Each(fn func(int64)) {
	c.EachIndex(func(item int64, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *Int64) IndexOf(el int64) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *Int64) First() int64 {
	return c.Nth(0)
}

// FirstN returns the first n int64 items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *Int64) FirstN(n int) []int64 {
	if n > c.Len() {
		return c.Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *Int64) Last() int64 {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n int64 items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *Int64) LastN(n int) []int64 {
	if c.Len()-n < 0 {
		return c.Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *Int64) Get(pos int) int64 {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *Int64) Nth(pos int) int64 {
	return c.items[pos]
}

// Len returns the length of the underlying int64 slice.
func (c *Int64) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying int64 slice.
func (c *Int64) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.
func (c *Int64) Append(items ...int64) *Int64 {
	c.items = append(c.items, items...)
	return c
}

// Prepend prepends items and returns the collection.
func (c *Int64) Prepend(items ...int64) *Int64 {
	c.items = append(items, c.items...)
	return c
}

// Copy creates a copy of the collection and the underlying int64 slice.
func (c *Int64) Copy() *Int64 {
	s := make([]int64, c.Len(), c.Len())
	copy(s, c.items)

	return NewInt64(s)
}

// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *Int64) Filter(fn func(int64) bool) *Int64 {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue int64

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = zeroValue
	}

	c.items = s

	return c
}

// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *Int64) Collect(fn func(int64) bool) *Int64 {
	return c.Filter(fn)
}

// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
func (c *Int64) Reject(fn func(int64) bool) *Int64 {
	return c.Filter(func(v int64) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *Int64) Partition(fn func(int64) bool) (*Int64, *Int64) {
	lhs := make([]int64, 0, c.Len())
	rhs := make([]int64, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewInt64(lhs), NewInt64(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *Int64) Map(fn func(int64) int64) *Int64 {
	for i, item := range c.items {
		c.items[i] = fn(item)

	}

	return c
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero int64 value on the first invocation.
func (c *Int64) Reduce(fn func(reducer int64, item int64) int64) int64 {
	var reducer int64

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// int64 value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *Int64) Find(fn func(int64) bool) int64 {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// int64 value. The second return value denotes whether the condition
// matched any item or not.
func (c *Int64) FindOk(fn func(int64) bool) (int64, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue int64
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *Int64) Any(fn func(int64) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *Int64) All(fn func(int64) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *Int64) Contains(el int64) bool {
	for _, item := range c.items {
		if item == el {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
func (c *Int64) Sort(fn func(int64, int64) bool) *Int64 {
	sort.Slice(c.items, c.lessFunc(fn))
	return c
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *Int64) IsSorted(fn func(int64, int64) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *Int64) lessFunc(fn func(int64, int64) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse reverses the order of the collection items in place and returns c.
func (c *Int64) Reverse() *Int64 {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
func (c *Int64) Remove(pos int) *Int64 {
	c.items = append(c.items[:pos], c.items[pos+1:]...)
	return c
}

// RemoveItem removes all instances of item from the collection and returns it.
func (c *Int64) RemoveItem(item int64) *Int64 {
	for i, el := range c.items {
		if el == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
func (c *Int64) InsertItem(item int64, pos int) *Int64 {
	var zeroValue int64
	c.items = append(c.items, zeroValue)
	copy(c.items[pos+1:], c.items[pos:])
	c.items[pos] = item
	return c
}

// Cut returns a copy of the underlying int64 slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *Int64) Cut(i, j int) []int64 {
	s := make([]int64, 0, c.Cap())
	s = append(s, c.items[:i]...)
	return append(s, c.items[j:]...)
}

// Slice returns the int64 items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *Int64) Slice(i, j int) []int64 {
	return c.items[i:j]
}
