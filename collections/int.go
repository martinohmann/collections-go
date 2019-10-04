// Code generated by collections-gen. DO NOT EDIT.

package collections

import (
	"sort"
)

// IntCollection is a collection of int values.
type IntCollection struct {
	items []int
}

// NewIntCollection creates a new collection from a slice of int.
func NewIntCollection(items []int) *IntCollection {
	return &IntCollection{items}
}

// Items returns the underlying slice of int values used by the
// collection.
func (c *IntCollection) Items() []int {
	return c.items
}

func (c *IntCollection) EachIndex(fn func(int, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

func (c *IntCollection) Each(fn func(int)) {
	c.EachIndex(func(item int, _ int) {
		fn(item)
	})
}

func (c *IntCollection) IndexOf(el int) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

func (c *IntCollection) First() int {
	return c.Nth(0)
}

func (c *IntCollection) FirstN(n int) *IntCollection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

func (c *IntCollection) Last() int {
	return c.Nth(c.Len() - 1)
}

func (c *IntCollection) LastN(n int) *IntCollection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

func (c *IntCollection) Get(idx int) int {
	return c.Nth(idx)
}

func (c *IntCollection) Nth(idx int) int {
	return c.items[idx]
}

func (c *IntCollection) Len() int {
	return len(c.items)
}

func (c *IntCollection) Cap() int {
	return cap(c.items)
}

func (c *IntCollection) Append(items ...int) *IntCollection {
	c.items = append(c.items, items...)

	return c
}

func (c *IntCollection) Prepend(items ...int) *IntCollection {
	c.items = append(items, c.items...)

	return c
}

func (c *IntCollection) Copy() *IntCollection {
	s := make([]int, c.Len(), c.Len())
	copy(s, c.items)

	return NewIntCollection(s)
}

func (c *IntCollection) Filter(fn func(int) bool) *IntCollection {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = 0
	}

	c.items = s

	return c
}

func (c *IntCollection) Collect(fn func(int) bool) *IntCollection {
	return c.Filter(fn)
}

func (c *IntCollection) Reject(fn func(int) bool) *IntCollection {
	return c.Filter(func(v int) bool {
		return !fn(v)
	})
}

func (c *IntCollection) Partition(fn func(int) bool) (*IntCollection, *IntCollection) {
	lhs := make([]int, 0, c.Len())
	rhs := make([]int, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewIntCollection(lhs), NewIntCollection(rhs)
}

func (c *IntCollection) Map(fn func(int) int) *IntCollection {
	for i, item := range c.items {
		c.items[i] = fn(item)

	}

	return c
}

func (c *IntCollection) Reduce(fn func(reducer int, item int) int) int {
	var reducer int

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

func (c *IntCollection) Find(fn func(int) bool) int {
	item, _ := c.FindOk(fn)

	return item
}

func (c *IntCollection) FindOk(fn func(int) bool) (int, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	return 0, false
}

func (c *IntCollection) Any(fn func(int) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

func (c *IntCollection) All(fn func(int) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

func (c *IntCollection) Contains(el int) bool {
	for _, item := range c.items {
		if item == el {
			return true
		}
	}

	return false
}

func (c *IntCollection) Sort(fn func(int, int) bool) *IntCollection {
	sort.Slice(c.items, c.lessFunc(fn))

	return c
}

func (c *IntCollection) IsSorted(fn func(int, int) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *IntCollection) lessFunc(fn func(int, int) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

func (c *IntCollection) Reverse() *IntCollection {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

func (c *IntCollection) Remove(idx int) *IntCollection {
	c.items = append(c.items[:idx], c.items[idx+1:]...)

	return c
}

func (c *IntCollection) RemoveItem(item int) *IntCollection {
	for i, el := range c.items {
		if el == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

func (c *IntCollection) InsertItem(item int, idx int) *IntCollection {
	c.items = append(c.items, 0)
	copy(c.items[idx+1:], c.items[idx:])
	c.items[idx] = item

	return c
}

func (c *IntCollection) Cut(i, j int) *IntCollection {
	c.items = append(c.items[:i], c.items[j:]...)

	return c
}

func (c *IntCollection) Slice(i, j int) *IntCollection {
	c.items = c.items[i:j]

	return c
}
