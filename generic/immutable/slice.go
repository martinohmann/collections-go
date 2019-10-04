package immutable

import "reflect"

func (c *Collection) Len() int {
	return c.rval.Len()
}

func (c *Collection) Cap() int {
	return c.rval.Cap()
}

func (c *Collection) Append(items ...interface{}) *Collection {
	s := c.copySlice()

	for _, item := range items {
		s = reflect.Append(s, reflect.ValueOf(item))
	}

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Prepend(items ...interface{}) *Collection {
	s := reflect.MakeSlice(c.sliceType, 0, c.rval.Len()+len(items))

	for _, item := range items {
		s = reflect.Append(s, reflect.ValueOf(item))
	}

	s = reflect.AppendSlice(s, c.rval)

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Copy() *Collection {
	s := c.copySlice()

	return newCollection(c.sliceType, s, s.Interface())
}
