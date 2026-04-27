package spec

import (
	"errors"

	"github.com/sv-tools/openapi"
)

var _ Comment = &TagDescriptionComment{}

func init() {
	Register(&TagDescriptionComment{})
}

type TagDescriptionComment struct{}

// Tag implements [Comment].
func (c *TagDescriptionComment) Tag() string {
	return "tag.description"
}

// Usage implements [Comment].
func (c *TagDescriptionComment) Usage() string {
	return "@tag.description <description>"
}

// ParseInto implements [Comment].
func (c *TagDescriptionComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if len(s.Spec.Tags) == 0 {
		return errors.New("a @tag.description comment requires a preceding @tag.name comment")
	}

	lastTag := s.Spec.Tags[len(s.Spec.Tags)-1]
	lastTag.Spec.Description = line

	return nil
}
