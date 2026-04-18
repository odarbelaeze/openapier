package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(ISO31661Alpha2Tag{})
}

type ISO31661Alpha2Tag struct{}

func (t ISO31661Alpha2Tag) Tag() string {
	return "iso3166_1_alpha2"
}

func (t ISO31661Alpha2Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("iso3166_1_alpha2 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[A-Z]{2}$"),
	}, nil
}

func (t ISO31661Alpha2Tag) Usage() string {
	return "iso3166_1_alpha2"
}
