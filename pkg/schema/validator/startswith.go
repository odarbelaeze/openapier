package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(StartsWithTag{})
}

type StartsWithTag struct{}

func (t StartsWithTag) Tag() string {
	return "startswith"
}

func (t StartsWithTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("startswith is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern(fmt.Sprintf("^%s.*$", value)),
	}, nil
}

func (t StartsWithTag) Usage() string {
	return "startswith=x"
}
