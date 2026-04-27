package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(E164Tag{})
}

type E164Tag struct{}

func (t E164Tag) Tag() string {
	return "e164"
}

func (t E164Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("e164 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^\\+[1-9]\\d{1,14}$"),
	}, nil
}

func (t E164Tag) Usage() string {
	return "e164"
}
