package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(AlphaNumTag{})
}

type AlphaNumTag struct{}

func (t AlphaNumTag) Tag() string {
	return "alphanum"
}

func (t AlphaNumTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("alphanum is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[a-zA-Z0-9]*$"),
	}, nil
}

func (t AlphaNumTag) Usage() string {
	return "alphanum"
}
