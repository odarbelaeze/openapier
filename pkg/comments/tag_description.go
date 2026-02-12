package comments

import (
	"fmt"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewTagDescriptionComment())
}

type tagDescriptionComment struct{}

// NewTagDescriptionComment creates a new tagDescriptionComment.
func NewTagDescriptionComment() *tagDescriptionComment {
	return &tagDescriptionComment{}
}

// ParseInto implements Comment.
func (t *tagDescriptionComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if len(s.Spec.Tags) == 0 {
		return fmt.Errorf("@tag.description must follow @tag.name")
	}
	tag := s.Spec.Tags[len(s.Spec.Tags)-1]
	tag.Spec.Description = c
	return nil
}

// Tag implements Comment.
func (t *tagDescriptionComment) Tag() string {
	return "tag.description"
}

// Usage implements Comment.
func (t *tagDescriptionComment) Usage() string {
	return `// @tag.description <description>`
}
