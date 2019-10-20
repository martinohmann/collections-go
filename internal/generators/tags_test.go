package generators

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/gengo/types"
)

func TestExtractOptionTags(t *testing.T) {
	tests := []struct {
		name        string
		typ         *types.Type
		expected    []typeOptions
		expectedErr error
	}{
		{
			name: "empty",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen:options"},
			},
			expected: nil,
		},
		{
			name: "mutable default",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen:options=mutable"},
			},
			expected: []typeOptions{
				{
					suffix: "collection",
				},
			},
		},
		{
			name: "bool options set",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen:options=immutable,pointer,underlying"},
			},
			expected: []typeOptions{
				{
					immutable:  true,
					pointer:    true,
					underlying: true,
					prefix:     "immutable",
					suffix:     "collection",
				},
			},
		},
		{
			name: "string options set",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen:options=name=foo,out-name=bar,prefix=abc,suffix=123,equality-func=reflect.DeepEqual"},
			},
			expected: []typeOptions{
				{
					name:         "foo",
					outName:      "bar",
					equalityFunc: "reflect.DeepEqual",
				},
			},
		},
		{
			name: "custom prefix and suffix",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen:options=prefix=abc,suffix=123"},
			},
			expected: []typeOptions{
				{
					prefix: "abc",
					suffix: "123",
				},
			},
		},
		{
			name: "ignore prefix and suffix if name is set",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen:options=name=foo,prefix=abc,suffix=123"},
			},
			expected: []typeOptions{
				{
					name: "foo",
				},
			},
		},
		{
			name: "ignore prefix and suffix if noprefix and nosuffix set",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen:options=noprefix,nosuffix"},
			},
			expected: []typeOptions{
				{},
			},
		},
		{
			name: "invalid string option",
			typ: &types.Type{
				Name:         types.Name{Name: "SomeType"},
				CommentLines: []string{"+collection-gen:options=name"},
			},
			expectedErr: errors.New(`malformed collection-gen:options tag for type SomeType: invalid option: "name"`),
		},
		{
			name: "unsupported option",
			typ: &types.Type{
				Name:         types.Name{Name: "SomeType"},
				CommentLines: []string{"+collection-gen:options=name=foo,unsupported-option=val"},
			},
			expectedErr: errors.New(`malformed collection-gen:options tag for type SomeType: unsupported option: "unsupported-option=val"`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts, err := extractOptionTags(test.typ)
			if test.expectedErr != nil {
				require.Error(t, err)
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, opts)
			}
		})
	}
}

func TestEnabledTag(t *testing.T) {
	tests := []struct {
		name     string
		typ      *types.Type
		expected bool
	}{
		{
			name: "missing tag",
			typ:  &types.Type{},
		},
		{
			name: "disabled",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen=false"},
			},
		},
		{
			name: "explicitly enabled",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen=true"},
			},
			expected: true,
		},
		{
			name: "implicitly enabled",
			typ: &types.Type{
				CommentLines: []string{"+collection-gen"},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, extractEnabledTag(test.typ))
		})
	}
}
