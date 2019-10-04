package immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionEach(t *testing.T) {
	input := []int{1, 5, 3}
	actual := make([]int, 0)

	c := NewCollection(input)

	c.Each(func(val interface{}) {
		actual = append(actual, val.(int))
	})

	assert.Equal(t, input, actual)
}

func TestCollectionEachIndex(t *testing.T) {
	input := []int{1, 5, 3}
	actual := make([]int, 0)

	c := NewCollection(input)

	c.EachIndex(func(val interface{}, idx int) {
		actual = append(actual, val.(int)+idx)
	})

	assert.Equal(t, []int{1, 6, 5}, actual)
}

func TestCollectionIndexOf(t *testing.T) {
	c := NewCollection([]string{"d", "b", "z"})

	assert.Equal(t, -1, c.IndexOf("a"))
	assert.Equal(t, 0, c.IndexOf("d"))
	assert.Equal(t, 2, c.IndexOf("z"))
	assert.Equal(t, -1, c.IndexOf(42))
}

func TestCollectionNth(t *testing.T) {
	c := NewCollection([]string{"d", "b", "z"})

	assert.Equal(t, "d", c.First())

	assert.Equal(t, "z", c.Last())

	assert.Equal(t, "b", c.Get(1))
}

func TestInsertItem(t *testing.T) {
	c := NewCollection([]string{"a", "c"})

	d := c.InsertItem("b", 1)

	assert.Equal(t, []string{"a", "b", "c"}, d.Items())

	d = c.InsertItem("b", 0)

	assert.Equal(t, []string{"b", "a", "c"}, d.Items())

	d = c.InsertItem("b", 2)

	assert.Equal(t, []string{"a", "c", "b"}, d.Items())

	assertDifferentItems(t, c, d)
}

func TestRemoveItem(t *testing.T) {
	c := NewCollection([]string{"a", "b", "c", "d"})

	d := c.RemoveItem("b")

	assert.Equal(t, []string{"a", "c", "d"}, d.Items())

	assertDifferentItems(t, c, d)
}

func TestRemove(t *testing.T) {
	c := NewCollection([]string{"a", "b", "c", "d"})

	d := c.Remove(3)

	assert.Equal(t, []string{"a", "b", "c"}, d.Items())

	d = c.Remove(0)

	assert.Equal(t, []string{"b", "c", "d"}, d.Items())

	d = c.Remove(2)

	assert.Equal(t, []string{"a", "b", "d"}, d.Items())

	assertDifferentItems(t, c, d)
}

func TestCut(t *testing.T) {
	c := NewCollection([]string{"a", "b", "c", "d"})

	d := c.Cut(1, 3)

	assert.Equal(t, []string{"a", "d"}, d.Items())

	d = c.Cut(1, 2)

	assert.Equal(t, []string{"a", "c", "d"}, d.Items())

	d = c.Cut(0, c.Len())

	assert.Equal(t, []string{}, d.Items())

	assertDifferentItems(t, c, d)
}

func TestSlice(t *testing.T) {
	c := NewCollection([]string{"a", "b", "c", "d"})

	d := c.Slice(1, 3)

	assert.Equal(t, []string{"b", "c"}, d.Items())

	d = c.Slice(0, c.Len())

	assert.Equal(t, []string{"a", "b", "c", "d"}, d.Items())

	d = c.Slice(1, 2)

	assert.Equal(t, []string{"b"}, d.Items())

	assertDifferentItems(t, c, d)
}

func TestFirstNLastN(t *testing.T) {
	c := NewCollection([]string{"a", "b", "c", "d"})

	d := c.FirstN(3)

	assert.Equal(t, []string{"a", "b", "c"}, d.Items())

	d = c.FirstN(4)

	assert.Equal(t, []string{"a", "b", "c", "d"}, d.Items())

	d = c.FirstN(0)

	assert.Equal(t, []string{}, d.Items())

	d = c.FirstN(5)

	assert.Equal(t, []string{"a", "b", "c", "d"}, d.Items())

	d = c.LastN(5)

	assert.Equal(t, []string{"a", "b", "c", "d"}, d.Items())

	d = c.LastN(0)

	assert.Equal(t, []string{}, d.Items())

	d = c.LastN(4)

	assert.Equal(t, []string{"a", "b", "c", "d"}, d.Items())

	d = c.LastN(3)

	assert.Equal(t, []string{"b", "c", "d"}, d.Items())

	assertDifferentItems(t, c, d)
}
