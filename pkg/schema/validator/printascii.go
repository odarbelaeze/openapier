package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(PrintASCIITag{})
}

type PrintASCIITag struct{}

func (t PrintASCIITag) Tag() string {
	return "printascii"
}

func (t PrintASCIITag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("printascii is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[\x20-\x7E]*$"),
	}, nil
}

func (t PrintASCIITag) Usage() string {
	return "printascii"
}
