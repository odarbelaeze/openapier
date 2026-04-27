package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(MACTag{})
}

type MACTag struct{}

func (t MACTag) Tag() string {
	return "mac"
}

func (t MACTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("mac is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("mac"),
	}, nil
}

func (t MACTag) Usage() string {
	return "mac"
}
