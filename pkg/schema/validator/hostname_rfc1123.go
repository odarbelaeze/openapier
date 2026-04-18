package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(HostnameRFC1123Tag{})
}

type HostnameRFC1123Tag struct{}

func (t HostnameRFC1123Tag) Tag() string {
	return "hostname_rfc1123"
}

func (t HostnameRFC1123Tag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("hostname_rfc1123 is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("hostname_rfc1123"),
	}, nil
}

func (t HostnameRFC1123Tag) Usage() string {
	return "hostname_rfc1123"
}
