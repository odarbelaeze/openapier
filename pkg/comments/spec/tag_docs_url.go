package spec

import (
	"fmt"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewTagDocsURLComment())
}

type tagDocsURLComment struct{}

// NewTagDocsURLComment creates a new tagDocsURLComment.
func NewTagDocsURLComment() *tagDocsURLComment {
	return &tagDocsURLComment{}
}

// ParseInto implements Comment.
func (t *tagDocsURLComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if len(s.Spec.Tags) == 0 {
		return fmt.Errorf("@tag.docs.url must follow @tag.name")
	}
	tag := s.Spec.Tags[len(s.Spec.Tags)-1]

	tag.Spec.ExternalDocs = &openapi.Extendable[openapi.ExternalDocs]{
		Spec: &openapi.ExternalDocs{
			URL: c,
		},
	}
	return nil
}

// Tag implements Comment.
func (t *tagDocsURLComment) Tag() string {
	return "tag.docs.url"
}

// Usage implements Comment.
func (t *tagDocsURLComment) Usage() string {
	return `// @tag.docs.url <url>`
}
