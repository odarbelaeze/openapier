package validator

import (
	"github.com/odarbelaeze/openapier/pkg/schema/options"
	"github.com/sv-tools/openapi"
)

func init() {
	Default().Register(NETag{})
}

type NETag struct{}

func (t NETag) Tag() string {
	return "ne"
}

func (t NETag) Parse(value string, as string) ([]options.SchemaOption, error) {
	return []options.SchemaOption{
		options.WithNot(openapi.NewSchemaBuilder().Enum(value).Build()),
	}, nil
}

func (t NETag) Usage() string {
	return "ne=value"
}
