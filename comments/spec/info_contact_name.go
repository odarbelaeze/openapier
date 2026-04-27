package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoContactNameComment{}

func init() {
	Register(&InfoContactNameComment{})
}

type InfoContactNameComment struct{}

// Tag implements [Comment].
func (c *InfoContactNameComment) Tag() string {
	return "info.contact.name"
}

// Usage implements [Comment].
func (c *InfoContactNameComment) Usage() string {
	return "@info.contact.name <name>"
}

// ParseInto implements [Comment].
func (c *InfoContactNameComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.Contact == nil {
		s.Spec.Info.Spec.Contact = openapi.NewContactBuilder().Build()
	}
	s.Spec.Info.Spec.Contact.Spec.Name = line
	return nil
}
