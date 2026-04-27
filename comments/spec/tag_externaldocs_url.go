package spec

import (
	"errors"

	"github.com/sv-tools/openapi"
)

var _ Comment = &TagExternalDocsURLComment{}

func init() {
	Register(&TagExternalDocsURLComment{})
}

type TagExternalDocsURLComment struct{}

// Tag implements [Comment].
func (c *TagExternalDocsURLComment) Tag() string {
	return "tag.externalDocs.url"
}

// Usage implements [Comment].
func (c *TagExternalDocsURLComment) Usage() string {
	return "@tag.externalDocs.url <url>"
}

// ParseInto implements [Comment].
func (c *TagExternalDocsURLComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if len(s.Spec.Tags) == 0 {
		return errors.New("a @tag.externalDocs.url comment requires a preceding @tag.name comment")
	}

	lastTag := s.Spec.Tags[len(s.Spec.Tags)-1]

	if lastTag.Spec.ExternalDocs == nil {
		lastTag.Spec.ExternalDocs = openapi.NewExternalDocsBuilder().Build()
	}

	lastTag.Spec.ExternalDocs.Spec.URL = line

	return nil
}
