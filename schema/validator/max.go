package validator

import (
	"fmt"
	"strconv"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(MaxTag{})
}

type MaxTag struct{}

func (t MaxTag) Tag() string {
	return "max"
}

func (t MaxTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	m, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid max value: %s", value)
	}

	switch as {
	case "integer", "number":
		return []options.SchemaOption{
			options.WithMaximum(int(m)),
		}, nil
	case "string":
		return []options.SchemaOption{
			options.WithMaxLength(int(m)),
		}, nil
	case "array":
		return []options.SchemaOption{
			options.WithMaxItems(int(m)),
		}, nil
	case "object":
		return []options.SchemaOption{
			options.WithMaxProperties(int(m)),
		}, nil
	default:
		return nil, fmt.Errorf("max is not supported for %s", as)
	}
}

func (t MaxTag) Usage() string {
	return "max=x"
}
