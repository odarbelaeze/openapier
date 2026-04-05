package schema

import (
	"fmt"
	"strings"
)

type locator struct {
	// Path is the path to the type.
	Path string

	// Package is the package of the type.
	Package string

	// Type is the type of the type.
	Name string

	// TypeParams are the type parameters of the type, if any.
	TypeParams []string
}

func (l locator) Prefix() string {
	prefix := strings.ReplaceAll(l.Path, "/", "_")
	prefix = strings.ReplaceAll(prefix, ".", "_")
	return prefix
}

func (l locator) TypeName() string {
	typeParams := ""
	if len(l.TypeParams) > 0 {
		typeParams = fmt.Sprintf("[%s]", strings.Join(l.TypeParams, ","))
	}
	return fmt.Sprintf("%s.%s%s", l.Package, l.Name, typeParams)
}

func (l locator) String() string {
	return fmt.Sprintf("%s:%s", l.Prefix(), l.TypeName())
}
