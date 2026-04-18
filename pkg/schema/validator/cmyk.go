package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(CMYKTag{})
}

type CMYKTag struct{}

func (t CMYKTag) Tag() string {
	return "cmyk"
}

func (t CMYKTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("cmyk is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^cmyk\\(.*\\)$"),
	}, nil
}

func (t CMYKTag) Usage() string {
	return "cmyk"
}
