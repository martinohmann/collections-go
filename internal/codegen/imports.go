package codegen

import (
	"errors"
	"fmt"
	"strings"
)

// Import holds information about about the package path of an import and its
// optional alias.
type Import struct {
	PkgPath string
	Alias   string
}

// String returns the formatted import valid for use as an import declaration.
func (i Import) String() string {
	if i.Alias == "" {
		return fmt.Sprintf("%q", i.PkgPath)
	}

	return fmt.Sprintf("%s %q", i.Alias, i.PkgPath)
}

func parseImport(pkgPath string) (Import, error) {
	if pkgPath == "" {
		return Import{}, errors.New("pkgPath cannot be empty")
	}

	parts := strings.SplitN(pkgPath, "=", 2)
	if len(parts) == 2 {
		return Import{PkgPath: parts[1], Alias: parts[0]}, nil
	}

	return Import{PkgPath: pkgPath}, nil
}
