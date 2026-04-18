package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(JSONTag{})
}

type JSONTag struct{}

func (t JSONTag) Tag() string {
	return "json"
}

func (t JSONTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("json is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("json"),
	}, nil
}

func (t JSONTag) Usage() string {
	return "json"
}
