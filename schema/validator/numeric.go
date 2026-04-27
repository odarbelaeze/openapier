package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(NumericTag{})
}

type NumericTag struct{}

func (t NumericTag) Tag() string {
	return "numeric"
}

func (t NumericTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("numeric is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[0-9]*$"),
	}, nil
}

func (t NumericTag) Usage() string {
	return "numeric"
}
