// +build !ignore_autogenerated

// Code generated by collection-gen. DO NOT EDIT.

package collections

import (
	"sort"
)

// ImmutableFloat64 is an immutable collection of float64 values.
type ImmutableFloat64 struct {
	items []float64
}

// NewImmutableFloat64 creates a new immutable collection from a slice of float64.
func NewImmutableFloat64(items []float64) *ImmutableFloat64 {
	return &ImmutableFloat64{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *ImmutableFloat64) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of float64 values used by the
// collection.
func (c *ImmutableFloat64) Items() []float64 {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *ImmutableFloat64) EachIndex(fn func(float64, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *ImmutableFloat64) Each(fn func(float64)) {
	c.EachIndex(func(item float64, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *ImmutableFloat64) IndexOf(el float64) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *ImmutableFloat64) First() float64 {
	return c.Nth(0)
}

// FirstN returns the first n float64 items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *ImmutableFloat64) FirstN(n int) []float64 {
	if n > c.Len() {
		return c.Copy().Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *ImmutableFloat64) Last() float64 {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n float64 items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *ImmutableFloat64) LastN(n int) []float64 {
	if c.Len()-n < 0 {
		return c.Copy().Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *ImmutableFloat64) Get(pos int) float64 {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *ImmutableFloat64) Nth(pos int) float64 {
	return c.items[pos]
}

// Len returns the length of the underlying float64 slice.
func (c *ImmutableFloat64) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying float64 slice.
func (c *ImmutableFloat64) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableFloat64) Append(items ...float64) *ImmutableFloat64 {
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
}

// Prepend prepends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableFloat64) Prepend(items ...float64) *ImmutableFloat64 {
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
}

// Copy creates a copy of the collection and the underlying float64 slice.
func (c *ImmutableFloat64) Copy() *ImmutableFloat64 {
	s := make([]float64, c.Len(), c.Len())
	copy(s, c.items)

	return NewImmutableFloat64(s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableFloat64) Filter(fn func(float64) bool) *ImmutableFloat64 {
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue float64

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = zeroValue
	}

	d.items = s

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableFloat64) Collect(fn func(float64) bool) *ImmutableFloat64 {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The original collection is not altered.
func (c *ImmutableFloat64) Reject(fn func(float64) bool) *ImmutableFloat64 {
	return c.Filter(func(v float64) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *ImmutableFloat64) Partition(fn func(float64) bool) (*ImmutableFloat64, *ImmutableFloat64) {
	lhs := make([]float64, 0, c.Len())
	rhs := make([]float64, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewImmutableFloat64(lhs), NewImmutableFloat64(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableFloat64) Map(fn func(float64) float64) *ImmutableFloat64 {
	return c.MapIndex(func(item float64, _ int) float64 {
		return fn(item)
	})
}

// MapIndex calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableFloat64) MapIndex(fn func(float64, int) float64) *ImmutableFloat64 {
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item, i)

	}

	return d
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero float64 value on the first invocation.
func (c *ImmutableFloat64) Reduce(fn func(reducer float64, item float64) float64) float64 {
	var reducer float64

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// float64 value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *ImmutableFloat64) Find(fn func(float64) bool) float64 {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// float64 value. The second return value denotes whether the condition
// matched any item or not.
func (c *ImmutableFloat64) FindOk(fn func(float64) bool) (float64, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue float64
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *ImmutableFloat64) Any(fn func(float64) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *ImmutableFloat64) All(fn func(float64) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *ImmutableFloat64) Contains(el float64) bool {
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
func (c *ImmutableFloat64) Sort(fn func(float64, float64) bool) *ImmutableFloat64 {
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *ImmutableFloat64) IsSorted(fn func(float64, float64) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *ImmutableFloat64) lessFunc(fn func(float64, float64) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *ImmutableFloat64) Reverse() *ImmutableFloat64 {
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableFloat64) Remove(pos int) *ImmutableFloat64 {
	d := c.Copy()
	d.items = append(d.items[:pos], d.items[pos+1:]...)
	return d
}

// RemoveItem removes all instances of item from the collection and returns it.
// The result is a new collection, the original is not modified.
func (c *ImmutableFloat64) RemoveItem(item float64) *ImmutableFloat64 {
	d := c.Copy()

	for i, el := range d.items {
		if el == item {
			d.items = append(d.items[:i], d.items[i+1:]...)
		}
	}

	return d
}

// InsertItem inserts item into the collection at position pos. Will panic if
// pos is out of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableFloat64) InsertItem(item float64, pos int) *ImmutableFloat64 {
	var zeroValue float64
	d := c.Copy()
	d.items = append(d.items, zeroValue)
	copy(d.items[pos+1:], d.items[pos:])
	d.items[pos] = item
	return d
}

// Cut returns a copy of the underlying float64 slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *ImmutableFloat64) Cut(i, j int) []float64 {
	d := c.Copy()
	return append(d.items[:i], d.items[j:]...)
}

// Slice returns the float64 items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *ImmutableFloat64) Slice(i, j int) []float64 {
	return c.Copy().items[i:j]
}
