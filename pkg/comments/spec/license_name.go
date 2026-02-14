package spec

import (
	"github.com/sv-tools/openapi"
)

var _ Comment = &licenseNameComment{}

func init() {
	Register(NewLicenseNameComment())
}

type licenseNameComment struct{}

// NewLicenseNameComment creates a new licenseNameComment.
func NewLicenseNameComment() *licenseNameComment {
	return &licenseNameComment{}
}

// ParseInto implements Comment.
func (*licenseNameComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.License == nil {
		s.Spec.Info.Spec.License = openapi.NewLicenseBuilder().Build()
	}
	s.Spec.Info.Spec.License.Spec.Name = c
	return nil
}

// Tag implements Comment.
func (l *licenseNameComment) Tag() string {
	return "license.name"
}

// Usage implements Comment.
func (l *licenseNameComment) Usage() string {
	return `// @license.name <name>`
}
