package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(AlphaTag{})
}

type AlphaTag struct{}

func (t AlphaTag) Tag() string {
	return "alpha"
}

func (t AlphaTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("alpha is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[a-zA-Z]*$"),
	}, nil
}

func (t AlphaTag) Usage() string {
	return "alpha"
}
