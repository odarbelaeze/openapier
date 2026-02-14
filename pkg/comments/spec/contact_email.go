package spec

import (
	"github.com/sv-tools/openapi"
)

var _ Comment = &contactEmailComment{}

func init() {
	Register(NewContactEmailComment())
}

type contactEmailComment struct{}

// NewContactEmailComment creates a new contactEmailComment.
func NewContactEmailComment() *contactEmailComment {
	return &contactEmailComment{}
}

// ParseInto implements Comment.
func (*contactEmailComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.Contact == nil {
		s.Spec.Info.Spec.Contact = openapi.NewContactBuilder().Build()
	}
	s.Spec.Info.Spec.Contact.Spec.Email = c
	return nil
}

// Tag implements Comment.
func (c *contactEmailComment) Tag() string {
	return "contact.email"
}

// Usage implements Comment.
func (c *contactEmailComment) Usage() string {
	return `// @contact.email <email>`
}
