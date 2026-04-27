package spec

import (
	"fmt"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&SecuritySchemeInComment{})
}

var _ Comment = &SecuritySchemeInComment{}

// SecuritySchemeInComment is a comment that updates the 'in' property of a security scheme.
type SecuritySchemeInComment struct{}

// Tag implements Comment.
func (c *SecuritySchemeInComment) Tag() string {
	return "securityScheme.in"
}

// Usage implements Comment.
func (c *SecuritySchemeInComment) Usage() string {
	return "@securityScheme.in <securitySchemeName> <in>"
}

// ParseInto implements Comment.
func (c *SecuritySchemeInComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return fmt.Errorf("invalid @securityScheme.in format, expected: %s", c.Usage())
	}

	schemeName := fields[0]
	in := fields[1]

	if s.Spec.Components == nil || s.Spec.Components.Spec.SecuritySchemes == nil {
		return fmt.Errorf("cannot add security scheme location without a preceding @securityScheme %s", schemeName)
	}

	scheme, ok := s.Spec.Components.Spec.SecuritySchemes[schemeName]
	if !ok {
		return fmt.Errorf("security scheme %s not found", schemeName)
	}

	scheme.Spec.Spec.In = in

	return nil
}
