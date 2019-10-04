// Code generated by collections-gen. DO NOT EDIT.

package internal

import (
	"sort"
)

type Collection struct {
	items []*Type
}

func NewCollection(items []*Type) *Collection {
	return &Collection{items}
}

func (c *Collection) Items() []*Type {
	return c.items
}

func (c *Collection) EachIndex(fn func(*Type, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

func (c *Collection) Each(fn func(*Type)) {
	c.EachIndex(func(item *Type, _ int) {
		fn(item)
	})
}

func (c *Collection) IndexOf(el *Type) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

func (c *Collection) First() *Type {
	return c.Nth(0)
}

func (c *Collection) FirstN(n int) *Collection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

func (c *Collection) Last() *Type {
	return c.Nth(c.Len() - 1)
}

func (c *Collection) LastN(n int) *Collection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

func (c *Collection) Get(idx int) *Type {
	return c.Nth(idx)
}

func (c *Collection) Nth(idx int) *Type {
	return c.items[idx]
}

func (c *Collection) Len() int {
	return len(c.items)
}

func (c *Collection) Cap() int {
	return cap(c.items)
}

func (c *Collection) Append(items ...*Type) *Collection {
	c.items = append(c.items, items...)

	return c
}

func (c *Collection) Prepend(items ...*Type) *Collection {
	c.items = append(items, c.items...)

	return c
}

func (c *Collection) Copy() *Collection {
	s := make([]*Type, c.Len(), c.Len())
	copy(s, c.items)

	return NewCollection(s)
}

func (c *Collection) Filter(fn func(*Type) bool) *Collection {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = nil
	}

	c.items = s

	return c
}

func (c *Collection) Collect(fn func(*Type) bool) *Collection {
	return c.Filter(fn)
}

func (c *Collection) Reject(fn func(*Type) bool) *Collection {
	return c.Filter(func(v *Type) bool {
		return !fn(v)
	})
}

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

func (c *Collection) Map(fn func(*Type) *Type) *Collection {
	for i, item := range c.items {
		c.items[i] = fn(item)

	}

	return c
}

func (c *Collection) Reduce(fn func(reducer *Type, item *Type) *Type) *Type {
	var reducer *Type

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

func (c *Collection) Find(fn func(*Type) bool) *Type {
	item, _ := c.FindOk(fn)

	return item
}

func (c *Collection) FindOk(fn func(*Type) bool) (*Type, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	return nil, false
}

func (c *Collection) Any(fn func(*Type) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

func (c *Collection) All(fn func(*Type) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

func (c *Collection) Contains(el *Type) bool {
	for _, item := range c.items {
		if item == el {
			return true
		}
	}

	return false
}

func (c *Collection) Sort(fn func(*Type, *Type) bool) *Collection {
	sort.Slice(c.items, c.lessFunc(fn))

	return c
}

func (c *Collection) IsSorted(fn func(*Type, *Type) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *Collection) lessFunc(fn func(*Type, *Type) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

func (c *Collection) Reverse() *Collection {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

func (c *Collection) Remove(idx int) *Collection {
	c.items = append(c.items[:idx], c.items[idx+1:]...)

	return c
}

func (c *Collection) RemoveItem(item *Type) *Collection {
	for i, el := range c.items {
		if el == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

func (c *Collection) InsertItem(item *Type, idx int) *Collection {
	c.items = append(c.items, nil)
	copy(c.items[idx+1:], c.items[idx:])
	c.items[idx] = item

	return c
}

func (c *Collection) Cut(i, j int) *Collection {
	c.items = append(c.items[:i], c.items[j:]...)

	return c
}

func (c *Collection) Slice(i, j int) *Collection {
	c.items = c.items[i:j]

	return c
}
