package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(UniqueTag{})
}

type UniqueTag struct{}

func (t UniqueTag) Tag() string {
	return "unique"
}

func (t UniqueTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "array" {
		return nil, fmt.Errorf("unique is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithUniqueItems(),
	}, nil
}

func (t UniqueTag) Usage() string {
	return "unique"
}
