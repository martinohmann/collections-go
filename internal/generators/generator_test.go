package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

func TestRemoveObsoleteAliases(t *testing.T) {
	tests := []struct {
		name     string
		in       []string
		expected []string
	}{
		{
			name:     "simple",
			in:       []string{`"bytes"`},
			expected: []string{`"bytes"`},
		},
		{
			name:     "unquoted",
			in:       []string{`bytes`},
			expected: []string{`bytes`},
		},
		{
			name:     "simple alias",
			in:       []string{`foo "bytes"`},
			expected: []string{`foo "bytes"`},
		},
		{
			name:     "simple, redundant alias",
			in:       []string{`bytes "bytes"`},
			expected: []string{`"bytes"`},
		},
		{
			name:     "deep path",
			in:       []string{`foobar "github.com/martinohmann/collections-go/collections"`},
			expected: []string{`foobar "github.com/martinohmann/collections-go/collections"`},
		},
		{
			name:     "deep path, redundant alias",
			in:       []string{`collections "github.com/martinohmann/collections-go/collections"`},
			expected: []string{`"github.com/martinohmann/collections-go/collections"`},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, removeObsoleteAliases(test.in))
		})
	}
}

func TestSanitizePackageName(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "simple",
			in:       "foo-bar",
			expected: "foobar",
		},
		{
			name:     "simple #2",
			in:       "foo.bar",
			expected: "foobar",
		},
		{
			name:     "simple #3",
			in:       "foo_bar",
			expected: "foo_bar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, sanitizePackageName(test.in))
		})
	}
}

func TestCollectionGen_equalityFuncName(t *testing.T) {
	tests := []struct {
		name            string
		gen             *collectionGen
		typ             *types.Type
		expected        string
		validateImports func(t *testing.T, tracker namer.ImportTracker)
	}{
		{
			name: "builtins do not need an equals func",
			gen: &collectionGen{
				imports: generator.NewImportTracker(),
			},
			typ: &types.Type{Kind: types.Builtin},
		},
		{
			name: "reflect.DeepEqual for slices if no equalityFunc defined",
			gen: &collectionGen{
				imports: generator.NewImportTracker(),
			},
			typ:      &types.Type{Kind: types.Slice},
			expected: "reflect.DeepEqual",
		},
		{
			name: "reflect.DeepEqual aliases with underlying slices if no equalityFunc defined",
			gen: &collectionGen{
				imports: generator.NewImportTracker(),
			},
			typ:      &types.Type{Kind: types.Alias, Underlying: &types.Type{Kind: types.Slice}},
			expected: "reflect.DeepEqual",
		},
		{
			name: "pointers do not require an equalityFunc",
			gen: &collectionGen{
				imports: generator.NewImportTracker(),
				options: typeOptions{pointer: true},
			},
			typ:      &types.Type{Kind: types.Slice},
			expected: "",
		},
		{
			name: "pointers do not require an equalityFunc #2",
			gen: &collectionGen{
				imports: generator.NewImportTracker(),
				options: typeOptions{},
			},
			typ:      &types.Type{Kind: types.Pointer, Elem: &types.Type{Kind: types.Slice}},
			expected: "",
		},
		{
			name: "package local equalityFunc",
			gen: &collectionGen{
				imports: generator.NewImportTracker(),
				options: typeOptions{equalityFunc: "Equal"},
			},
			expected: "Equal",
			validateImports: func(t *testing.T, tracker namer.ImportTracker) {
				lines := tracker.ImportLines()

				assert.Len(t, lines, 0)
			},
		},
		{
			name: "equalityFunc from another package",
			gen: &collectionGen{
				imports: generator.NewImportTracker(),
				options: typeOptions{equalityFunc: "bytes.Equal"},
			},
			expected: "bytes.Equal",
			validateImports: func(t *testing.T, tracker namer.ImportTracker) {
				lines := tracker.ImportLines()

				assert.Len(t, lines, 1)
				assert.Equal(t, `bytes "bytes"`, lines[0])
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &generator.Context{
				Universe: types.Universe{},
			}
			result := test.gen.equalityFuncName(c, test.typ)

			assert.Equal(t, test.expected, result)

			if test.validateImports != nil {
				test.validateImports(t, test.gen.imports)
			}
		})
	}
}
