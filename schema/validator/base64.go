package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(Base64Tag{})
}

type Base64Tag struct{}

func (t Base64Tag) Tag() string {
	return "base64"
}

func (t Base64Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("base64 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("byte"),
	}, nil
}

func (t Base64Tag) Usage() string {
	return "base64"
}
