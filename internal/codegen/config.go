package codegen

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var nonAlphaRegexp = regexp.MustCompile("[^a-zA-Z0-9_]+")

type Config struct {
	Package      string
	Name         string
	ItemType     string
	Immutable    bool
	EqualityFunc string
	Imports      []string
	OutputFile   string
}

func NewDefaultConfig() *Config {
	return &Config{
		Package: os.Getenv("GOPACKAGE"),
	}
}

func (c *Config) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&c.Package, "package", "p", c.Package, "The package of the generated file")
	cmd.Flags().StringVarP(&c.Name, "name", "n", c.Name, "The name of the collection type")
	cmd.Flags().StringVarP(&c.ItemType, "item-type", "t", c.ItemType, "The item")
	cmd.Flags().StringVarP(&c.EqualityFunc, "equality-func", "e", c.EqualityFunc, "Custom equality func. Must have signature func(item-type, item-type) bool.")
	cmd.Flags().BoolVarP(&c.Immutable, "immutable", "i", c.Immutable, "If set to true, an immutable collection will be generated")
	cmd.Flags().StringSliceVar(&c.Imports, "imports", c.Imports, "Additional imports to add to the generated code. Use this to import types or the equals func from a different package. Format: [alias=]fullPkgPath")
}

func (c *Config) Validate() error {
	if len(c.Package) == 0 {
		return errors.New("package must not be empty")
	}

	if len(c.ItemType) == 0 {
		return errors.New("item type must not be empty")
	}

	if nonAlphaRegexp.MatchString(c.Name) {
		return errors.Errorf("name contains non-alphanumeric characters: %s", c.Name)
	}

	return nil
}

func (c *Config) Complete(args []string) error {
	c.OutputFile = args[0]

	c.ensureEqualityFunc()

	return nil
}

func (c *Config) ensureEqualityFunc() {
	if c.EqualityFunc != "" || (!strings.HasPrefix(c.ItemType, "[]") && !strings.HasPrefix(c.ItemType, "map[")) {
		return
	}

	// This will be slow but it is good enough as default and users can
	// override this explicitly if required.
	c.EqualityFunc = "reflect.DeepEqual"

	if c.Imports == nil {
		c.Imports = make([]string, 0, 1)
	}

	c.Imports = append(c.Imports, "reflect")
}

func (c *Config) getCollectionName() string {
	if c.Name != "" {
		return c.Name
	}

	itemType := nonAlphaRegexp.ReplaceAllString(c.ItemType, "")

	return fmt.Sprintf("%sCollection", strings.Title(itemType))
}
