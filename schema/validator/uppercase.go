package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(UppercaseTag{})
}

type UppercaseTag struct{}

func (t UppercaseTag) Tag() string {
	return "uppercase"
}

func (t UppercaseTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("uppercase is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[A-Z]*$"),
	}, nil
}

func (t UppercaseTag) Usage() string {
	return "uppercase"
}
