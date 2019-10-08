package collections

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPanic(t *testing.T) {
	defer func() {
		r := recover()

		assert.NotNil(t, r)
	}()

	New(nil)
}

func TestNewSafeGeneric(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expectedErr error
	}{
		{
			name:        "nil",
			input:       nil,
			expectedErr: errors.New("cannot create *Generic for nil interface{}"),
		},
		{
			name:        "primitive type",
			input:       "some string",
			expectedErr: errors.New("expected slice type, got string"),
		},
		{
			name:        "struct pointer",
			input:       &FooType{},
			expectedErr: errors.New("expected slice type, got *collections.FooType"),
		},
		{
			name:        "struct",
			input:       FooType{},
			expectedErr: errors.New("expected slice type, got collections.FooType"),
		},
		{
			name:        "typed nil",
			input:       (*FooType)(nil),
			expectedErr: errors.New("expected slice type, got *collections.FooType"),
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
			c, err := SafeNew(test.input)

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

func TestGenericInterface(t *testing.T) {
	c := New([]string{"a"})

	assert.Equal(t, c.Interface(), c.Items())
}

func TestGenericEach(t *testing.T) {
	expected := []string{"a", "b", "c"}

	c := New(expected)

	result := make([]string, 0)

	c.Each(func(item interface{}) {
		result = append(result, item.(string))
	})

	assert.Equal(t, expected, result)
}

func TestGenericIndexOf(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	assert.Equal(t, 2, c.IndexOf("c"))
	assert.Equal(t, -1, c.IndexOf("d"))
}

func TestGenericGet(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	assert.Equal(t, "a", c.First())
	assert.Equal(t, "b", c.Get(1))
	assert.Equal(t, "c", c.Last())

	assert.Equal(t, []string{"a", "b"}, c.FirstN(2))
	assert.Equal(t, []string{"a", "b", "c"}, c.FirstN(4))
	assert.Equal(t, []string{"b", "c"}, c.LastN(2))
	assert.Equal(t, []string{"a", "b", "c"}, c.LastN(4))
}

func TestGenericAppend(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	d := c.Append("d").Append([]interface{}{"e", "f"}...)

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f"}, c.Items())
}

func TestGenericPrepend(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	d := c.Prepend("d").Prepend([]interface{}{"e", "f"}...)

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"e", "f", "d", "a", "b", "c"}, c.Items())
}

func TestGenericCopy(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	d := c.Copy()

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, c.Items(), d.Items())
}

func TestGenericCollectReject(t *testing.T) {
	c := New([]string{"foo", "foobar", "baz"})

	d := c.Collect(func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "foo")
	}).Reject(func(item interface{}) bool {
		return item.(string) == "foo"
	})

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"foobar"}, c.Items())
}

func TestGenericPartition(t *testing.T) {
	c := New([]string{"bar", "foo", "foobar", "baz"})

	d, e := c.Partition(func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "foo")
	})

	if c == d || c == e {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"foo", "foobar"}, d.Items())
	assert.Equal(t, []string{"bar", "baz"}, e.Items())
}

func TestGenericMap(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	d := c.Map(func(item interface{}) interface{} {
		return item.(string) + item.(string)
	})

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"aa", "bb", "cc"}, c.Items())
}

func TestGenericMapIndex(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	d := c.MapIndex(func(item interface{}, i int) interface{} {
		return fmt.Sprintf("%s%d", item.(string), i)
	})

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"a0", "b1", "c2"}, c.Items())
}

func TestGenericReduce(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	result := c.Reduce(func(reducer, item interface{}) interface{} {
		return reducer.(string) + item.(string)
	})

	assert.Equal(t, "abc", result)
}

func TestGenericFind(t *testing.T) {
	c := New([]string{"aa", "bb", "cc"})

	result := c.Find(func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "c")
	})

	assert.Equal(t, "cc", result)

	_, ok := c.FindOk(func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "d")
	})
	assert.False(t, ok)
}

func TestGenericAnyAll(t *testing.T) {
	c := New([]string{"foo", "foobar", "foobarbaz"})

	hasFooPrefix := func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "foo")
	}

	hasBarPrefix := func(item interface{}) bool {
		return strings.HasPrefix(item.(string), "bar")
	}

	hasBarSuffix := func(item interface{}) bool {
		return strings.HasSuffix(item.(string), "bar")
	}

	assert.True(t, c.Any(hasFooPrefix))
	assert.True(t, c.Any(hasBarSuffix))
	assert.False(t, c.Any(hasBarPrefix))

	assert.True(t, c.All(hasFooPrefix))
	assert.False(t, c.All(hasBarSuffix))
	assert.False(t, c.All(hasBarPrefix))
}

func TestGenericContains(t *testing.T) {
	c := New([]string{"a", "b"})

	assert.True(t, c.Contains("a"))
	assert.False(t, c.Contains("c"))
}

func TestGenericSort(t *testing.T) {
	c := New([]string{"z", "b", "y", "a"})

	sortFunc := func(a, b interface{}) bool {
		return a.(string) < b.(string)
	}

	d := c.Sort(sortFunc)

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.True(t, c.IsSorted(sortFunc))
	assert.Equal(t, []string{"a", "b", "y", "z"}, c.Items())
}

func TestGenericReverse(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	d := c.Reverse()

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"c", "b", "a"}, c.Items())
}

func TestGenericRemove(t *testing.T) {
	c := New([]string{"a", "b", "c", "d"})

	d := c.Remove(1)

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"a", "c", "d"}, c.Items())
	assert.Equal(t, []string{"a", "c"}, c.Remove(c.Len()-1).Items())

	d = c.RemoveItem("c")

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"a"}, c.Items())
}

func TestGenericInsertItem(t *testing.T) {
	c := New([]string{"a", "b"})

	d := c.InsertItem("c", 1)

	if c != d {
		t.Fatal("expected pointers to be the same")
	}

	assert.Equal(t, []string{"a", "c", "b"}, c.Items())
	assert.Equal(t, []string{"d", "a", "c", "b"}, c.InsertItem("d", 0).Items())
	assert.Equal(t, []string{"d", "a", "c", "b", "e"}, c.InsertItem("e", c.Len()).Items())
}

func TestGenericCut(t *testing.T) {
	c := New([]string{"a", "b", "c", "d", "e"})

	assert.Equal(t, []string{"a", "c", "d", "e"}, c.Cut(1, 2))

	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, c.Items())
}

func TestGenericSlice(t *testing.T) {
	c := New([]string{"a", "b", "c", "d", "e"})

	assert.Equal(t, []string{"b", "c"}, c.Slice(1, 3))

	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, c.Items())

	c.Slice(1, 3).([]string)[0] = "a"

	assert.Equal(t, []string{"a", "a", "c", "d", "e"}, c.Items())
}
