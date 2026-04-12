package validator

import (
	"fmt"
	"strconv"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(LTTag{})
}

type LTTag struct{}

func (t LTTag) Tag() string {
	return "lt"
}

func (t LTTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	m, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid lt value: %s", value)
	}

	switch as {
	case "integer", "number":
		return []options.SchemaOption{
			options.WithExclusiveMaximum(int(m)),
		}, nil
	case "string":
		return []options.SchemaOption{
			options.WithMaxLength(int(m) - 1),
		}, nil
	case "array":
		return []options.SchemaOption{
			options.WithMaxItems(int(m) - 1),
		}, nil
	case "object":
		return []options.SchemaOption{
			options.WithMaxProperties(int(m) - 1),
		}, nil
	default:
		return nil, fmt.Errorf("lt is not supported for %s", as)
	}
}

func (t LTTag) Usage() string {
	return "lt=x"
}
