package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/gengo/types"
)

func TestPublicNamer(t *testing.T) {
	n := newPublicNamer(typeOptions{})

	assert.Equal(t, "String", n.Name(types.String))

	n = newPublicNamer(typeOptions{name: "foobar"})

	assert.Equal(t, "Foobar", n.Name(types.String))

	n = newPublicNamer(typeOptions{prefix: "immutable", suffix: "collection"})

	assert.Equal(t, "ImmutableStringCollection", n.Name(types.String))
}

func TestFileNamer(t *testing.T) {
	n := newFileNamer(typeOptions{})

	assert.Equal(t, "string", n.Name(types.String))

	n = newFileNamer(typeOptions{outName: "foo-bar_baz"})

	assert.Equal(t, "foo-bar_baz", n.Name(types.String))

	n = newFileNamer(typeOptions{name: "FooBar"})

	assert.Equal(t, "foo_bar", n.Name(types.String))

	n = newFileNamer(typeOptions{prefix: "immutable", suffix: "collection"})

	assert.Equal(t, "immutable_string_collection", n.Name(types.String))
}

func TestRawNamer(t *testing.T) {
	n := newRawNamer("", nil, typeOptions{})

	assert.Equal(t, "string", n.Name(types.String))

	n = newRawNamer("", nil, typeOptions{pointer: true})

	assert.Equal(t, "*string", n.Name(types.String))
}
