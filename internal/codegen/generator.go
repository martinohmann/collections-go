package codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"text/template"
)

type Parameters struct {
	Package   string
	Name      string
	Exported  string
	ItemType  string
	ZeroValue string
	Immutable bool
	Imports   []*Import
}

func Generate(c *Config) ([]byte, error) {
	t, err := parseTemplate(c.EqualityFunc, Template)
	if err != nil {
		return nil, err
	}

	imports, err := parseImports(c.Imports)
	if err != nil {
		return nil, err
	}

	p := &Parameters{
		Package:   c.Package,
		Name:      c.getCollectionName(),
		ItemType:  c.ItemType,
		Immutable: c.Immutable,
		Imports:   imports,
	}

	return generate(t, p)
}

func generate(t *template.Template, p *Parameters) ([]byte, error) {
	var buf bytes.Buffer

	err := t.Execute(&buf, p)
	if err != nil {
		return nil, err
	}

	code, err := format.Source(buf.Bytes())
	if err != nil {
		return buf.Bytes(), err
	}

	return code, nil
}

func parseImports(pkgPaths []string) ([]*Import, error) {
	imports := make([]*Import, len(pkgPaths))
	for i, pkgPath := range pkgPaths {
		imp, err := parseImport(pkgPath)
		if err != nil {
			return nil, err
		}

		imports[i] = imp
	}

	return imports, nil
}

func parseTemplate(equalityFuncName, text string) (*template.Template, error) {
	return template.New("").
		Funcs(template.FuncMap{
			"title": func(s string) string {
				return strings.Title(s)
			},
			"equals": func(a, b string) string {
				if len(equalityFuncName) > 0 {
					return fmt.Sprintf("%s(%s, %s)", equalityFuncName, a, b)
				}

				return fmt.Sprintf("%s == %s", a, b)
			},
		}).
		Parse(text)
}
