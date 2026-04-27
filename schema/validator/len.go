package validator

import (
	"fmt"
	"strconv"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(LenTag{})
}

type LenTag struct{}

func (t LenTag) Tag() string {
	return "len"
}

func (t LenTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	l, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("invalid len value: %s", value)
	}
	switch as {
	case "object":
		return []options.SchemaOption{
			options.WithMaxProperties(l),
			options.WithMinProperties(l),
		}, nil
	case "array":
		return []options.SchemaOption{
			options.WithMaxItems(l),
			options.WithMinItems(l),
		}, nil
	case "string":
		return []options.SchemaOption{
			options.WithMaxLength(l),
			options.WithMinLength(l),
		}, nil
	case "integer", "number":
		return nil, fmt.Errorf("len is not supported for %s", as)
	}
	return nil, fmt.Errorf("invalid schema type: %s", as)
}

func (t LenTag) Usage() string {
	return "len=x"
}
