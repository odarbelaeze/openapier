package validator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/odarbelaeze/openapier/schema/options"
	"github.com/sv-tools/openapi"
)

func init() {
	Default().Register(NoneOfTag{})
}

type NoneOfTag struct{}

func (t NoneOfTag) Tag() string {
	return "noneof"
}

func (t NoneOfTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	parts := strings.Split(value, " ")
	if len(parts) == 0 {
		return nil, errors.New("noneof requires at least one value")
	}
	values := make([]any, len(parts))
	for i, part := range parts {
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
			return nil, fmt.Errorf("noneof is not supported for %s", as)
		}

		if err != nil {
			return nil, fmt.Errorf("invalid noneof value for %s: %s", as, part)
		}

		values[i] = val
	}
	return []options.SchemaOption{
		options.WithNot(openapi.NewSchemaBuilder().Enum(values...).Build()),
	}, nil
}

func (t NoneOfTag) Usage() string {
	return "noneof=value1 value2"
}
