package validator

import (
	"github.com/odarbelaeze/openapier/pkg/schema/options"
	"github.com/sv-tools/openapi"
	"strings"
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
	values := make([]any, len(parts))
	for i, part := range parts {
		values[i] = part
	}
	return []options.SchemaOption{
		options.WithNot(openapi.NewSchemaBuilder().Enum(values...).Build()),
	}, nil
}

func (t NoneOfTag) Usage() string {
	return "noneof=value1 value2"
}
