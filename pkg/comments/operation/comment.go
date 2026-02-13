package operation

import "github.com/sv-tools/openapi"

// Comment defines the interface for operation-level OpenAPI annotations.
type Comment interface {
	Tag() string
	Usage() string
	ParseInto(c string, target *openapi.Extendable[openapi.Operation]) error
}
