package validator

import (
	"fmt"
	"strconv"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(EqTag{})
}

type EqTag struct{}

func (t EqTag) Tag() string {
	return "eq"
}

func (t EqTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	var val any
	var err error

	switch as {
	case "integer":
		val, err = strconv.Atoi(value)
	case "number":
		val, err = strconv.ParseFloat(value, 64)
	case "boolean":
		val, err = strconv.ParseBool(value)
	case "string":
		val = value
	default:
		return nil, fmt.Errorf("eq is not supported for %s", as)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid eq value for %s: %s", as, value)
	}

	return []options.SchemaOption{
		options.WithEnum(val),
	}, nil
}

func (t EqTag) Usage() string {
	return "eq=x"
}
