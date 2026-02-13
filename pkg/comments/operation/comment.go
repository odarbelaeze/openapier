package operation

// Comment defines the interface for operation-level OpenAPI annotations.
type Comment interface {
	Tag() string
	Usage() string
	ParseInto(c string, op *Operation) error
}
