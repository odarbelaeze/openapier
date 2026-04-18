package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(IPv4Tag{})
}

type IPv4Tag struct{}

func (t IPv4Tag) Tag() string {
	return "ipv4"
}

func (t IPv4Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ipv4 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("ipv4"),
	}, nil
}

func (t IPv4Tag) Usage() string {
	return "ipv4"
}
