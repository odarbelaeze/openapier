package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(HSLATag{})
}

type HSLATag struct{}

func (t HSLATag) Tag() string {
	return "hsla"
}

func (t HSLATag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("hsla is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^hsla\\(.*\\)$"),
	}, nil
}

func (t HSLATag) Usage() string {
	return "hsla"
}
