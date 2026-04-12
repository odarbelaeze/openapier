package validator

import (
	"fmt"
	"strconv"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(LTETag{})
}

type LTETag struct{}

func (t LTETag) Tag() string {
	return "lte"
}

func (t LTETag) Parse(value string, as string) ([]options.SchemaOption, error) {
	m, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid lte value: %s", value)
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
		return nil, fmt.Errorf("lte is not supported for %s", as)
	}
}

func (t LTETag) Usage() string {
	return "lte=x"
}
