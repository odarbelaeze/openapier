package spec

import "github.com/sv-tools/openapi"

var _ Comment = &ExternalDocsURLComment{}

func init() {
	Register(&ExternalDocsURLComment{})
}

type ExternalDocsURLComment struct{}

// Tag implements [Comment].
func (c *ExternalDocsURLComment) Tag() string {
	return "externalDocs.url"
}

// Usage implements [Comment].
func (c *ExternalDocsURLComment) Usage() string {
	return "@externalDocs.url <url>"
}

// ParseInto implements [Comment].
func (c *ExternalDocsURLComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.ExternalDocs == nil {
		s.Spec.ExternalDocs = openapi.NewExternalDocsBuilder().Build()
	}
	s.Spec.ExternalDocs.Spec.URL = line
	return nil
}
