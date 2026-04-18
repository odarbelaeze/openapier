package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(HexadecimalTag{})
}

type HexadecimalTag struct{}

func (t HexadecimalTag) Tag() string {
	return "hexadecimal"
}

func (t HexadecimalTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("hexadecimal is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[0-9a-fA-F]*$"),
	}, nil
}

func (t HexadecimalTag) Usage() string {
	return "hexadecimal"
}
