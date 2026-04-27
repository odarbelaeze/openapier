package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(AlphaSpaceTag{})
}

type AlphaSpaceTag struct{}

func (t AlphaSpaceTag) Tag() string {
	return "alphaspace"
}

func (t AlphaSpaceTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("alphaspace is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[a-zA-Z ]*$"),
	}, nil
}

func (t AlphaSpaceTag) Usage() string {
	return "alphaspace"
}
