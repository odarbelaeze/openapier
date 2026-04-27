package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(JWTTag{})
}

type JWTTag struct{}

func (t JWTTag) Tag() string {
	return "jwt"
}

func (t JWTTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("jwt is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$"),
	}, nil
}

func (t JWTTag) Usage() string {
	return "jwt"
}
