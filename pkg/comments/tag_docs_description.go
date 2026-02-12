package comments

import (
	"fmt"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewTagDocsDescriptionComment())
}

type tagDocsDescriptionComment struct{}

// NewTagDocsDescriptionComment creates a new tagDocsDescriptionComment.
func NewTagDocsDescriptionComment() *tagDocsDescriptionComment {
	return &tagDocsDescriptionComment{}
}

// ParseInto implements Comment.
func (t *tagDocsDescriptionComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if len(s.Spec.Tags) == 0 {
		return fmt.Errorf("@tag.docs.description must follow @tag.name")
	}
	tag := s.Spec.Tags[len(s.Spec.Tags)-1]

	if tag.Spec.ExternalDocs == nil {
		return fmt.Errorf("@tag.docs.description needs to come after a @tag.docs.url")
	}

	tag.Spec.ExternalDocs.Spec.Description = c
	return nil
}

// Tag implements Comment.
func (t *tagDocsDescriptionComment) Tag() string {
	return "tag.docs.description"
}

// Usage implements Comment.
func (t *tagDocsDescriptionComment) Usage() string {
	return `// @tag.docs.description <description>`
}
