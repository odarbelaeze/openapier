package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(FQDNTag{})
}

type FQDNTag struct{}

func (t FQDNTag) Tag() string {
	return "fqdn"
}

func (t FQDNTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("fqdn is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("hostname"),
	}, nil
}

func (t FQDNTag) Usage() string {
	return "fqdn"
}
