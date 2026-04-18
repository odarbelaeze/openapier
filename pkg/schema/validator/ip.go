package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(IPTag{})
}

type IPTag struct{}

func (t IPTag) Tag() string {
	return "ip"
}

func (t IPTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("ip is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("ip"),
	}, nil
}

func (t IPTag) Usage() string {
	return "ip"
}
