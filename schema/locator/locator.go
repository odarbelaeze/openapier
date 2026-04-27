package locator

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

	// TypeParams are the type parameters of the type, if any.
	TypeParams []string
}

func (l Locator) Prefix() string {
	prefix := strings.ReplaceAll(l.Path, "/", "_")
	prefix = strings.ReplaceAll(prefix, ".", "_")
	return prefix
}

func (l Locator) Namespace() string {
	return l.Package
}

func (l Locator) TypeName() string {
	typeParams := ""
	if len(l.TypeParams) > 0 {
		typeParams = fmt.Sprintf("[%s]", strings.Join(l.TypeParams, ","))
	}
	return fmt.Sprintf("%s%s", l.Name, typeParams)
}

func (l Locator) String() string {
	return fmt.Sprintf("%s:%s.%s", l.Prefix(), l.Namespace(), l.TypeName())
}
