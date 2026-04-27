package spec

import (
	"fmt"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&SecuritySchemeSchemeComment{})
}

var _ Comment = &SecuritySchemeSchemeComment{}

// SecuritySchemeSchemeComment is a comment that updates the 'scheme' property of a security scheme.
type SecuritySchemeSchemeComment struct{}

// Tag implements Comment.
func (c *SecuritySchemeSchemeComment) Tag() string {
	return "securityScheme.scheme"
}

// Usage implements Comment.
func (c *SecuritySchemeSchemeComment) Usage() string {
	return "@securityScheme.scheme <securitySchemeName> <scheme>"
}

// ParseInto implements Comment.
func (c *SecuritySchemeSchemeComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return fmt.Errorf("invalid @securityScheme.scheme format, expected: %s", c.Usage())
	}

	schemeName := fields[0]
	schemeValue := fields[1]

	if s.Spec.Components == nil || s.Spec.Components.Spec.SecuritySchemes == nil {
		return fmt.Errorf("cannot add security scheme HTTP scheme without a preceding @securityScheme %s", schemeName)
	}

	scheme, ok := s.Spec.Components.Spec.SecuritySchemes[schemeName]
	if !ok {
		return fmt.Errorf("security scheme %s not found", schemeName)
	}

	scheme.Spec.Spec.Scheme = schemeValue

	return nil
}
