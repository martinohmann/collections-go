package immutable

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FooType struct {
	Bar int
	Baz string
}

func TestCollectionMap(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		fn          func(interface{}) interface{}
		expected    interface{}
		expectedErr error
	}{
		{
			name:        "nil",
			input:       nil,
			expectedErr: errors.New("cannot create *Collection for nil interface{}"),
		},
		{
			name:  "strings",
			input: []string{"a", "b", "c", "d", "e"},
			fn: func(v interface{}) interface{} {
				return v.(string) + "1"
			},
			expected: []string{"a1", "b1", "c1", "d1", "e1"},
		},
		{
			name:  "ints",
			input: []int{1, 2, 3, 4, 5},
			fn: func(v interface{}) interface{} {
				return v.(int) * 2
			},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name: "FooType",
			input: []FooType{
				{Bar: 1, Baz: "a"},
				{Bar: 2, Baz: "b"},
				{Bar: 3, Baz: "c"},
			},
			fn: func(v interface{}) interface{} {
				foo := v.(FooType)

				foo.Bar++

				return foo
			},
			expected: []FooType{
				{Bar: 2, Baz: "a"},
				{Bar: 3, Baz: "b"},
				{Bar: 4, Baz: "c"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, err := SafeNewCollection(test.input)

			if test.expectedErr != nil {
				require.Error(t, err)
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				c = c.Map(test.fn)

				if !reflect.DeepEqual(test.expected, c.Items()) {
					t.Fatalf("expected: %+v, got: %+v", test.expected, c.Items())
				}
			}
		})
	}
}
