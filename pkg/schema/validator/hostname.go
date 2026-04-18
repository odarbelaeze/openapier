package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(HostnameTag{})
}

type HostnameTag struct{}

func (t HostnameTag) Tag() string {
	return "hostname"
}

func (t HostnameTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("hostname is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("hostname"),
	}, nil
}

func (t HostnameTag) Usage() string {
	return "hostname"
}
