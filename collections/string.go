// Code generated by collections-gen. DO NOT EDIT.

package collections

import (
	"sort"
)

// String is a collection of string values.
type String struct {
	items []string
}

// NewString creates a new collection from a slice of string.
func NewString(items []string) *String {
	return &String{items}
}

// Items returns the underlying slice of string values used by the
// collection.
func (c *String) Items() []string {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *String) EachIndex(fn func(string, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *String) Each(fn func(string)) {
	c.EachIndex(func(item string, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *String) IndexOf(el string) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *String) First() string {
	return c.Nth(0)
}

// FirstN returns a new collection containing the first n items. Will return
// less than n items if the underlying slice's length is < n.
func (c *String) FirstN(n int) *String {
	if n > c.Len() {
		return c
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *String) Last() string {
	return c.Nth(c.Len() - 1)
}

// LastN returns a new collection containing the last n items. Will return less
// than n items if the underlying slice's length is < n.
func (c *String) LastN(n int) *String {
	if c.Len()-n < 0 {
		return c
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at idx from the collection. Will panic if the
// underlying slice is shorter than idx+1.
func (c *String) Get(idx int) string {
	return c.Nth(idx)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than idx+1.
func (c *String) Nth(idx int) string {
	return c.items[idx]
}

// Len returns the length of the underlying string slice.
func (c *String) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying string slice.
func (c *String) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.
func (c *String) Append(items ...string) *String {
	c.items = append(c.items, items...)
	return c
}

// Prepend prepends items and returns the collection.
func (c *String) Prepend(items ...string) *String {
	c.items = append(items, c.items...)
	return c
}

// Copy creates a copy of the collection and the underlying string slice.
func (c *String) Copy() *String {
	s := make([]string, c.Len(), c.Len())
	copy(s, c.items)

	return NewString(s)
}

// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *String) Filter(fn func(string) bool) *String {
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue string

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = zeroValue
	}

	c.items = s

	return c
}

// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
func (c *String) Collect(fn func(string) bool) *String {
	return c.Filter(fn)
}

// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
func (c *String) Reject(fn func(string) bool) *String {
	return c.Filter(func(v string) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *String) Partition(fn func(string) bool) (*String, *String) {
	lhs := make([]string, 0, c.Len())
	rhs := make([]string, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewString(lhs), NewString(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.
func (c *String) Map(fn func(string) string) *String {
	for i, item := range c.items {
		c.items[i] = fn(item)

	}

	return c
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero string value on the first invocation.
func (c *String) Reduce(fn func(reducer string, item string) string) string {
	var reducer string

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// string value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *String) Find(fn func(string) bool) string {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// string value. The second return value denotes whether the condition
// matched any item or not.
func (c *String) FindOk(fn func(string) bool) (string, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue string
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *String) Any(fn func(string) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *String) All(fn func(string) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *String) Contains(el string) bool {
	for _, item := range c.items {
		if item == el {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
func (c *String) Sort(fn func(string, string) bool) *String {
	sort.Slice(c.items, c.lessFunc(fn))
	return c
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *String) IsSorted(fn func(string, string) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *String) lessFunc(fn func(string, string) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse reverses the order of the collection items in place and returns c.
func (c *String) Reverse() *String {
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
}

// Remove removes the collection item at position idx. Will panic if idx is out
// of bounds.
func (c *String) Remove(idx int) *String {
	c.items = append(c.items[:idx], c.items[idx+1:]...)
	return c
}

// RemoveItem removes all instances of item from the collection and returns it.
func (c *String) RemoveItem(item string) *String {
	for i, el := range c.items {
		if el == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
}

// InsertItem inserts item into the collection at position idx. Will panic if
// idx is out of bounds.
func (c *String) InsertItem(item string, idx int) *String {
	var zeroValue string
	c.items = append(c.items, zeroValue)
	copy(c.items[idx+1:], c.items[idx:])
	c.items[idx] = item
	return c
}

// Cut removes all items between index i and j from the collection and returns
// it. Will panic if i or j is out of bounds of the underlying slice.
func (c *String) Cut(i, j int) *String {
	c.items = append(c.items[:i], c.items[j:]...)
	return c
}

// Slice replaces the underlying slice of c with the items between i and j and
// returns the collection. Will panic if i or j is out of bounds.
func (c *String) Slice(i, j int) *String {
	c.items = c.items[i:j]
	return c
}
