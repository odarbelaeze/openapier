package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(HostnamePortTag{})
}

type HostnamePortTag struct{}

func (t HostnamePortTag) Tag() string {
	return "hostname_port"
}

func (t HostnamePortTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("hostname_port is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^.*:[0-9]+$"),
	}, nil
}

func (t HostnamePortTag) Usage() string {
	return "hostname_port"
}
