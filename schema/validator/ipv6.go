package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(IPv6Tag{})
}

type IPv6Tag struct{}

func (t IPv6Tag) Tag() string {
	return "ipv6"
}

func (t IPv6Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ipv6 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("ipv6"),
	}, nil
}

func (t IPv6Tag) Usage() string {
	return "ipv6"
}
