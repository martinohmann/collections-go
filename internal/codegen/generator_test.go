package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	tpl := `package {{.Package}}
import (
{{- range .Imports }}
  {{ . }}
{{- end }}
)
	
func {{.Name}}(a, b {{.ItemType}}) bool {
  return {{ equals "a" "b"}}	
}`

	c := &Config{
		Package:      "foo",
		Name:         "Bar",
		ItemType:     "Baz",
		EqualityFunc: "reflect.DeepEqual",
		Imports:      []string{"reflect"},
	}

	buf, err := Generate(c, tpl)

	require.NoError(t, err)

	expected := `package foo

import (
	"reflect"
)

func Bar(a, b Baz) bool {
	return reflect.DeepEqual(a, b)
}
`

	assert.Equal(t, expected, string(buf))
}
