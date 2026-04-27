package validator

import (
	"fmt"
	"strconv"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(GTETag{})
}

type GTETag struct{}

func (t GTETag) Tag() string {
	return "gte"
}

func (t GTETag) Parse(value string, as string) ([]options.SchemaOption, error) {
	m, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid gte value: %s", value)
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
		return nil, fmt.Errorf("gte is not supported for %s", as)
	}
}

func (t GTETag) Usage() string {
	return "gte=x"
}
