package templates

var Collection = `
// Code generated by collections-gen. DO NOT EDIT.

package {{.Package}}

import (
	"sort"
)

// {{.Name}} is a{{if .Immutable}}n immutable{{end}} collection of {{.ItemType}} values.
type {{.Name}} struct {
	items []{{.ItemType}}
}

// New{{.Name}} creates a new{{if .Immutable}} immutable{{end}} collection from a slice of {{.ItemType}}.
func New{{.Name}}(items []{{.ItemType}}) *{{.Name}} {
	return &{{.Name}}{items}
}

// Items returns the underlying slice of {{.ItemType}} values used by the
// collection.
func (c *{{.Name}}) Items() []{{.ItemType}} {
	return c.items
}

// EachIndex calls fn for every item in the collection. The slice index of the
// item is passed to fn as the second argument.
func (c *{{.Name}}) EachIndex(fn func({{.ItemType}}, int)) {
	for i, item := range c.items {
		fn(item, i)
	}
}

// Each calls fn for every item in the collection.
func (c *{{.Name}}) Each(fn func({{.ItemType}})) {
	c.EachIndex(func(item {{.ItemType}}, _ int) {
		fn(item)
	})
}

// IndexOf searches for el in the collection and returns the first index where
// el is found. If el is not present in the collection IndexOf will return -1.
func (c *{{.Name}}) IndexOf(el {{.ItemType}}) int {
	for i, item := range c.items {
		if item == el {
			return i
		}
	}

	return -1
}

// First returns the first item from the collection. Will panic if the
// underlying slice is empty.
func (c *{{.Name}}) First() {{.ItemType}} {
	return c.Nth(0)
}

// FirstN returns a new collection containing the first n items. Will return
// less than n items if the underlying slice's length is < n.
func (c *{{.Name}}) FirstN(n int) *{{.Name}} {
	if n > c.Len() {
		n = c.Len()
	}

	return c.Slice(0, n)
}

// Last returns the last item from the collection. Will panic if the underlying
// slice is empty.
func (c *{{.Name}}) Last() {{.ItemType}} {
	return c.Nth(c.Len() - 1)
}

// LastN returns a new collection containing the last n items. Will return less
// than n items if the underlying slice's length is < n.
func (c *{{.Name}}) LastN(n int) *{{.Name}} {
	if c.Len()-n < 0 {
		n = c.Len()
	}

	return c.Slice(c.Len()-n, c.Len())
}

// Get returns the item at idx from the collection. Will panic if the
// underlying slice is shorter than idx+1.
func (c *{{.Name}}) Get(idx int) {{.ItemType}} {
	return c.Nth(idx)
}

// Nth returns the nth item from the collection. Will panic if the underlying
// slice is shorter than idx+1.
func (c *{{.Name}}) Nth(idx int) {{.ItemType}} {
	return c.items[idx]
}

// Len returns the length of the underlying {{.ItemType}} slice.
func (c *{{.Name}}) Len() int {
	return len(c.items)
}

// Cap returns the capacity of the underlying {{.ItemType}} slice.
func (c *{{.Name}}) Cap() int {
	return cap(c.items)
}

// Append appends items and returns the collection.{{if .Immutable}} The
// initial collection will not be modified.{{end}}
func (c *{{.Name}}) Append(items ...{{.ItemType}}) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	d.items = append(d.items, items...)
	return d
{{ else -}}
	c.items = append(c.items, items...)
	return c
{{ end -}}
}

// Prepend prepends items and returns the collection.{{if .Immutable}} The
// initial collection will not be modified.{{end}}
func (c *{{.Name}}) Prepend(items ...{{.ItemType}}) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	d.items = append(items, d.items...)
	return d
{{ else -}}
	c.items = append(items, c.items...)
	return c
{{ end -}}
}

// Copy creates a copy of the collection and the underlying {{.ItemType}} slice.
func (c *{{.Name}}) Copy() *{{.Name}} {
	s := make([]{{.ItemType}}, c.Len(), c.Len())
	copy(s, c.items)

	return New{{.Name}}(s)
}

{{ if .Immutable -}}
// Filter collects all items for which fn evaluates to true into a new
// collection. The inital collection is not altered.
{{- else -}}
// Filter removes all items from the collection for which fn evaluates to
// false and returns c.
{{- end}}
func (c *{{.Name}}) Filter(fn func({{.ItemType}}) bool) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	s := d.items[:0]

	for _, item := range d.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	for i := len(s); i < len(d.items); i++ {
		d.items[i] = {{.ZeroValue}}
	}

	d.items = s

	return d
{{ else -}}
	s := c.items[:0]

	for _, item := range c.items {
		if fn(item) {
			s = append(s, item)
		}
	}

	for i := len(s); i < len(c.items); i++ {
		c.items[i] = {{.ZeroValue}}
	}

	c.items = s

	return c
{{ end -}}
}

{{ if .Immutable -}}
// Collect collects all items for which fn evaluates to true into a new
// collection. The inital collection is not altered.
{{- else -}}
// Collect removes all items from the collection for which fn evaluates to
// false and returns c.
{{- end }}
func (c *{{.Name}}) Collect(fn func({{.ItemType}}) bool) *{{.Name}} {
	return c.Filter(fn)
}

{{ if .Immutable -}}
// Reject collects all items for which fn evaluates to false into a new
// collection. The inital collection is not altered.
{{- else -}}
// Reject removes all items from the collection for which fn evaluates to
// true and returns c.
{{- end }}
func (c *{{.Name}}) Reject(fn func({{.ItemType}}) bool) *{{.Name}} {
	return c.Filter(func(v {{.ItemType}}) bool {
		return !fn(v)
	})
}

// Partition partitions the collection into two new collections. The first
// collection contains all items where fn evaluates to true, the second one all
// items where fn evaluates to false.
func (c *{{.Name}}) Partition(fn func({{.ItemType}}) bool) (*{{.Name}}, *{{.Name}}) {
	lhs := make([]{{.ItemType}}, 0, c.Len())
	rhs := make([]{{.ItemType}}, 0, c.Len())

	for _, item := range c.items {
		if fn(item) {
			lhs = append(lhs, item)
		} else {
			rhs = append(rhs, item)
		}
	}

	return New{{.Name}}(lhs), New{{.Name}}(rhs)
}

// Map calls fn for each item in the collection an replaces its value with the
// result of fn.{{if .Immutable}} The result is a new collection. The initial
// collection is not modified.{{end}}
func (c *{{.Name}}) Map(fn func({{.ItemType}}) {{.ItemType}}) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()

	for i, item := range d.items {
		d.items[i] = fn(item)

	}

	return d
{{ else -}}
	for i, item := range c.items {
		c.items[i] = fn(item)

	}

	return c
{{ end -}}
}

// Reduce calls fn for each item in c and reduces the result into reducer. The
// reducer contains the value returned by the call to fn for the previous item.
// Reducer will be the zero {{.ItemType}} value on the first invocation.
func (c *{{.Name}}) Reduce(fn func(reducer {{.ItemType}}, item {{.ItemType}}) {{.ItemType}}) {{.ItemType}} {
	var reducer {{.ItemType}}

	for _, item := range c.items {
		reducer = fn(reducer, item)
	}

	return reducer
}

// Find returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, Find will return the zero
// {{.ItemType}} value. If you need to distinguish zero values from a condition
// that did not match any item consider using FindOk instead.
func (c *{{.Name}}) Find(fn func({{.ItemType}}) bool) {{.ItemType}} {
	item, _ := c.FindOk(fn)

	return item
}

// FindOk returns the first item for which fn evaluates to true. If the
// collection does not contain a matching item, FindOk will return the zero
// {{.ItemType}} value. The second return value denotes whether the condition
// matched any item or not.
func (c *{{.Name}}) FindOk(fn func({{.ItemType}}) bool) ({{.ItemType}}, bool) {
	for _, item := range c.items {
		if fn(item) {
			return item, true
		}
	}

	return {{.ZeroValue}}, false
}

// Any returns true as soon as fn evaluates to true for one item in c.
func (c *{{.Name}}) Any(fn func({{.ItemType}}) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}

	return false
}

// All returns true if fn evaluates to true for all items in c.
func (c *{{.Name}}) All(fn func({{.ItemType}}) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}

	return true
}

// Contains returns true if the collection contains el.
func (c *{{.Name}}) Contains(el {{.ItemType}}) bool {
	for _, item := range c.items {
		if item == el {
			return true
		}
	}

	return false
}

// Sort sorts the collection using the passed in comparator func.
{{- if .Immutable }}
// The result will be a copy of c which is sorted, the original collection is
// not altered.
{{- end }}
func (c *{{.Name}}) Sort(fn func({{.ItemType}}, {{.ItemType}}) bool) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	sort.Slice(d.items, d.lessFunc(fn))
	return d
{{ else -}}
	sort.Slice(c.items, c.lessFunc(fn))
	return c
{{ end -}}
}

// IsSorted returns true if the collection is sorted in the order defined by
// the passed in comparator func.
func (c *{{.Name}}) IsSorted(fn func({{.ItemType}}, {{.ItemType}}) bool) bool {
	return sort.SliceIsSorted(c.items, c.lessFunc(fn))
}

func (c *{{.Name}}) lessFunc(fn func({{.ItemType}}, {{.ItemType}}) bool) func(int, int) bool {
	return func(i, j int) bool {
		return fn(c.items[i], c.items[j])
	}
}

{{ if .Immutable -}}
// Reverse copies the collection and returns it with the order of all items
// reversed.
{{- else -}}
// Reverse reverses the order of the collection items in place and returns c.
{{- end}}
func (c *{{.Name}}) Reverse() *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	for l, r := 0, len(d.items)-1; l < r; l, r = l+1, r-1 {
		d.items[l], d.items[r] = d.items[r], d.items[l]
	}

	return d
{{ else -}}
	for l, r := 0, len(c.items)-1; l < r; l, r = l+1, r-1 {
		c.items[l], c.items[r] = c.items[r], c.items[l]
	}

	return c
{{ end -}}
}

// Remove removes the collection item at position idx. Will panic if idx is out
// of bounds.
{{- if .Immutable }}
// The result is a new collection, the original is not modified.
{{- end }}
func (c *{{.Name}}) Remove(idx int) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	d.items = append(d.items[:idx], d.items[idx+1:]...)
	return d
{{ else -}}
	c.items = append(c.items[:idx], c.items[idx+1:]...)
	return c
{{ end -}}
}

// RemoveItem removes all instances of item from the collection and returns it.
{{- if .Immutable }}
// The result is a new collection, the original is not modified.
{{- end }}
func (c *{{.Name}}) RemoveItem(item {{.ItemType}}) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()

	for i, el := range c.items {
		if el == item {
			d.items = append(d.items[:i], d.items[i+1:]...)
		}
	}

	return d
{{ else -}}
	for i, el := range c.items {
		if el == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
		}
	}

	return c
{{ end -}}
}

// InsertItem inserts item into the collection at position idx. Will panic if
// idx is out of bounds.
{{- if .Immutable }}
// The result is a new collection, the original is not modified.
{{- end }}
func (c *{{.Name}}) InsertItem(item {{.ItemType}}, idx int) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	d.items = append(d.items, {{.ZeroValue}})
	copy(d.items[idx+1:], d.items[idx:])
	d.items[idx] = item
	return d
{{ else -}}
	c.items = append(c.items, {{.ZeroValue}})
	copy(c.items[idx+1:], c.items[idx:])
	c.items[idx] = item
	return c
{{ end -}}
}

// Cut removes all items between index i and j from the collection and returns
// it. Will panic if i or j is out of bounds of the underlying slice.
{{- if .Immutable }}
// The result is a new collection, the original is not modified.
{{- end }}
func (c *{{.Name}}) Cut(i, j int) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	d.items = append(d.items[:i], d.items[j:]...)
	return d
{{ else -}}
	c.items = append(c.items[:i], c.items[j:]...)
	return c
{{ end -}}
}

// Slice replaces the underlying slice of c with the items between i and j and
// returns the collection. Will panic if i or j is out of bounds.
{{- if .Immutable }}
// The result is a new collection, the original is not modified.
{{- end }}
func (c *{{.Name}}) Slice(i, j int) *{{.Name}} {
{{ if .Immutable -}}
	d := c.Copy()
	d.items = d.items[i:j]
	return d
{{ else -}}
	c.items = c.items[i:j]
	return c
{{ end -}}
}
`
