package immutable

import "reflect"

func (c *Collection) Filter(fn func(interface{}) bool) *Collection {
	s := c.makeSlice()

	for i := 0; i < c.rval.Len(); i++ {
		v := c.rval.Index(i)

		if fn(v.Interface()) {
			s = reflect.Append(s, v)
		}
	}

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Collect(fn func(interface{}) bool) *Collection {
	return c.Filter(fn)
}

func (c *Collection) Reject(fn func(interface{}) bool) *Collection {
	return c.Filter(func(v interface{}) bool {
		return !fn(v)
	})
}

func (c *Collection) Partition(fn func(interface{}) bool) (*Collection, *Collection) {
	lhs, rhs := c.makeSlice(), c.makeSlice()

	for i := 0; i < c.rval.Len(); i++ {
		v := c.rval.Index(i)

		if fn(v.Interface()) {
			lhs = reflect.Append(lhs, v)
		} else {
			rhs = reflect.Append(rhs, v)
		}
	}

	return newCollection(c.sliceType, lhs, lhs.Interface()),
		newCollection(c.sliceType, rhs, rhs.Interface())
}
