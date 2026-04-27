package operation

import (
	"go/ast"
)

func init() {
	Register(NewSecurityNoneComment())
}

// SecurityNoneComment is a comment that clears the security of an operation.
type SecurityNoneComment struct{}

func NewSecurityNoneComment() *SecurityNoneComment {
	return &SecurityNoneComment{}
}

func (d *SecurityNoneComment) Tag() string {
	return "security.none"
}

func (d *SecurityNoneComment) Usage() string {
	return "@security.none"
}

func (d *SecurityNoneComment) ParseInto(content string, f *ast.File, op *Operation) error {
	op.Builder.Security()
	return nil
}
