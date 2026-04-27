package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(IP4AddrTag{})
}

type IP4AddrTag struct{}

func (t IP4AddrTag) Tag() string {
	return "ip4_addr"
}

func (t IP4AddrTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ip4_addr is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("ipv4"),
	}, nil
}

func (t IP4AddrTag) Usage() string {
	return "ip4_addr"
}
