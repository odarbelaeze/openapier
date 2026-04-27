package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(SemverTag{})
}

type SemverTag struct{}

func (t SemverTag) Tag() string {
	return "semver"
}

func (t SemverTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("semver is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("semver"),
	}, nil
}

func (t SemverTag) Usage() string {
	return "semver"
}
