package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoContactURLComment{}

func init() {
	Register(&InfoContactURLComment{})
}

type InfoContactURLComment struct{}

// Tag implements [Comment].
func (c *InfoContactURLComment) Tag() string {
	return "info.contact.url"
}

// Usage implements [Comment].
func (c *InfoContactURLComment) Usage() string {
	return "@info.contact.url <url>"
}

// ParseInto implements [Comment].
func (c *InfoContactURLComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.Contact == nil {
		s.Spec.Info.Spec.Contact = openapi.NewContactBuilder().Build()
	}
	s.Spec.Info.Spec.Contact.Spec.URL = line
	return nil
}
