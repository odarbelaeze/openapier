package validator

import "github.com/odarbelaeze/openapier/schema/options"

func init() {
	Default().Register(RequiredTag{})
}

type RequiredTag struct{}

func (t RequiredTag) Tag() string {
	return "required"
}

func (t RequiredTag) Parse(string, string) ([]options.SchemaOption, error) {
	return []options.SchemaOption{
		options.WithRequired(),
	}, nil
}

func (t RequiredTag) Usage() string {
	return "required"
}
