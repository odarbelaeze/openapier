package spec

import "github.com/sv-tools/openapi"

type Comment interface {
	Tag() string
	Usage() string
	ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error
}
