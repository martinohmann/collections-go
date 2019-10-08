package collections

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImmutableStringInterface(t *testing.T) {
	c := NewImmutableString([]string{"a"})

	assert.Equal(t, c.Interface(), c.Items())
}

func TestImmutableStringEach(t *testing.T) {
	expected := []string{"a", "b", "c"}

	c := NewImmutableString(expected)

	result := make([]string, 0)

	c.Each(func(item string) {
		result = append(result, item)
	})

	assert.Equal(t, expected, result)
}

func TestImmutableStringIndexOf(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	assert.Equal(t, 2, c.IndexOf("c"))
	assert.Equal(t, -1, c.IndexOf("d"))
}

func TestImmutableStringGet(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	assert.Equal(t, "a", c.First())
	assert.Equal(t, "b", c.Get(1))
	assert.Equal(t, "c", c.Last())

	assert.Equal(t, []string{"a", "b"}, c.FirstN(2))
	assert.Equal(t, []string{"a", "b", "c"}, c.FirstN(4))
	assert.Equal(t, []string{"b", "c"}, c.LastN(2))
	assert.Equal(t, []string{"a", "b", "c"}, c.LastN(4))
}

func TestImmutableStringAppend(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	d := c.Append("d").Append([]string{"e", "f"}...)

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f"}, d.Items())
	assert.Equal(t, []string{"a", "b", "c"}, c.Items())
}

func TestImmutableStringPrepend(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	d := c.Prepend("d").Prepend([]string{"e", "f"}...)

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"e", "f", "d", "a", "b", "c"}, d.Items())
	assert.Equal(t, []string{"a", "b", "c"}, c.Items())
}

func TestImmutableStringCopy(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	d := c.Copy()

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, c.Items(), d.Items())
}

func TestImmutableStringCollectReject(t *testing.T) {
	c := NewImmutableString([]string{"foo", "foobar", "baz"})

	d := c.Collect(func(item string) bool {
		return strings.HasPrefix(item, "foo")
	}).Reject(func(item string) bool {
		return item == "foo"
	})

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"foobar"}, d.Items())
	assert.Equal(t, []string{"foo", "foobar", "baz"}, c.Items())
}

func TestImmutableStringPartition(t *testing.T) {
	c := NewImmutableString([]string{"bar", "foo", "foobar", "baz"})

	d, e := c.Partition(func(item string) bool {
		return strings.HasPrefix(item, "foo")
	})

	if c == d || c == e {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"foo", "foobar"}, d.Items())
	assert.Equal(t, []string{"bar", "baz"}, e.Items())
}

func TestImmutableStringMap(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	d := c.Map(func(item string) string {
		return item + item
	})

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"aa", "bb", "cc"}, d.Items())
	assert.Equal(t, []string{"a", "b", "c"}, c.Items())
}

func TestImmutableStringMapIndex(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	d := c.MapIndex(func(item string, i int) string {
		return fmt.Sprintf("%s%d", item, i)
	})

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"a0", "b1", "c2"}, d.Items())
	assert.Equal(t, []string{"a", "b", "c"}, c.Items())
}

func TestImmutableStringReduce(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	result := c.Reduce(func(reducer, item string) string {
		return reducer + item
	})

	assert.Equal(t, "abc", result)
}

func TestImmutableStringFind(t *testing.T) {
	c := NewImmutableString([]string{"aa", "bb", "cc"})

	result := c.Find(func(item string) bool {
		return strings.HasPrefix(item, "c")
	})

	assert.Equal(t, "cc", result)

	_, ok := c.FindOk(func(item string) bool {
		return strings.HasPrefix(item, "d")
	})
	assert.False(t, ok)
}

func TestImmutableStringAnyAll(t *testing.T) {
	c := NewImmutableString([]string{"foo", "foobar", "foobarbaz"})

	hasFooPrefix := func(item string) bool {
		return strings.HasPrefix(item, "foo")
	}

	hasBarPrefix := func(item string) bool {
		return strings.HasPrefix(item, "bar")
	}

	hasBarSuffix := func(item string) bool {
		return strings.HasSuffix(item, "bar")
	}

	assert.True(t, c.Any(hasFooPrefix))
	assert.True(t, c.Any(hasBarSuffix))
	assert.False(t, c.Any(hasBarPrefix))

	assert.True(t, c.All(hasFooPrefix))
	assert.False(t, c.All(hasBarSuffix))
	assert.False(t, c.All(hasBarPrefix))
}

func TestImmutableStringContains(t *testing.T) {
	c := NewImmutableString([]string{"a", "b"})

	assert.True(t, c.Contains("a"))
	assert.False(t, c.Contains("c"))
}

func TestImmutableStringSort(t *testing.T) {
	c := NewImmutableString([]string{"z", "b", "y", "a"})

	sortFunc := func(a, b string) bool {
		return a < b
	}

	d := c.Sort(sortFunc)

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.True(t, d.IsSorted(sortFunc))
	assert.Equal(t, []string{"a", "b", "y", "z"}, d.Items())
	assert.Equal(t, []string{"z", "b", "y", "a"}, c.Items())
}

func TestImmutableStringReverse(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c"})

	d := c.Reverse()

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"c", "b", "a"}, d.Items())
	assert.Equal(t, []string{"a", "b", "c"}, c.Items())
}

func TestImmutableStringRemove(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c", "d"})

	d := c.Remove(1)

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"a", "c", "d"}, d.Items())
	assert.Equal(t, []string{"a", "c"}, d.Remove(d.Len()-1).Items())

	d = c.RemoveItem("c")

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"a", "b", "d"}, d.Items())
	assert.Equal(t, []string{"a", "b", "c", "d"}, c.Items())
}

func TestImmutableStringInsertItem(t *testing.T) {
	c := NewImmutableString([]string{"a", "b"})

	d := c.InsertItem("c", 1)

	if c == d {
		t.Fatal("expected pointers to be different")
	}

	assert.Equal(t, []string{"a", "c", "b"}, d.Items())
	assert.Equal(t, []string{"d", "a", "b"}, c.InsertItem("d", 0).Items())
	assert.Equal(t, []string{"a", "b", "e"}, c.InsertItem("e", c.Len()).Items())
	assert.Equal(t, []string{"a", "b"}, c.Items())
}

func TestImmutableStringCut(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c", "d", "e"})

	assert.Equal(t, []string{"a", "c", "d", "e"}, c.Cut(1, 2))

	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, c.Items())
}

func TestImmutableStringSlice(t *testing.T) {
	c := NewImmutableString([]string{"a", "b", "c", "d", "e"})

	assert.Equal(t, []string{"b", "c"}, c.Slice(1, 3))

	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, c.Items())

	c.Slice(1, 3)[0] = "a"

	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, c.Items())
}
