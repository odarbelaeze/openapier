package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(DateTimeTag{})
}

type DateTimeTag struct{}

func (t DateTimeTag) Tag() string {
	return "datetime"
}

func (t DateTimeTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("datetime is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("date-time"),
	}, nil
}

func (t DateTimeTag) Usage() string {
	return "datetime"
}
