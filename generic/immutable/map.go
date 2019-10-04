package immutable

import (
	"reflect"
)

func (c *Collection) Map(fn func(interface{}) interface{}) *Collection {
	s := c.makeSlice()

	for i := 0; i < c.Len(); i++ {
		s = reflect.Append(s, reflect.ValueOf(fn(c.valueAt(i))))
	}

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Reduce(fn func(reducer interface{}, item interface{}) interface{}) interface{} {
	reducer := c.zeroValue

	for i := 0; i < c.Len(); i++ {
		reducer = fn(reducer, c.valueAt(i))
	}

	return reducer
}
