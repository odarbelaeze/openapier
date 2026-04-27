package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(ISO31661Alpha3Tag{})
}

type ISO31661Alpha3Tag struct{}

func (t ISO31661Alpha3Tag) Tag() string {
	return "iso3166_1_alpha3"
}

func (t ISO31661Alpha3Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("iso3166_1_alpha3 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^[A-Z]{3}$"),
	}, nil
}

func (t ISO31661Alpha3Tag) Usage() string {
	return "iso3166_1_alpha3"
}
