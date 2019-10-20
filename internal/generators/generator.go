package generators

import (
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"strings"

	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"k8s.io/klog"
)

// cleanPkgNameRegexp is used to strip unsupported characters from package
// names.
var cleanPkgNameRegexp = regexp.MustCompile("[^a-zA-Z0-9_]+")

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"file":   newFileNamer(typeOptions{}),
		"public": newPublicNamer(typeOptions{}),
		"raw":    newRawNamer("", nil, typeOptions{}),
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

// Packages makes the collection package definition.
func Packages(_ *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		klog.Fatalf("Failed loading boilerplate: %v", err)
	}

	headerText := append([]byte(fmt.Sprintf("// +build !%s\n\n", arguments.GeneratedBuildTag)), boilerplate...)

	return generator.Packages{&generator.DefaultPackage{
		PackageName: sanitizePackageName(filepath.Base(arguments.OutputPackagePath)),
		PackagePath: arguments.OutputPackagePath,
		HeaderText:  headerText,
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			for _, t := range c.Order {
				options, err := extractOptionTags(t)
				if err != nil {
					klog.Fatalf("Failed to parse option tags: %v", err)
				}

				if len(options) == 0 {
					generators = append(generators, &collectionGen{
						DefaultGen: generator.DefaultGen{
							OptionalName: c.Namers["file"].Name(t),
						},
						imports:       generator.NewImportTracker(),
						outputPackage: arguments.OutputPackagePath,
						typeToMatch:   t,
					})
					continue
				}

				for _, opts := range options {
					generators = append(generators, &collectionGen{
						DefaultGen: generator.DefaultGen{
							OptionalName: newFileNamer(opts).Name(t),
						},
						imports:       generator.NewImportTracker(),
						options:       opts,
						outputPackage: arguments.OutputPackagePath,
						typeToMatch:   t,
					})
				}
			}

			return generators
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			return extractEnabledTag(t)
		},
	}}
}

type collectionGen struct {
	generator.DefaultGen

	outputPackage string
	options       typeOptions
	typeToMatch   *types.Type
	imports       namer.ImportTracker
}

// Filter implements the generator.Generator interface.
func (g *collectionGen) Filter(c *generator.Context, t *types.Type) bool {
	return t == g.typeToMatch
}

// Namers implements the generator.Generator interface.
func (g *collectionGen) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"public": newPublicNamer(g.options),
		"raw":    newRawNamer(g.outputPackage, g.imports, g.options),
	}
}

// GenerateType implements the generator.Generator interface.
func (g *collectionGen) Imports(c *generator.Context) []string {
	imports := removeObsoleteAliases(g.imports.ImportLines())
	return append(imports, "reflect", "sort")
}

// GenerateType implements the generator.Generator interface.
func (g *collectionGen) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	if g.options.underlying {
		t = underlying(t)
	}

	klog.Infof("Generating collection for type %s (%s)", t.Name, t.Kind)
	klog.V(1).Infof("with options: %+v", g.options)

	sw := generator.NewSnippetWriter(w, c, "{{", "}}")
	sw.Do(collectionCode, generator.Args{
		"type":         t,
		"immutable":    g.options.immutable,
		"equalityFunc": g.equalityFuncName(c, t),
	})

	return sw.Error()
}

// equalityFuncName builds the name of the func used for equality checks on the
// type the collection will be generated for. The returned name includes the
// possible package name if it is not package local. If no custom equality func
// was defined in the generator options and t is a slice, map or func,
// "reflect.DeepEqual" will be returned. equalityFuncName also ensures that the
// type of the equality func is added to the import tracker so that is
// automatically included in the imports of the generated file. If
// equalityFuncName returns an empty string it means that no equality func is
// needed and "==" can be used.
func (g *collectionGen) equalityFuncName(c *generator.Context, t *types.Type) string {
	equalityFunc := g.options.equalityFunc
	if equalityFunc == "" {
		// if no equalityFunc is defined, set a sane default. Pointers can be
		// compared by pointer equality and do not need a special equality
		// func.
		if g.options.pointer {
			return ""
		}

		switch underlying(t).Kind {
		case types.Slice, types.Map, types.Func:
			return "reflect.DeepEqual"
		}

		return ""
	}

	name := types.ParseFullyQualifiedName(equalityFunc)
	if name.Package == "" {
		return name.Name
	}

	f := c.Universe.Function(name)

	g.imports.AddType(f)

	pkgName := g.imports.LocalNameOf(name.Package)

	return pkgName + "." + name.Name
}

// underlying returns the underlying type if t is an alias. Returns t
// unmodified if it is not an alias.
func underlying(t *types.Type) *types.Type {
	if t.Kind == types.Alias {
		return t.Underlying
	}

	return t
}

// sanitizePackageName removes all non-letter, non-digit and non-underscore
// characters from a package name.
func sanitizePackageName(name string) string {
	return cleanPkgNameRegexp.ReplaceAllString(name, "")
}

// removeObsoleteAliases removes aliases from import lines that are
// unnecessary. An alias is unnecessary if it exactly matches the name of the
// imported package. For example the import line `http "net/http"` will be
// rewritten to `"net/http"`.
func removeObsoleteAliases(lines []string) []string {
	for i, line := range lines {
		aliasedPkgPath := strings.Split(line, " ")
		if len(aliasedPkgPath) == 1 {
			continue
		}

		alias := aliasedPkgPath[0]
		quotedPkgPath := aliasedPkgPath[1]
		pkgPath := quotedPkgPath[1 : len(quotedPkgPath)-1]
		pathParts := strings.Split(pkgPath, "/")

		pkgName := pathParts[len(pathParts)-1]
		if pkgName == alias {
			lines[i] = quotedPkgPath
		}
	}

	return lines
}
