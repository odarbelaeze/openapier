package validator

import "github.com/odarbelaeze/openapier/schema/options"

type Tag interface {
	Tag() string
	Usage() string
	Parse(value string, as string) ([]options.SchemaOption, error)
}
