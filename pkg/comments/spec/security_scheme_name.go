package spec

import (
	"fmt"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&SecuritySchemeNameComment{})
}

var _ Comment = &SecuritySchemeNameComment{}

// SecuritySchemeNameComment is a comment that updates the name of a security scheme.
type SecuritySchemeNameComment struct{}

// Tag implements Comment.
func (c *SecuritySchemeNameComment) Tag() string {
	return "securityScheme.name"
}

// Usage implements Comment.
func (c *SecuritySchemeNameComment) Usage() string {
	return "@securityScheme.name <securitySchemeName> <name>"
}

// ParseInto implements Comment.
func (c *SecuritySchemeNameComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return fmt.Errorf("invalid @securityScheme.name format, expected: %s", c.Usage())
	}

	schemeName := fields[0]
	name := fields[1]

	if s.Spec.Components == nil || s.Spec.Components.Spec.SecuritySchemes == nil {
		return fmt.Errorf("cannot add security scheme name without a preceding @securityScheme %s", schemeName)
	}

	scheme, ok := s.Spec.Components.Spec.SecuritySchemes[schemeName]
	if !ok {
		return fmt.Errorf("security scheme %s not found", schemeName)
	}

	scheme.Spec.Spec.Name = name

	return nil
}
