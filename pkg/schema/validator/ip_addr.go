package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(IPAddrTag{})
}

type IPAddrTag struct{}

func (t IPAddrTag) Tag() string {
	return "ip_addr"
}

func (t IPAddrTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ip_addr is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("ip"),
	}, nil
}

func (t IPAddrTag) Usage() string {
	return "ip_addr"
}
