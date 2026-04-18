package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(LowercaseTag{})
}

type LowercaseTag struct{}

func (t LowercaseTag) Tag() string {
	return "lowercase"
}

func (t LowercaseTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("lowercase is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[a-z]*$"),
	}, nil
}

func (t LowercaseTag) Usage() string {
	return "lowercase"
}
