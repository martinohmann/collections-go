package immutable

import "reflect"

func (c *Collection) EachIndex(fn func(interface{}, int)) {
	for i := 0; i < c.rval.Len(); i++ {
		fn(c.valueAt(i), i)
	}
}

func (c *Collection) Each(fn func(interface{})) {
	c.EachIndex(func(val interface{}, _ int) {
		fn(val)
	})
}

func (c *Collection) IndexOf(el interface{}) int {
	for i := 0; i < c.rval.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return i
		}
	}

	return -1
}

func (c *Collection) First() interface{} {
	return c.Nth(0)
}

func (c *Collection) FirstN(n int) *Collection {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

func (c *Collection) Last() interface{} {
	return c.Nth(c.Len() - 1)
}

func (c *Collection) LastN(n int) *Collection {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

func (c *Collection) Get(idx int) interface{} {
	return c.Nth(idx)
}

func (c *Collection) Nth(idx int) interface{} {
	return c.valueAt(idx)
}

func (c *Collection) Remove(idx int) *Collection {
	s := c.copySlice()

	s = reflect.AppendSlice(s.Slice(0, idx), s.Slice(idx+1, s.Len()))

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) RemoveItem(item interface{}) *Collection {
	s := c.copySlice()

	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), item) {
			s = reflect.AppendSlice(s.Slice(0, i), s.Slice(i+1, s.Len()))
		}
	}

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) InsertItem(item interface{}, idx int) *Collection {
	s := c.copySlice()
	s = reflect.Append(s, reflect.ValueOf(c.zeroValue))
	reflect.Copy(s.Slice(idx+1, s.Len()), s.Slice(idx, s.Len()-1))
	s.Index(idx).Set(reflect.ValueOf(item))

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Cut(i, j int) *Collection {
	s := c.copySlice()
	s = reflect.AppendSlice(s.Slice(0, i), s.Slice(j, s.Len()))

	return newCollection(c.sliceType, s, s.Interface())
}

func (c *Collection) Slice(i, j int) *Collection {
	s := c.copySlice()
	s = c.rval.Slice(i, j)

	return newCollection(c.sliceType, s, s.Interface())
}
