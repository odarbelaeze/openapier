package spec

import (
	"github.com/sv-tools/openapi"
)

var _ Comment = &titleComment{}

func init() {
	Register(NewTitleComment())
}

type titleComment struct{}

// NewTitleComment creates a new titleComment.
func NewTitleComment() *titleComment {
	return &titleComment{}
}

// ParseInto implements Comment.
func (*titleComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.Title = c
	return nil
}

// Tag implements Comment.
func (t *titleComment) Tag() string {
	return "title"
}

// Usage implements Comment.
func (t *titleComment) Usage() string {
	return `// @title <title>`
}
