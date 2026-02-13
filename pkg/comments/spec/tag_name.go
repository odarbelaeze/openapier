package spec

import "github.com/sv-tools/openapi"

func init() {
	Register(NewTagNameComment())
}

type tagNameComment struct{}

// NewTagNameComment creates a new tagNameComment.
func NewTagNameComment() *tagNameComment {
	return &tagNameComment{}
}

// ParseInto implements Comment.
func (t *tagNameComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	tag := &openapi.Extendable[openapi.Tag]{
		Spec: &openapi.Tag{
			Name: c,
		},
	}
	s.Spec.Tags = append(s.Spec.Tags, tag)
	return nil
}

// Tag implements Comment.
func (t *tagNameComment) Tag() string {
	return "tag.name"
}

// Usage implements Comment.
func (t *tagNameComment) Usage() string {
	return `// @tag.name <name>`
}
