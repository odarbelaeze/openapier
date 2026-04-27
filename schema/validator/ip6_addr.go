package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(IP6AddrTag{})
}

type IP6AddrTag struct{}

func (t IP6AddrTag) Tag() string {
	return "ip6_addr"
}

func (t IP6AddrTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ip6_addr is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("ipv6"),
	}, nil
}

func (t IP6AddrTag) Usage() string {
	return "ip6_addr"
}
