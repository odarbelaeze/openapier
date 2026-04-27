package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(ISO4217Tag{})
}

type ISO4217Tag struct{}

func (t ISO4217Tag) Tag() string {
	return "iso4217"
}

func (t ISO4217Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("iso4217 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[A-Z]{3}$"),
	}, nil
}

func (t ISO4217Tag) Usage() string {
	return "iso4217"
}
