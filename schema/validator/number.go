package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(NumberTag{})
}

type NumberTag struct{}

func (t NumberTag) Tag() string {
	return "number"
}

func (t NumberTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("number is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[0-9]+$"),
	}, nil
}

func (t NumberTag) Usage() string {
	return "number"
}
