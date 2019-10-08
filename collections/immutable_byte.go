// Code generated by collections-gen. DO NOT EDIT.

package collections

import (
	"sort"
)

// ImmutableByte is an immutable collection of byte values.
type ImmutableByte struct {
	items []byte
}

// NewImmutableByte creates a new immutable collection from a slice of byte.
func NewImmutableByte(items []byte) *ImmutableByte {
	return &ImmutableByte{items}
}

// Interface returns the underlying slice used by the collection as interface{}
// value.
func (c *ImmutableByte) Interface() interface{} {
	return c.items
}

// Items returns the underlying slice of byte values used by the
// collection.
func (c *ImmutableByte) Items() []byte {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *ImmutableByte) EachIndex(fn func(byte, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *ImmutableByte) Each(fn func(byte)) {
	c.EachIndex(func(item byte, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *ImmutableByte) IndexOf(el byte) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *ImmutableByte) First() byte {
	return c.Nth(0)
}

// FirstN returns the first n byte items of the collection. Will
// return less than n items if the underlying slice's length is < n.
func (c *ImmutableByte) FirstN(n int) []byte {
	if n > c.Len() {
		return c.Copy().Items()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *ImmutableByte) Last() byte {
	return c.Nth(c.Len() - 1)
}

// LastN returns the last n byte items of the collection. Will return
// less than n items if the underlying slice's length is < n.
func (c *ImmutableByte) LastN(n int) []byte {
	if c.Len()-n < 0 {
		return c.Copy().Items()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at pos from the collection. Will panic if the
// underlying slice is shorter than pos+1.
func (c *ImmutableByte) Get(pos int) byte {
	return c.Nth(pos)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than pos+1.
func (c *ImmutableByte) Nth(pos int) byte {
	return c.items[pos]
}

// Len returns the length of the underlying byte slice.
func (c *ImmutableByte) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying byte slice.
func (c *ImmutableByte) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableByte) Append(items ...byte) *ImmutableByte {
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
}

// Prepend prepends items and returns the collection. The
// original collection will not be modified.
func (c *ImmutableByte) Prepend(items ...byte) *ImmutableByte {
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
}

// Copy creates a copy of the collection and the underlying byte slice.
func (c *ImmutableByte) Copy() *ImmutableByte {
	s := make([]byte, c.Len(), c.Len())
	copy(s, c.items)

	return NewImmutableByte(s)
}

// Filter collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableByte) Filter(fn func(byte) bool) *ImmutableByte {
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	var zeroValue byte

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = zeroValue
	}

	d.items = s

	return d
}

// Collect collects all items for which fn evaluates to true into a new
// collection. The original collection is not altered.
func (c *ImmutableByte) Collect(fn func(byte) bool) *ImmutableByte {
	return c.Filter(fn)
}

// Reject collects all items for which fn evaluates to false into a new
// collection. The original collection is not altered.
func (c *ImmutableByte) Reject(fn func(byte) bool) *ImmutableByte {
	return c.Filter(func(v byte) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *ImmutableByte) Partition(fn func(byte) bool) (*ImmutableByte, *ImmutableByte) {
	lhs := make([]byte, 0, c.Len())
	rhs := make([]byte, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return NewImmutableByte(lhs), NewImmutableByte(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn. The result is a new collection. The original
// collection is not modified.
func (c *ImmutableByte) Map(fn func(byte) byte) *ImmutableByte {
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item)

	}

	return d
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero byte value on the first invocation.
func (c *ImmutableByte) Reduce(fn func(reducer byte, item byte) byte) byte {
	var reducer byte

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// byte value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *ImmutableByte) Find(fn func(byte) bool) byte {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// byte value. The second return value denotes whether the condition
// matched any item or not.
func (c *ImmutableByte) FindOk(fn func(byte) bool) (byte, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	var zeroValue byte
	return zeroValue, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *ImmutableByte) Any(fn func(byte) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *ImmutableByte) All(fn func(byte) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *ImmutableByte) Contains(el byte) bool {
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
func (c *ImmutableByte) Sort(fn func(byte, byte) bool) *ImmutableByte {
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *ImmutableByte) IsSorted(fn func(byte, byte) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *ImmutableByte) lessFunc(fn func(byte, byte) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

// Reverse copies the collection and returns it with the order of all items
// reversed.
func (c *ImmutableByte) Reverse() *ImmutableByte {
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
}

// Remove removes the collection item at position pos. Will panic if pos is out
// of bounds.
// The result is a new collection, the original is not modified.
func (c *ImmutableByte) Remove(pos int) *ImmutableByte {
	d := c.Copy()
	d.items = append(d.items[:pos], d.items[pos+1:]...)
	return d
}

// RemoveItem removes all instances of item from the collection and returns it.
// The result is a new collection, the original is not modified.
func (c *ImmutableByte) RemoveItem(item byte) *ImmutableByte {
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
func (c *ImmutableByte) InsertItem(item byte, pos int) *ImmutableByte {
	var zeroValue byte
	d := c.Copy()
	d.items = append(d.items, zeroValue)
	copy(d.items[pos+1:], d.items[pos:])
	d.items[pos] = item
	return d
}

// Cut returns a copy of the underlying byte slice with the items
// between index i and j removed. Will panic if i or j is out of bounds of the
// underlying slice.
func (c *ImmutableByte) Cut(i, j int) []byte {
	d := c.Copy()
	return append(d.items[:i], d.items[j:]...)
}

// Slice returns the byte items between slice index i and j. Will
// panic if i or j is out of bounds.
func (c *ImmutableByte) Slice(i, j int) []byte {
	return c.Copy().items[i:j]
}
