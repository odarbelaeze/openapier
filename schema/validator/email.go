package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(EmailTag{})
}

type EmailTag struct{}

func (t EmailTag) Tag() string {
	return "email"
}

func (t EmailTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("email is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("email"),
	}, nil
}

func (t EmailTag) Usage() string {
	return "email"
}
