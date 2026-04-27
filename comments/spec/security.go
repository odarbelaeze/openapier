package spec

import (
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&SecurityComment{})
}

var _ Comment = &SecurityComment{}

// SecurityComment is a comment that updates the security of the global specification.
type SecurityComment struct{}

// Tag implements Comment.
func (c *SecurityComment) Tag() string {
	return "security"
}

// Usage implements Comment.
func (c *SecurityComment) Usage() string {
	return "@security <name> [scope1] [scope2] ..."
}

// ParseInto implements Comment.
func (c *SecurityComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil
	}

	name := parts[0]
	var scopes []string
	if len(parts) > 1 {
		scopes = parts[1:]
	}

	req := *openapi.NewSecurityRequirementBuilder().Add(name, scopes...).Build()

	if s.Spec.Security == nil {
		s.Spec.Security = make([]openapi.SecurityRequirement, 0)
	}
	s.Spec.Security = append(s.Spec.Security, req)

	return nil
}
