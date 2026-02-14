package spec

import (
	"github.com/sv-tools/openapi"
)

var _ Comment = &descriptionComment{}

func init() {
	Register(NewDescriptionComment())
}

type descriptionComment struct{}

// NewDescriptionComment creates a new descriptionComment.
func NewDescriptionComment() *descriptionComment {
	return &descriptionComment{}
}

// ParseInto implements Comment.
func (*descriptionComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.Description = c
	return nil
}

// Tag implements Comment.
func (d *descriptionComment) Tag() string {
	return "description"
}

// Usage implements Comment.
func (d *descriptionComment) Usage() string {
	return `// @description <description>`
}
