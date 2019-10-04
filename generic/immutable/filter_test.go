package immutable

import (
	"reflect"
	"testing"
)

func TestCollectionCollect(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		fn       func(interface{}) bool
		expected interface{}
	}{
		{
			name:  "strings",
			input: []string{"a", "b", "c", "d", "e"},
			fn: func(v interface{}) bool {
				return v.(string) < "c"
			},
			expected: []string{"a", "b"},
		},
		{
			name:  "ints",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int)%2 == 0
			},
			expected: []int{2, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCollection(test.input)
			d := c.Collect(test.fn)
			if !reflect.DeepEqual(test.expected, d.Items()) {
				t.Fatalf("expected: %+v, got: %+v", test.expected, d.Items())
			}

			assertDifferentItems(t, c, d)
		})
	}
}

func TestCollectionReject(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		fn       func(interface{}) bool
		expected interface{}
	}{
		{
			name:  "strings",
			input: []string{"a", "b", "c", "d", "e"},
			fn: func(v interface{}) bool {
				return v.(string) < "c"
			},
			expected: []string{"c", "d", "e"},
		},
		{
			name:  "ints",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int)%2 == 0
			},
			expected: []int{1, 3, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCollection(test.input)
			d := c.Reject(test.fn)
			if !reflect.DeepEqual(test.expected, d.Items()) {
				t.Fatalf("expected: %+v, got: %+v", test.expected, d.Items())
			}

			assertDifferentItems(t, c, d)
		})
	}
}

func TestCollectionPartition(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		fn          func(interface{}) bool
		expectedLHS interface{}
		expectedRHS interface{}
	}{
		{
			name:  "strings",
			input: []string{"a", "b", "c", "d", "e"},
			fn: func(v interface{}) bool {
				return v.(string) < "c"
			},
			expectedLHS: []string{"a", "b"},
			expectedRHS: []string{"c", "d", "e"},
		},
		{
			name:  "ints",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int)%2 == 0
			},
			expectedLHS: []int{2, 4},
			expectedRHS: []int{1, 3, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCollection(test.input)
			d := c.Copy()
			lhs, rhs := c.Partition(test.fn)
			if !reflect.DeepEqual(test.expectedLHS, lhs.Items()) {
				t.Fatalf("expected: %+v, got: %+v", test.expectedLHS, lhs.Items())
			}

			if !reflect.DeepEqual(test.expectedRHS, rhs.Items()) {
				t.Fatalf("expected: %+v, got: %+v", test.expectedRHS, rhs.Items())
			}

			assertEqualItems(t, c, d)
		})
	}
}
