package validator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(OneOfTag{})
}

type OneOfTag struct{}

func (t OneOfTag) Tag() string {
	return "oneof"
}

func (t OneOfTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	parts := strings.Fields(value)
	if len(parts) == 0 {
		return nil, fmt.Errorf("oneof requires at least one value")
	}
	values := make([]any, 0, len(parts))
	for _, part := range parts {
		var val any
		var err error

		switch as {
		case "integer":
			val, err = strconv.Atoi(part)
		case "number":
			val, err = strconv.ParseFloat(part, 64)
		case "boolean":
			val, err = strconv.ParseBool(part)
		case "string":
			val = part
		default:
			return nil, fmt.Errorf("oneof is not supported for %s", as)
		}

		if err != nil {
			return nil, fmt.Errorf("invalid oneof value for %s: %s", as, part)
		}

		values = append(values, val)
	}

	return []options.SchemaOption{
		options.WithEnum(values...),
	}, nil
}

func (t OneOfTag) Usage() string {
	return "oneof=val1 val2 val3"
}
