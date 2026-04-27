package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(HSLTag{})
}

type HSLTag struct{}

func (t HSLTag) Tag() string {
	return "hsl"
}

func (t HSLTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("hsl is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^hsl\\(.*\\)$"),
	}, nil
}

func (t HSLTag) Usage() string {
	return "hsl"
}
