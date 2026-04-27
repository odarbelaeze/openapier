package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(ISBNTag{})
}

type ISBNTag struct{}

func (t ISBNTag) Tag() string {
	return "isbn"
}

func (t ISBNTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("isbn is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("isbn"),
	}, nil
}

func (t ISBNTag) Usage() string {
	return "isbn"
}
