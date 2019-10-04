package immutable

import "reflect"

func (c *Collection) Find(fn func(interface{}) bool) interface{} {
	item, _ := c.FindOk(fn)

	return item
}

func (c *Collection) FindOk(fn func(interface{}) bool) (interface{}, bool) {
	for i := 0; i < c.rval.Len(); i++ {
		item := c.valueAt(i)

		if fn(item) {
			return item, true
		}
	}

	return reflect.Zero(c.elemType).Interface(), false
}

func (c *Collection) Any(fn func(interface{}) bool) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if fn(c.valueAt(i)) {
			return true
		}
	}

	return false
}

func (c *Collection) All(fn func(interface{}) bool) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if !fn(c.valueAt(i)) {
			return false
		}
	}

	return true
}

func (c *Collection) Contains(el interface{}) bool {
	for i := 0; i < c.rval.Len(); i++ {
		if reflect.DeepEqual(c.valueAt(i), el) {
			return true
		}
	}

	return false
}
