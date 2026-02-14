package spec

import (
	"github.com/sv-tools/openapi"
)

var _ Comment = &termsOfServiceComment{}

func init() {
	Register(NewTermsOfServiceComment())
}

type termsOfServiceComment struct{}

// NewTermsOfServiceComment creates a new termsOfServiceComment.
func NewTermsOfServiceComment() *termsOfServiceComment {
	return &termsOfServiceComment{}
}

// ParseInto implements Comment.
func (*termsOfServiceComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.TermsOfService = c
	return nil
}

// Tag implements Comment.
func (t *termsOfServiceComment) Tag() string {
	return "termsofservice"
}

// Usage implements Comment.
func (t *termsOfServiceComment) Usage() string {
	return `// @termsofservice <url>`
}
