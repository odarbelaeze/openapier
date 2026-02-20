package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoContactEmailComment{}

func init() {
	Register(&InfoContactEmailComment{})
}

type InfoContactEmailComment struct{}

// Tag implements [Comment].
func (c *InfoContactEmailComment) Tag() string {
	return "info.contact.email"
}

// Usage implements [Comment].
func (c *InfoContactEmailComment) Usage() string {
	return "@info.contact.email <email>"
}

// ParseInto implements [Comment].
func (c *InfoContactEmailComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.Contact == nil {
		s.Spec.Info.Spec.Contact = openapi.NewContactBuilder().Build()
	}
	s.Spec.Info.Spec.Contact.Spec.Email = line
	return nil
}
