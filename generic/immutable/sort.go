package immutable

import (
	"reflect"
	"sort"
)

func (c *Collection) Sort(fn func(interface{}, interface{}) bool) *Collection {
	d := c.Copy()

	sort.Slice(d.items, d.lessFunc(fn))

	return d
}

func (c *Collection) IsSorted(fn func(interface{}, interface{}) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *Collection) lessFunc(fn func(interface{}, interface{}) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.valueAt(i), c.valueAt(j))
	}
}

func (c *Collection) Reverse() *Collection {
	newSlice := c.makeSlice()

	for i := c.rval.Len() - 1; i >= 0; i-- {
		newSlice = reflect.Append(newSlice, c.rval.Index(i))
	}

	return newCollection(c.sliceType, newSlice, newSlice.Interface())
}
