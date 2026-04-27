package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(HTTPURLTag{})
}

type HTTPURLTag struct{}

func (t HTTPURLTag) Tag() string {
	return "http_url"
}

func (t HTTPURLTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("http_url is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^https?://.*"),
	}, nil
}

func (t HTTPURLTag) Usage() string {
	return "http_url"
}
