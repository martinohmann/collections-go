package immutable

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionLenCap(t *testing.T) {
	s := make([]int, 1, 3)

	c := NewCollection(s)

	assert.Equal(t, 1, c.Len())
	assert.Equal(t, 3, c.Cap())
}

func TestCollectionAppend(t *testing.T) {
	actual := NewCollection([]int{1, 2, 3}).Append([]interface{}{4, 5}...)
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(expected, actual.Items()) {
		t.Fatalf("expected %+v, got %+v", expected, actual.Items())
	}
}

func TestCollectionPrepend(t *testing.T) {
	actual := NewCollection([]int{1, 2, 3}).Prepend([]interface{}{4, 5}...)
	expected := []int{4, 5, 1, 2, 3}

	if !reflect.DeepEqual(expected, actual.Items()) {
		t.Fatalf("expected %+v, got %+v", expected, actual.Items())
	}
}

func TestCopy(t *testing.T) {
	c1 := NewCollection([]int{1, 2, 3})
	c2 := c1.Copy()

	assertEqualItems(t, c1, c2)

	c1 = c1.Append(4)

	assertDifferentItems(t, c1, c2)
}
