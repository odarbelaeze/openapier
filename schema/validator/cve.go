package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(CVETag{})
}

type CVETag struct{}

func (t CVETag) Tag() string {
	return "cve"
}

func (t CVETag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("cve is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("cve"),
	}, nil
}

func (t CVETag) Usage() string {
	return "cve"
}
