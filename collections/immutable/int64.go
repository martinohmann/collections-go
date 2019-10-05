// Code generated by collections-gen. DO NOT EDIT.

package immutable

import (
	"sort"
)

// Int64Collection is an immutable collection of int64 values.
type Int64Collection struct {
	items []int64
}

// NewInt64Collection creates a new immutable collection from a slice of int64.
func NewInt64Collection(items []int64) *Int64Collection {
	return &Int64Collection{items}
}

// Items returns the underlying slice of int64 values used by the
// collection.
func (c *Int64Collection) Items() []int64 {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *Int64Collection) EachIndex(fn func(int64, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *Int64Collection) Each(fn func(int64)) {
	c.EachIndex(func(item int64, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *Int64Collection) IndexOf(el int64) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *Int64Collection) First() int64 {
	return c.Nth(0)
}

// FirstN returns a new collection containing the first n items. Will return
// less than n items if the underlying slice's length is < n.
func (c *Int64Collection) FirstN(n int) *Int64Collection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *Int64Collection) Last() int64 {
	return c.Nth(c.Len() - 1)
}

// LastN returns a new collection containing the last n items. Will return less
// than n items if the underlying slice's length is < n.
func (c *Int64Collection) LastN(n int) *Int64Collection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at idx from the collection. Will panic if the
// underlying slice is shorter than idx+1.
func (c *Int64Collection) Get(idx int) int64 {
	return c.Nth(idx)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than idx+1.
func (c *Int64Collection) Nth(idx int) int64 {
	return c.items[idx]
}

// Len returns the length of the underlying int64 slice.
func (c *Int64Collection) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying int64 slice.
func (c *Int64Collection) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection. The
// initial collection will not be modified.
func (c *Int64Collection) Append(items ...int64) *Int64Collection {
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
}

// Prepend prepends items and returns the collection. The
// initial collection will not be modified.
func (c *Int64Collection) Prepend(items ...int64) *Int64Collection {
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
}

// Copy creates a copy of the collection and the underlying int64 slice.
func (c *Int64Collection) Copy() *Int64Collection {
	s := make([]int64, c.Len(), c.Len())
	copy(s, c.items)

	return NewInt64Collection(s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The inital collection is not altered.
func (c *Int64Collection) Filter(fn func(int64) bool) *Int64Collection {
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = 0
	}

	d.items = s

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The inital collection is not altered.
func (c *Int64Collection) Collect(fn func(int64) bool) *Int64Collection {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The inital collection is not altered.
func (c *Int64Collection) Reject(fn func(int64) bool) *Int64Collection {
	return c.Filter(func(v int64) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *Int64Collection) Partition(fn func(int64) bool) (*Int64Collection, *Int64Collection) {
	lhs := make([]int64, 0, c.Len())
	rhs := make([]int64, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewInt64Collection(lhs), NewInt64Collection(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The initial
// collection is not modified.
func (c *Int64Collection) Map(fn func(int64) int64) *Int64Collection {
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item)

	}

	return d
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero int64 value on the first invocation.
func (c *Int64Collection) Reduce(fn func(reducer int64, item int64) int64) int64 {
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
func (c *Int64Collection) Find(fn func(int64) bool) int64 {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// int64 value. The second return value denotes whether the condition
// matched any item or not.
func (c *Int64Collection) FindOk(fn func(int64) bool) (int64, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	return 0, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *Int64Collection) Any(fn func(int64) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *Int64Collection) All(fn func(int64) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *Int64Collection) Contains(el int64) bool {
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
func (c *Int64Collection) Sort(fn func(int64, int64) bool) *Int64Collection {
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *Int64Collection) IsSorted(fn func(int64, int64) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *Int64Collection) lessFunc(fn func(int64, int64) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *Int64Collection) Reverse() *Int64Collection {
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
}

// Remove removes the collection item at position idx. Will panic if idx is out
// of bounds.
// The result is a new collection, the original is not modified.
func (c *Int64Collection) Remove(idx int) *Int64Collection {
	d := c.Copy()
	d.items = append(d.items[:idx], d.items[idx+1:]...)
	return d
}

// RemoveItem removes all instances of item from the collection and returns it.
// The result is a new collection, the original is not modified.
func (c *Int64Collection) RemoveItem(item int64) *Int64Collection {
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
func (c *Int64Collection) InsertItem(item int64, idx int) *Int64Collection {
	d := c.Copy()
	d.items = append(d.items, 0)
	copy(d.items[idx+1:], d.items[idx:])
	d.items[idx] = item
	return d
}

// Cut removes all items between index i and j from the collection and returns
// it. Will panic if i or j is out of bounds of the underlying slice.
// The result is a new collection, the original is not modified.
func (c *Int64Collection) Cut(i, j int) *Int64Collection {
	d := c.Copy()
	d.items = append(d.items[:i], d.items[j:]...)
	return d
}

// Slice replaces the underlying slice of c with the items between i and j and
// returns the collection. Will panic if i or j is out of bounds.
// The result is a new collection, the original is not modified.
func (c *Int64Collection) Slice(i, j int) *Int64Collection {
	d := c.Copy()
	d.items = d.items[i:j]
	return d
}
