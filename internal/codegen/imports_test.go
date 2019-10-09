package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseImport(t *testing.T) {
	_, err := parseImport("")

	assert.Error(t, err)

	i, err := parseImport("foo=github.com/bar/baz")

	assert.NoError(t, err)
	assert.Equal(t, Import{Alias: "foo", PkgPath: "github.com/bar/baz"}, i)

	i, err = parseImport("github.com/bar/baz")

	assert.NoError(t, err)
	assert.Equal(t, Import{PkgPath: "github.com/bar/baz"}, i)
}

func TestImportString(t *testing.T) {
	i1 := Import{Alias: "foo", PkgPath: "github.com/bar/baz"}

	assert.Equal(t, `foo "github.com/bar/baz"`, i1.String())

	i2 := Import{PkgPath: "github.com/bar/baz"}

	assert.Equal(t, `"github.com/bar/baz"`, i2.String())
}
