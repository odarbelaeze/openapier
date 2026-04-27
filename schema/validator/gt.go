package validator

import (
	"fmt"
	"strconv"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(GTTag{})
}

type GTTag struct{}

func (t GTTag) Tag() string {
	return "gt"
}

func (t GTTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	m, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid gt value: %s", value)
	}

	switch as {
	case "integer", "number":
		return []options.SchemaOption{
			options.WithExclusiveMinimum(int(m)),
		}, nil
	case "string":
		return []options.SchemaOption{
			options.WithMinLength(int(m) + 1),
		}, nil
	case "array":
		return []options.SchemaOption{
			options.WithMinItems(int(m) + 1),
		}, nil
	case "object":
		return []options.SchemaOption{
			options.WithMinProperties(int(m) + 1),
		}, nil
	default:
		return nil, fmt.Errorf("gt is not supported for %s", as)
	}
}

func (t GTTag) Usage() string {
	return "gt=x"
}
