package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(URITag{})
}

type URITag struct{}

func (t URITag) Tag() string {
	return "uri"
}

func (t URITag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("uri is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("uri"),
	}, nil
}

func (t URITag) Usage() string {
	return "uri"
}
