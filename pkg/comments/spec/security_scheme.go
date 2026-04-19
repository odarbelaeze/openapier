package spec

import (
	"fmt"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&SecuritySchemeComment{})
}

var _ Comment = &SecuritySchemeComment{}

// SecuritySchemeComment is a comment that adds a security scheme to the specification.
type SecuritySchemeComment struct{}

// Tag implements Comment.
func (c *SecuritySchemeComment) Tag() string {
	return "securityScheme"
}

// Usage implements Comment.
func (c *SecuritySchemeComment) Usage() string {
	return "@securityScheme <name> <type> [<description>...]"
}

// ParseInto implements Comment.
func (c *SecuritySchemeComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return fmt.Errorf("invalid @securityScheme format, expected: %s", c.Usage())
	}

	name := fields[0]
	typ := fields[1]
	var description string
	if len(fields) > 2 {
		description = strings.Join(fields[2:], " ")
	}

	if s.Spec.Components == nil {
		s.Spec.Components = openapi.NewComponents()
	}

	if s.Spec.Components.Spec.SecuritySchemes == nil {
		s.Spec.Components.Spec.SecuritySchemes = make(map[string]*openapi.RefOrSpec[openapi.Extendable[openapi.SecurityScheme]])
	}

	builder := openapi.NewSecuritySchemeBuilder().
		Type(typ).
		Description(description)

	s.Spec.Components.Spec.SecuritySchemes[name] = builder.Build()

	return nil
}
