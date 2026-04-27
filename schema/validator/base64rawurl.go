package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(Base64RawURLTag{})
}

type Base64RawURLTag struct{}

func (t Base64RawURLTag) Tag() string {
	return "base64rawurl"
}

func (t Base64RawURLTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("base64rawurl is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[0-9a-zA-Z-_]*$"),
	}, nil
}

func (t Base64RawURLTag) Usage() string {
	return "base64rawurl"
}
