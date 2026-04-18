package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(AlphaNumSpaceTag{})
}

type AlphaNumSpaceTag struct{}

func (t AlphaNumSpaceTag) Tag() string {
	return "alphanumspace"
}

func (t AlphaNumSpaceTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("alphanumspace is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[a-zA-Z0-9 ]*$"),
	}, nil
}

func (t AlphaNumSpaceTag) Usage() string {
	return "alphanumspace"
}
