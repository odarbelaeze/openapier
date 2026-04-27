package operation

import "go/ast"

func init() {
	Register(NewDeprecatedComment())
}

// DeprecatedComment marks an operation as deprecated.
type DeprecatedComment struct{}

func NewDeprecatedComment() *DeprecatedComment {
	return &DeprecatedComment{}
}

func (c *DeprecatedComment) Tag() string {
	return "deprecated"
}

func (c *DeprecatedComment) Usage() string {
	return "@deprecated"
}

func (c *DeprecatedComment) ParseInto(content string, f *ast.File, op *Operation) error {
	op.Builder.Deprecated(true)
	return nil
}
