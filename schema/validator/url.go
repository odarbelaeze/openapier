package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(URLTag{})
}

type URLTag struct{}

func (t URLTag) Tag() string {
	return "url"
}

func (t URLTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("url is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("url"),
	}, nil
}

func (t URLTag) Usage() string {
	return "url"
}
