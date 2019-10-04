package immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectionFindOk(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		fn       func(interface{}) bool
		expected interface{}
		found    bool
	}{
		{
			name:  "strings",
			input: []string{"a", "b", "c", "d", "e"},
			fn: func(v interface{}) bool {
				return v.(string) < "c"
			},
			expected: "a",
			found:    true,
		},
		{
			name:  "ints",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int)%2 == 0
			},
			expected: 2,
			found:    true,
		},
		{
			name:  "ints #2",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int) > 5
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, found := NewCollection(test.input).FindOk(test.fn)
			require.Equal(t, test.found, found)
			if found {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestCollectionAny(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		fn       func(interface{}) bool
		expected bool
	}{
		{
			name:  "strings",
			input: []string{"a", "b", "c", "d", "e"},
			fn: func(v interface{}) bool {
				return v.(string) < "c"
			},
			expected: true,
		},
		{
			name:  "ints",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int)%2 == 0
			},
			expected: true,
		},
		{
			name:  "ints #2",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int) > 5
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := NewCollection(test.input).Any(test.fn)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestCollectionAll(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		fn       func(interface{}) bool
		expected bool
	}{
		{
			name:  "strings",
			input: []string{"a", "b", "c", "d", "e"},
			fn: func(v interface{}) bool {
				return v.(string) < "c"
			},
			expected: false,
		},
		{
			name:  "ints",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int)%2 == 0
			},
			expected: false,
		},
		{
			name:  "ints #2",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) bool {
				return v.(int) < 10
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := NewCollection(test.input).All(test.fn)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestCollectionContains(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		val      interface{}
		expected bool
	}{
		{
			name:     "strings",
			input:    []string{"a", "b", "c", "d", "e"},
			val:      "a",
			expected: true,
		},
		{
			name:     "ints",
			input:    []int{1, 2, 3, 4, 5},
			val:      5,
			expected: true,
		},
		{
			name:     "ints #2",
			input:    []int{1, 2, 3, 4, 5},
			val:      "a",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := NewCollection(test.input).Contains(test.val)
			assert.Equal(t, test.expected, actual)
		})
	}
}
