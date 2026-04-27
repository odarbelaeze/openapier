package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(HTTPSURLTag{})
}

type HTTPSURLTag struct{}

func (t HTTPSURLTag) Tag() string {
	return "https_url"
}

func (t HTTPSURLTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("https_url is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^https://.*"),
	}, nil
}

func (t HTTPSURLTag) Usage() string {
	return "https_url"
}
