package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(CIDRv6Tag{})
}

type CIDRv6Tag struct{}

func (t CIDRv6Tag) Tag() string {
	return "cidrv6"
}

func (t CIDRv6Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("cidrv6 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("cidrv6"),
	}, nil
}

func (t CIDRv6Tag) Usage() string {
	return "cidrv6"
}
