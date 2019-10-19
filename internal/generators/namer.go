package generators

import (
	"strings"

	"github.com/iancoleman/strcase"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

func newPublicNamer(opts typeOptions) *strategicNamer {
	return &strategicNamer{
		NameStrategy: newNameStrategy(opts, joiner("", namer.IC, namer.IC)),
		caser:        strcase.ToCamel,
		opts:         opts,
	}
}

func newFileNamer(opts typeOptions) *fileNamer {
	return &fileNamer{
		strategicNamer: &strategicNamer{
			NameStrategy: newNameStrategy(opts, joiner("_", namer.IL, namer.IL)),
			caser:        strcase.ToSnake,
			opts:         opts,
		},
	}
}

type strategicNamer struct {
	*namer.NameStrategy
	caser func(string) string
	opts  typeOptions
}

func (n *strategicNamer) Name(t *types.Type) string {
	name := n.opts.name
	if name == "" {
		name = n.NameStrategy.Name(t)
	}

	return n.caser(name)
}

type fileNamer struct {
	*strategicNamer
}

func (n *fileNamer) Name(t *types.Type) string {
	if n.opts.outName != "" {
		return n.opts.outName
	}

	return n.strategicNamer.Name(t)
}

type rawNamer struct {
	namer.Namer
	opts typeOptions
}

func newRawNamer(pkg string, tracker namer.ImportTracker, opts typeOptions) *rawNamer {
	return &rawNamer{
		Namer: namer.NewRawNamer(pkg, tracker),
		opts:  opts,
	}
}

func (n *rawNamer) Name(t *types.Type) string {
	name := n.Namer.Name(t)
	if n.opts.pointer {
		name = "*" + name
	}

	return name
}

func newNameStrategy(opts typeOptions, join joinFunc) *namer.NameStrategy {
	ns := &namer.NameStrategy{
		Join:        join,
		IgnoreWords: map[string]bool{},
		Prefix:      opts.prefix,
		Suffix:      opts.suffix,
	}

	return ns
}

type joinFunc func(pre string, in []string, post string) string

func joiner(glue string, first, others func(string) string) func(pre string, in []string, post string) string {
	return func(pre string, in []string, post string) string {
		tmp := []string{}

		if pre = others(pre); pre != "" {
			tmp = append(tmp, pre)
		}

		for i := range in {
			tmp = append(tmp, others(in[i]))
		}

		if post = others(post); post != "" {
			tmp = append(tmp, post)
		}

		return first(strings.Join(tmp, glue))
	}
}
