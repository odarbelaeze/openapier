package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(ASCIITag{})
}

type ASCIITag struct{}

func (t ASCIITag) Tag() string {
	return "ascii"
}

func (t ASCIITag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ascii is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[\x00-\x7F]*$"),
	}, nil
}

func (t ASCIITag) Usage() string {
	return "ascii"
}
