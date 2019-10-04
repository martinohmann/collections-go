package immutable

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		sortFunc func(interface{}, interface{}) bool
		expected interface{}
	}{
		{
			name:  "strings",
			input: []string{"b", "a", "d", "c", "e"},
			sortFunc: func(a interface{}, b interface{}) bool {
				return a.(string) > b.(string)
			},
			expected: []string{"e", "d", "c", "b", "a"},
		},
		{
			name:  "ints desc",
			input: []int{3, 1, 5, 2, 4},
			sortFunc: func(a interface{}, b interface{}) bool {
				return a.(int) > b.(int)
			},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:  "FooType by Baz",
			input: []FooType{{Baz: "123"}, {Baz: "xyz"}, {Baz: "asdf"}},
			sortFunc: func(a interface{}, b interface{}) bool {
				return a.(FooType).Baz < b.(FooType).Baz
			},
			expected: []FooType{{Baz: "123"}, {Baz: "asdf"}, {Baz: "xyz"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCollection(test.input)
			d := c.Sort(test.sortFunc)
			if !reflect.DeepEqual(test.expected, d.Items()) {
				t.Fatalf("expected: %+v, got: %+v", test.expected, d.Items())
			}

			assertDifferentItems(t, c, d)
		})
	}
}

func TestIsSorted(t *testing.T) {
	sortInts := func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	}

	c := NewCollection([]int{5, 1, 4})

	assert.False(t, c.IsSorted(sortInts))

	d := c.Sort(sortInts)

	assert.True(t, d.IsSorted(sortInts))

	assertDifferentItems(t, c, d)
}

func TestCollectionReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "strings",
			input:    []string{"a", "b", "c", "d", "e"},
			expected: []string{"e", "d", "c", "b", "a"},
		},
		{
			name:     "ints",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCollection(test.input)
			d := c.Reverse()
			if !reflect.DeepEqual(test.expected, d.Items()) {
				t.Fatalf("expected: %+v, got: %+v", test.expected, d.Items())
			}

			assertDifferentItems(t, c, d)
		})
	}
}
