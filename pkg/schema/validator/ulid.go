package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(ULIDTag{})
}

type ULIDTag struct{}

func (t ULIDTag) Tag() string {
	return "ulid"
}

func (t ULIDTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ulid is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("ulid"),
	}, nil
}

func (t ULIDTag) Usage() string {
	return "ulid"
}
