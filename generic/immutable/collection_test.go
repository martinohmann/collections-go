package immutable

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCollectionPanic(t *testing.T) {
	defer func() {
		r := recover()

		assert.NotNil(t, r)
	}()

	NewCollection(nil)
}

func TestNewSafeCollection(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expectedErr error
	}{
		{
			name:        "nil",
			input:       nil,
			expectedErr: errors.New("cannot create *Collection for nil interface{}"),
		},
		{
			name:        "primitive type",
			input:       "some string",
			expectedErr: errors.New("expected slice type, got string"),
		},
		{
			name:        "struct pointer",
			input:       &FooType{},
			expectedErr: errors.New("expected slice type, got *immutable.FooType"),
		},
		{
			name:        "struct",
			input:       FooType{},
			expectedErr: errors.New("expected slice type, got immutable.FooType"),
		},
		{
			name:        "typed nil",
			input:       (*FooType)(nil),
			expectedErr: errors.New("expected slice type, got *immutable.FooType"),
		},
		{
			name:  "nil slice",
			input: []int(nil),
		},
		{
			name:  "int slice",
			input: []int{1, 2},
		},
		{
			name:  "interface{} slice",
			input: []interface{}{1, 2},
		},
		{
			name:  "FooType slice",
			input: []FooType{},
		},
		{
			name:        "map",
			input:       map[string]interface{}{},
			expectedErr: errors.New("expected slice type, got map[string]interface {}"),
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
				assert.NotNil(t, c)
			}
		})
	}
}

func TestCollectionMethodChain(t *testing.T) {
	c := NewCollection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	d := c.Copy()

	result := c.Map(func(val interface{}) interface{} {
		return val.(int) * 3
	}).Reject(func(val interface{}) bool {
		return val.(int)%2 == 0
	}).Collect(func(val interface{}) bool {
		return val.(int) > 10
	}).Reduce(func(r interface{}, val interface{}) interface{} {
		return r.(int) + val.(int)
	})

	expected := 63

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected: %+v, got: %+v", expected, result)
	}

	assertEqualItems(t, c, d)
}

func assertEqualItems(t *testing.T, c1, c2 *Collection) {
	if !reflect.DeepEqual(c1.Items(), c2.Items()) {
		t.Fatalf("collections have different items, c1: %+v, c2: %+v", c1.Items(), c2.Items())
	}
}

func assertDifferentItems(t *testing.T, c1, c2 *Collection) {
	if reflect.DeepEqual(c1.Items(), c2.Items()) {
		t.Fatalf("collections have same items: %+v", c1.Items())
	}
}
