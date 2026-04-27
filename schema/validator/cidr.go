package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(CIDRTag{})
}

type CIDRTag struct{}

func (t CIDRTag) Tag() string {
	return "cidr"
}

func (t CIDRTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("cidr is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("cidr"),
	}, nil
}

func (t CIDRTag) Usage() string {
	return "cidr"
}
