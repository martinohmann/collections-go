// Code generated by collections-gen. DO NOT EDIT.

package immutable

import (
	"sort"
)

// IntCollection is an immutable collection of int values.
type IntCollection struct {
	items []int
}

// NewIntCollection creates a new immutable collection from a slice of int.
func NewIntCollection(items []int) *IntCollection {
	return &IntCollection{items}
}

// Items returns the underlying slice of int values used by the
// collection.
func (c *IntCollection) Items() []int {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *IntCollection) EachIndex(fn func(int, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *IntCollection) Each(fn func(int)) {
	c.EachIndex(func(item int, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *IntCollection) IndexOf(el int) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *IntCollection) First() int {
	return c.Nth(0)
}

// FirstN returns a new collection containing the first n items. Will return
// less than n items if the underlying slice's length is < n.
func (c *IntCollection) FirstN(n int) *IntCollection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *IntCollection) Last() int {
	return c.Nth(c.Len() - 1)
}

// LastN returns a new collection containing the last n items. Will return less
// than n items if the underlying slice's length is < n.
func (c *IntCollection) LastN(n int) *IntCollection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at idx from the collection. Will panic if the
// underlying slice is shorter than idx+1.
func (c *IntCollection) Get(idx int) int {
	return c.Nth(idx)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than idx+1.
func (c *IntCollection) Nth(idx int) int {
	return c.items[idx]
}

// Len returns the length of the underlying int slice.
func (c *IntCollection) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying int slice.
func (c *IntCollection) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection. The
// original collection will not be modified.
func (c *IntCollection) Append(items ...int) *IntCollection {
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
}

// Prepend prepends items and returns the collection. The
// original collection will not be modified.
func (c *IntCollection) Prepend(items ...int) *IntCollection {
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
}

// Copy creates a copy of the collection and the underlying int slice.
func (c *IntCollection) Copy() *IntCollection {
	s := make([]int, c.Len(), c.Len())
	copy(s, c.items)

	return NewIntCollection(s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *IntCollection) Filter(fn func(int) bool) *IntCollection {
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue int

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = zeroValue
	}

	d.items = s

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *IntCollection) Collect(fn func(int) bool) *IntCollection {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The original collection is not altered.
func (c *IntCollection) Reject(fn func(int) bool) *IntCollection {
	return c.Filter(func(v int) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
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

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *IntCollection) Map(fn func(int) int) *IntCollection {
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item)

	}

	return d
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero int value on the first invocation.
func (c *IntCollection) Reduce(fn func(reducer int, item int) int) int {
	var reducer int

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// int value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *IntCollection) Find(fn func(int) bool) int {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// int value. The second return value denotes whether the condition
// matched any item or not.
func (c *IntCollection) FindOk(fn func(int) bool) (int, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue int
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *IntCollection) Any(fn func(int) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *IntCollection) All(fn func(int) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *IntCollection) Contains(el int) bool {
	for _, item := range c.items {
		if item == el {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
// The result will be a copy of c which is sorted, the original collection is
// not altered.
func (c *IntCollection) Sort(fn func(int, int) bool) *IntCollection {
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *IntCollection) IsSorted(fn func(int, int) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *IntCollection) lessFunc(fn func(int, int) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *IntCollection) Reverse() *IntCollection {
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
}

// Remove removes the collection item at position idx. Will panic if idx is out
// of bounds.
// The result is a new collection, the original is not modified.
func (c *IntCollection) Remove(idx int) *IntCollection {
	d := c.Copy()
	d.items = append(d.items[:idx], d.items[idx+1:]...)
	return d
}

// RemoveItem removes all instances of item from the collection and returns it.
// The result is a new collection, the original is not modified.
func (c *IntCollection) RemoveItem(item int) *IntCollection {
	d := c.Copy()

	for i, el := range d.items {
		if el == item {
			d.items = append(d.items[:i], d.items[i+1:]...)
		}
	}

	return d
}

// InsertItem inserts item into the collection at position idx. Will panic if
// idx is out of bounds.
// The result is a new collection, the original is not modified.
func (c *IntCollection) InsertItem(item int, idx int) *IntCollection {
	var zeroValue int
	d := c.Copy()
	d.items = append(d.items, zeroValue)
	copy(d.items[idx+1:], d.items[idx:])
	d.items[idx] = item
	return d
}

// Cut removes all items between index i and j from the collection and returns
// it. Will panic if i or j is out of bounds of the underlying slice.
// The result is a new collection, the original is not modified.
func (c *IntCollection) Cut(i, j int) *IntCollection {
	d := c.Copy()
	d.items = append(d.items[:i], d.items[j:]...)
	return d
}

// Slice replaces the underlying slice of c with the items between i and j and
// returns the collection. Will panic if i or j is out of bounds.
// The result is a new collection, the original is not modified.
func (c *IntCollection) Slice(i, j int) *IntCollection {
	d := c.Copy()
	d.items = d.items[i:j]
	return d
}
