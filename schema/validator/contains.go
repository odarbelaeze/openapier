package validator

import (
	"fmt"
	"regexp"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(ContainsTag{})
}

type ContainsTag struct{}

func (t ContainsTag) Tag() string {
	return "contains"
}

func (t ContainsTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("contains is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithPattern(regexp.QuoteMeta(value)),
	}, nil
}

func (t ContainsTag) Usage() string {
	return "contains=value"
}
