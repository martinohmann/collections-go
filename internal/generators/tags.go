package generators

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/gengo/types"
)

const (
	tagName        = "collection-gen"
	optionsTagName = tagName + ":options"
)

type typeOptions struct {
	pointer      bool
	immutable    bool
	underlying   bool
	equalityFunc string
	name         string
	outName      string
	prefix       string
	suffix       string
}

func extractEnabledTag(t *types.Type) bool {
	values := extractCommentTags(t)[tagName]
	if values == nil {
		return false
	}

	return len(values) > 0 && values[0] != "false"
}

func extractOptionTags(t *types.Type) ([]typeOptions, error) {
	values := extractCommentTags(t)[optionsTagName]
	var optionTags []typeOptions

	for _, v := range values {
		if len(v) == 0 {
			continue
		}

		options := typeOptions{
			suffix: "collection",
		}

		var noPrefix, noSuffix bool

		opts := strings.Split(v, ",")
		for _, opt := range opts {
			var err error

			kv := strings.Split(opt, "=")
			switch kv[0] {
			case "mutable":
				// This is a no-op and just here for convenience.
				continue
			case "immutable":
				options.prefix = "immutable"
				options.immutable = true
			case "pointer":
				options.pointer = true
			case "underlying":
				options.underlying = true
			case "noprefix":
				noPrefix = true
			case "nosuffix":
				noSuffix = true
			case "equality-func":
				options.equalityFunc, err = optionValue(opt)
			case "name":
				options.name, err = optionValue(opt)
			case "out-name":
				options.outName, err = optionValue(opt)
			case "prefix":
				options.prefix, err = optionValue(opt)
			case "suffix":
				options.suffix, err = optionValue(opt)
			default:
				err = fmt.Errorf("unsupported option: %q", opt)
			}

			if err != nil {
				return nil, errors.Wrapf(err, "malformed %s tag for type %s", optionsTagName, t.Name)
			}

			if noPrefix || options.name != "" {
				options.prefix = ""
			}

			if noSuffix || options.name != "" {
				options.suffix = ""
			}
		}

		optionTags = append(optionTags, options)
	}

	return optionTags, nil
}

func optionValue(opt string) (string, error) {
	kv := strings.Split(opt, "=")
	if len(kv) != 2 {
		return "", fmt.Errorf("invalid option: %q", opt)
	}

	return kv[1], nil
}

func extractCommentTags(t *types.Type) map[string][]string {
	comments := append(append([]string{}, t.SecondClosestCommentLines...), t.CommentLines...)
	return types.ExtractCommentTags("+", comments)
}
