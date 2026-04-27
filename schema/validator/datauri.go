package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(DataURITag{})
}

type DataURITag struct{}

func (t DataURITag) Tag() string {
	return "datauri"
}

func (t DataURITag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("datauri is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern("^data:.*"),
	}, nil
}

func (t DataURITag) Usage() string {
	return "datauri"
}
