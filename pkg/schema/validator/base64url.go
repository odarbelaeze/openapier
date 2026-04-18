package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(Base64URLTag{})
}

type Base64URLTag struct{}

func (t Base64URLTag) Tag() string {
	return "base64url"
}

func (t Base64URLTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("base64url is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[0-9a-zA-Z-_]*$"),
	}, nil
}

func (t Base64URLTag) Usage() string {
	return "base64url"
}
