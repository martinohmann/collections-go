// Code generated by collections-gen. DO NOT EDIT.

package collections

import (
	"sort"
)

// BoolCollection is a collection of bool values.
type BoolCollection struct {
	items []bool
}

// NewBoolCollection creates a new collection from a slice of bool.
func NewBoolCollection(items []bool) *BoolCollection {
	return &BoolCollection{items}
}

// Items returns the underlying slice of bool values used by the
// collection.
func (c *BoolCollection) Items() []bool {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *BoolCollection) EachIndex(fn func(bool, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *BoolCollection) Each(fn func(bool)) {
	c.EachIndex(func(item bool, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *BoolCollection) IndexOf(el bool) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *BoolCollection) First() bool {
	return c.Nth(0)
}

// FirstN returns a new collection containing the first n items. Will return
// less than n items if the underlying slice's length is < n.
func (c *BoolCollection) FirstN(n int) *BoolCollection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *BoolCollection) Last() bool {
	return c.Nth(c.Len() - 1)
}

// LastN returns a new collection containing the last n items. Will return less
// than n items if the underlying slice's length is < n.
func (c *BoolCollection) LastN(n int) *BoolCollection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at idx from the collection. Will panic if the
// underlying slice is shorter than idx+1.
func (c *BoolCollection) Get(idx int) bool {
	return c.Nth(idx)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than idx+1.
func (c *BoolCollection) Nth(idx int) bool {
	return c.items[idx]
}

// Len returns the length of the underlying bool slice.
func (c *BoolCollection) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying bool slice.
func (c *BoolCollection) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.
func (c *BoolCollection) Append(items ...bool) *BoolCollection {
	c.items = append(c.items, items...)
	return c
}

// Prepend prepends items and returns the collection.
func (c *BoolCollection) Prepend(items ...bool) *BoolCollection {
	c.items = append(items, c.items...)
	return c
}

// Copy creates a copy of the collection and the underlying bool slice.
func (c *BoolCollection) Copy() *BoolCollection {
	s := make([]bool, c.Len(), c.Len())
	copy(s, c.items)

	return NewBoolCollection(s)
}

// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *BoolCollection) Filter(fn func(bool) bool) *BoolCollection {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = false
	}

	c.items = s

	return c
}

// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *BoolCollection) Collect(fn func(bool) bool) *BoolCollection {
	return c.Filter(fn)
}

// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
func (c *BoolCollection) Reject(fn func(bool) bool) *BoolCollection {
	return c.Filter(func(v bool) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *BoolCollection) Partition(fn func(bool) bool) (*BoolCollection, *BoolCollection) {
	lhs := make([]bool, 0, c.Len())
	rhs := make([]bool, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewBoolCollection(lhs), NewBoolCollection(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *BoolCollection) Map(fn func(bool) bool) *BoolCollection {
	for i, item := range c.items {
		c.items[i] = fn(item)

	}

	return c
}

func (c *BoolCollection) Reduce(fn func(reducer bool, item bool) bool) bool {
	var reducer bool

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

func (c *BoolCollection) Find(fn func(bool) bool) bool {
	item, _ := c.FindOk(fn)

	return item
}

func (c *BoolCollection) FindOk(fn func(bool) bool) (bool, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	return false, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *BoolCollection) Any(fn func(bool) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *BoolCollection) All(fn func(bool) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *BoolCollection) Contains(el bool) bool {
	for _, item := range c.items {
		if item == el {
			return true
		}
	}

	return false
}

func (c *BoolCollection) Sort(fn func(bool, bool) bool) *BoolCollection {
	sort.Slice(c.items, c.lessFunc(fn))
	return c
}

func (c *BoolCollection) IsSorted(fn func(bool, bool) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *BoolCollection) lessFunc(fn func(bool, bool) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

func (c *BoolCollection) Reverse() *BoolCollection {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

func (c *BoolCollection) Remove(idx int) *BoolCollection {
	c.items = append(c.items[:idx], c.items[idx+1:]...)
	return c
}

func (c *BoolCollection) RemoveItem(item bool) *BoolCollection {
	for i, el := range c.items {
		if el == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

func (c *BoolCollection) InsertItem(item bool, idx int) *BoolCollection {
	c.items = append(c.items, false)
	copy(c.items[idx+1:], c.items[idx:])
	c.items[idx] = item
	return c
}

func (c *BoolCollection) Cut(i, j int) *BoolCollection {
	c.items = append(c.items[:i], c.items[j:]...)
	return c
}

func (c *BoolCollection) Slice(i, j int) *BoolCollection {
	c.items = c.items[i:j]
	return c
}
