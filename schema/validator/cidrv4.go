package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(CIDRv4Tag{})
}

type CIDRv4Tag struct{}

func (t CIDRv4Tag) Tag() string {
	return "cidrv4"
}

func (t CIDRv4Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("cidrv4 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("cidrv4"),
	}, nil
}

func (t CIDRv4Tag) Usage() string {
	return "cidrv4"
}
