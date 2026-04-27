package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(RGBTag{})
}

type RGBTag struct{}

func (t RGBTag) Tag() string {
	return "rgb"
}

func (t RGBTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("rgb is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^rgb\\(.*\\)$"),
	}, nil
}

func (t RGBTag) Usage() string {
	return "rgb"
}
