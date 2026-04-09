package validator

import "github.com/odarbelaeze/openapier/pkg/schema/options"

type ValidatorTag interface {
	Tag() string
	Parse(string, string) ([]options.SchemaOption, error)
	Usage() string
}
