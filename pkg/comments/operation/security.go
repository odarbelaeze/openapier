package operation

import (
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewSecurityComment())
}

// SecurityComment is a comment that updates the security of an operation.
type SecurityComment struct{}

func NewSecurityComment() *SecurityComment {
	return &SecurityComment{}
}

func (d *SecurityComment) Tag() string {
	return "security"
}

func (d *SecurityComment) Usage() string {
	return "@security <name> [scope1] [scope2] ..."
}

func (d *SecurityComment) ParseInto(content string, op *Operation) error {
	parts := strings.Fields(content)
	if len(parts) == 0 {
		return nil
	}

	name := parts[0]
	var scopes []string
	if len(parts) > 1 {
		scopes = parts[1:]
	}

	req := *openapi.NewSecurityRequirementBuilder().Add(name, scopes...).Build()

	op.Builder.AddSecurity(req)

	return nil
}
