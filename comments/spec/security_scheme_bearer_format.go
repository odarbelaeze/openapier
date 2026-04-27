package spec

import (
	"fmt"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&SecuritySchemeBearerFormatComment{})
}

var _ Comment = &SecuritySchemeBearerFormatComment{}

// SecuritySchemeBearerFormatComment is a comment that updates the 'bearerFormat' property of a security scheme.
type SecuritySchemeBearerFormatComment struct{}

// Tag implements Comment.
func (c *SecuritySchemeBearerFormatComment) Tag() string {
	return "securityScheme.bearerFormat"
}

// Usage implements Comment.
func (c *SecuritySchemeBearerFormatComment) Usage() string {
	return "@securityScheme.bearerFormat <securitySchemeName> <format>"
}

// ParseInto implements Comment.
func (c *SecuritySchemeBearerFormatComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return fmt.Errorf("invalid @securityScheme.bearerFormat format, expected: %s", c.Usage())
	}

	schemeName := fields[0]
	format := fields[1]

	if s.Spec.Components == nil || s.Spec.Components.Spec.SecuritySchemes == nil {
		return fmt.Errorf("cannot add security scheme bearer format without a preceding @securityScheme %s", schemeName)
	}

	scheme, ok := s.Spec.Components.Spec.SecuritySchemes[schemeName]
	if !ok {
		return fmt.Errorf("security scheme %s not found", schemeName)
	}

	scheme.Spec.Spec.BearerFormat = format

	return nil
}
