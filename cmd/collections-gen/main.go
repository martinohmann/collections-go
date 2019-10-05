package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/martinohmann/collections-go/internal/templates"
)

var (
	config = &Config{
		Package: os.Getenv("GOPACKAGE"),
	}

	alphanumericRegexp = regexp.MustCompile("[^a-zA-Z0-9]+")
)

func init() {
	flag.StringVar(&config.Package, "package", config.Package, "The package of the generated file")
	flag.StringVar(&config.Name, "name", config.Name, "The name of the collection type")
	flag.StringVar(&config.ItemType, "item-type", config.ItemType, "The item")
	flag.StringVar(&config.ZeroValue, "zero-value", config.ZeroValue, "The zero value of the item type")
	flag.StringVar(&config.EqualsFunc, "equals-func", config.EqualsFunc, "Custom equality func. Must have signature func(item-type, item-type) bool.")
	flag.BoolVar(&config.Immutable, "immutable", config.Immutable, "If set to true, an immutable collection will be generated")

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "usage: %s [flags] <output-file>\n", os.Args[0])
		flag.PrintDefaults()
	}
}

type Config struct {
	Package    string
	Name       string
	ItemType   string
	ZeroValue  string
	OutputFile string
	Immutable  bool
	EqualsFunc string
}

func (c *Config) validate() error {
	if len(c.Package) == 0 {
		return errors.New("package must not be empty")
	}

	if len(c.ItemType) == 0 {
		return errors.New("item type must not be empty")
	}

	return nil
}

func (c *Config) generateName() string {
	itemType := alphanumericRegexp.ReplaceAllString(c.ItemType, "")

	return fmt.Sprintf("%sCollection", strings.Title(itemType))
}

func (c *Config) complete(args []string) error {
	c.OutputFile = args[0]

	if len(c.Name) == 0 {
		c.Name = c.generateName()
	}

	if len(c.ZeroValue) > 0 {
		return nil
	}

	if c.ItemType[0] == '*' || strings.HasPrefix(c.ItemType, "[]") {
		c.ZeroValue = "nil"
		return nil
	}

	switch c.ItemType {
	case "bool":
		c.ZeroValue = "false"
	case "int", "int8", "int16", "int32", "int64":
		c.ZeroValue = "0"
	case "uint", "uint8", "uint16", "uint32", "uint64":
		c.ZeroValue = "0"
	case "byte", "rune", "uintptr":
		c.ZeroValue = "0"
	case "float32", "float64", "complex64", "complex128":
		c.ZeroValue = "0.0"
	case "string":
		c.ZeroValue = `""`
	case "interface{}", "error":
		c.ZeroValue = `nil`
	}

	if len(c.ZeroValue) == 0 {
		return fmt.Errorf(
			"cannot guess zero value for type %q, please provide it via the --zero-item-value flag",
			c.ItemType,
		)
	}

	return nil
}

func (c *Config) parseTemplate(text string) (*template.Template, error) {
	return template.New("").
		Funcs(template.FuncMap{
			"equals": func(a, b string) string {
				if len(c.EqualsFunc) > 0 {
					return fmt.Sprintf("%s(%s, %s)", c.EqualsFunc, a, b)
				}

				return fmt.Sprintf("%s == %s", a, b)
			},
		}).
		Parse(text)
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	err := config.validate()
	if err != nil {
		panic(err)
	}

	err = config.complete(args)
	if err != nil {
		panic(err)
	}

	config.OutputFile = args[0]

	tpl, err := config.parseTemplate(templates.Collection)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer

	err = tpl.Execute(&buf, config)
	if err != nil {
		panic(err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, buf.String())
		panic(err)
	}

	if config.OutputFile == "-" {
		fmt.Fprintf(os.Stdout, string(p))
		os.Exit(0)
	}

	err = ioutil.WriteFile(config.OutputFile, p, 0666)
	if err != nil {
		panic(err)
	}
}
