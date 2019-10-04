// Code generated by collections-gen. DO NOT EDIT.

package immutable

import (
	"sort"
)

// Float32Collection is an immutable collection of float32 values.
type Float32Collection struct {
	items []float32
}

// NewFloat32Collection creates a new immutable collection from a slice of float32.
func NewFloat32Collection(items []float32) *Float32Collection {
	return &Float32Collection{items}
}

// Items returns the underlying slice of float32 values used by the
// collection.
func (c *Float32Collection) Items() []float32 {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *Float32Collection) EachIndex(fn func(float32, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *Float32Collection) Each(fn func(float32)) {
	c.EachIndex(func(item float32, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *Float32Collection) IndexOf(el float32) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *Float32Collection) First() float32 {
	return c.Nth(0)
}

// FirstN returns a new collection containing the first n items. Will return
// less than n items if the underlying slice's length is < n.
func (c *Float32Collection) FirstN(n int) *Float32Collection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *Float32Collection) Last() float32 {
	return c.Nth(c.Len() - 1)
}

// LastN returns a new collection containing the last n items. Will return less
// than n items if the underlying slice's length is < n.
func (c *Float32Collection) LastN(n int) *Float32Collection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at idx from the collection. Will panic if the
// underlying slice is shorter than idx+1.
func (c *Float32Collection) Get(idx int) float32 {
	return c.Nth(idx)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than idx+1.
func (c *Float32Collection) Nth(idx int) float32 {
	return c.items[idx]
}

// Len returns the length of the underlying float32 slice.
func (c *Float32Collection) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying float32 slice.
func (c *Float32Collection) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection. The
// initial collection will not be modified.
func (c *Float32Collection) Append(items ...float32) *Float32Collection {
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
}

// Prepend prepends items and returns the collection. The
// initial collection will not be modified.
func (c *Float32Collection) Prepend(items ...float32) *Float32Collection {
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
}

// Copy creates a copy of the collection and the underlying float32 slice.
func (c *Float32Collection) Copy() *Float32Collection {
	s := make([]float32, c.Len(), c.Len())
	copy(s, c.items)

	return NewFloat32Collection(s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The inital collection is not altered.
func (c *Float32Collection) Filter(fn func(float32) bool) *Float32Collection {
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = 0.0
	}

	d.items = s

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The inital collection is not altered.
func (c *Float32Collection) Collect(fn func(float32) bool) *Float32Collection {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The inital collection is not altered.
func (c *Float32Collection) Reject(fn func(float32) bool) *Float32Collection {
	return c.Filter(func(v float32) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *Float32Collection) Partition(fn func(float32) bool) (*Float32Collection, *Float32Collection) {
	lhs := make([]float32, 0, c.Len())
	rhs := make([]float32, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewFloat32Collection(lhs), NewFloat32Collection(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The initial
// collection is not modified.
func (c *Float32Collection) Map(fn func(float32) float32) *Float32Collection {
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item)

	}

	return d
}

func (c *Float32Collection) Reduce(fn func(reducer float32, item float32) float32) float32 {
	var reducer float32

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

func (c *Float32Collection) Find(fn func(float32) bool) float32 {
	item, _ := c.FindOk(fn)

	return item
}

func (c *Float32Collection) FindOk(fn func(float32) bool) (float32, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	return 0.0, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *Float32Collection) Any(fn func(float32) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *Float32Collection) All(fn func(float32) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *Float32Collection) Contains(el float32) bool {
	for _, item := range c.items {
		if item == el {
			return true
		}
	}

	return false
}

func (c *Float32Collection) Sort(fn func(float32, float32) bool) *Float32Collection {
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
}

func (c *Float32Collection) IsSorted(fn func(float32, float32) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *Float32Collection) lessFunc(fn func(float32, float32) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

func (c *Float32Collection) Reverse() *Float32Collection {
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
}

func (c *Float32Collection) Remove(idx int) *Float32Collection {
	d := c.Copy()
	d.items = append(d.items[:idx], d.items[idx+1:]...)
	return d
}

func (c *Float32Collection) RemoveItem(item float32) *Float32Collection {
	d := c.Copy()

	for i, el := range c.items {
		if el == item {
			d.items = append(d.items[:i], d.items[i+1:]...)
		}
	}

	return d
}

func (c *Float32Collection) InsertItem(item float32, idx int) *Float32Collection {
	d := c.Copy()
	d.items = append(d.items, 0.0)
	copy(d.items[idx+1:], d.items[idx:])
	d.items[idx] = item
	return d
}

func (c *Float32Collection) Cut(i, j int) *Float32Collection {
	d := c.Copy()
	d.items = append(d.items[:i], d.items[j:]...)
	return d
}

func (c *Float32Collection) Slice(i, j int) *Float32Collection {
	d := c.Copy()
	d.items = d.items[i:j]
	return d
}
