package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(EndsWithTag{})
}

type EndsWithTag struct{}

func (t EndsWithTag) Tag() string {
	return "endswith"
}

func (t EndsWithTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("endswith is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern(fmt.Sprintf("^.*%s$", value)),
	}, nil
}

func (t EndsWithTag) Usage() string {
	return "endswith=x"
}
