package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(UUIDTag{})
}

type UUIDTag struct{}

func (t UUIDTag) Tag() string {
	return "uuid"
}

func (t UUIDTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("uuid is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("uuid"),
	}, nil
}

func (t UUIDTag) Usage() string {
	return "uuid"
}
