package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(LongitudeTag{})
}

type LongitudeTag struct{}

func (t LongitudeTag) Tag() string {
	return "longitude"
}

func (t LongitudeTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "number" && as != "integer" {
		return nil, fmt.Errorf("longitude is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithMinimum(-180),
		options.WithMaximum(180),
	}, nil
}

func (t LongitudeTag) Usage() string {
	return "longitude"
}
