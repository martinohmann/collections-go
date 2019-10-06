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
