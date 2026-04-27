package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(LatitudeTag{})
}

type LatitudeTag struct{}

func (t LatitudeTag) Tag() string {
	return "latitude"
}

func (t LatitudeTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "number" && as != "integer" {
		return nil, fmt.Errorf("latitude is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithMinimum(-90),
		options.WithMaximum(90),
	}, nil
}

func (t LatitudeTag) Usage() string {
	return "latitude"
}
