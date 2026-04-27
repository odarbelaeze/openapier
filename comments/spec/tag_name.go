package spec

import "github.com/sv-tools/openapi"

var _ Comment = &TagNameComment{}

func init() {
	Register(&TagNameComment{})
}

type TagNameComment struct{}

// Tag implements [Comment].
func (c *TagNameComment) Tag() string {
	return "tag.name"
}

// Usage implements [Comment].
func (c *TagNameComment) Usage() string {
	return "@tag.name <name>"
}

// ParseInto implements [Comment].
func (c *TagNameComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	tag := openapi.NewTagBuilder().Name(line).Build()
	s.Spec.Tags = append(s.Spec.Tags, tag)
	return nil
}
