package spec

import (
	"errors"

	"github.com/sv-tools/openapi"
)

var _ Comment = &TagExternalDocsDescriptionComment{}

func init() {
	Register(&TagExternalDocsDescriptionComment{})
}

type TagExternalDocsDescriptionComment struct{}

// Tag implements [Comment].
func (c *TagExternalDocsDescriptionComment) Tag() string {
	return "tag.externalDocs.description"
}

// Usage implements [Comment].
func (c *TagExternalDocsDescriptionComment) Usage() string {
	return "@tag.externalDocs.description <description>"
}

// ParseInto implements [Comment].
func (c *TagExternalDocsDescriptionComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if len(s.Spec.Tags) == 0 {
		return errors.New("a @tag.externalDocs.description comment requires a preceding @tag.name comment")
	}

	lastTag := s.Spec.Tags[len(s.Spec.Tags)-1]

	if lastTag.Spec.ExternalDocs == nil {
		lastTag.Spec.ExternalDocs = openapi.NewExternalDocsBuilder().Build()
	}

	lastTag.Spec.ExternalDocs.Spec.Description = line

	return nil
}
