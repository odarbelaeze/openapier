package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(RGBATag{})
}

type RGBATag struct{}

func (t RGBATag) Tag() string {
	return "rgba"
}

func (t RGBATag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("rgba is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^rgba\\(.*\\)$"),
	}, nil
}

func (t RGBATag) Usage() string {
	return "rgba"
}
