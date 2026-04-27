package spec

import "github.com/sv-tools/openapi"

var _ Comment = &ExternalDocsDescriptionComment{}

func init() {
	Register(&ExternalDocsDescriptionComment{})
}

type ExternalDocsDescriptionComment struct{}

// Tag implements [Comment].
func (c *ExternalDocsDescriptionComment) Tag() string {
	return "externalDocs.description"
}

// Usage implements [Comment].
func (c *ExternalDocsDescriptionComment) Usage() string {
	return "@externalDocs.description <description>"
}

// ParseInto implements [Comment].
func (c *ExternalDocsDescriptionComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.ExternalDocs == nil {
		s.Spec.ExternalDocs = openapi.NewExternalDocsBuilder().Build()
	}
	s.Spec.ExternalDocs.Spec.Description = line
	return nil
}
