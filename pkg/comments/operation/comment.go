package operation

import "go/ast"

// Comment defines the interface for operation-level OpenAPI annotations.
type Comment interface {
	Tag() string
	Usage() string
	ParseInto(c string, f *ast.File, op *Operation) error
}
