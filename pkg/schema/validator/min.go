package validator

import (
	"fmt"
	"strconv"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(MinTag{})
}

type MinTag struct{}

func (t MinTag) Tag() string {
	return "min"
}

func (t MinTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	m, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid min value: %s", value)
	}

	switch as {
	case "integer", "number":
		return []options.SchemaOption{
			options.WithMinimum(int(m)),
		}, nil
	case "string":
		return []options.SchemaOption{
			options.WithMinLength(int(m)),
		}, nil
	case "array":
		return []options.SchemaOption{
			options.WithMinItems(int(m)),
		}, nil
	case "object":
		return []options.SchemaOption{
			options.WithMinProperties(int(m)),
		}, nil
	default:
		return nil, fmt.Errorf("min is not supported for %s", as)
	}
}

func (t MinTag) Usage() string {
	return "min=x"
}
