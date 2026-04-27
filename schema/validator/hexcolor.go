package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(HexColorTag{})
}

type HexColorTag struct{}

func (t HexColorTag) Tag() string {
	return "hexcolor"
}

func (t HexColorTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("hexcolor is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"),
	}, nil
}

func (t HexColorTag) Usage() string {
	return "hexcolor"
}
