package schema

import (
	"fmt"
	"strings"
)

type Locator struct {
	// Path is the path to the type.
	Path string

	// Package is the package of the type.
	Package string

	// Type is the type of the type.
	Name string
}

func (l Locator) String() string {
	prefix := strings.ReplaceAll(l.Path, "/", "_")
	prefix = strings.ReplaceAll(prefix, ".", "_")
	return fmt.Sprintf("%s:%s.%s", prefix, l.Package, l.Name)
}
