package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(SSNTag{})
}

type SSNTag struct{}

func (t SSNTag) Tag() string {
	return "ssn"
}

func (t SSNTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ssn is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^\\d{3}-\\d{2}-\\d{4}$"),
	}, nil
}

func (t SSNTag) Usage() string {
	return "ssn"
}
