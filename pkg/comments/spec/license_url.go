package spec

import (
	"github.com/sv-tools/openapi"
)

var _ Comment = &licenseURLComment{}

func init() {
	Register(NewLicenseURLComment())
}

type licenseURLComment struct{}

// NewLicenseURLComment creates a new licenseURLComment.
func NewLicenseURLComment() *licenseURLComment {
	return &licenseURLComment{}
}

// ParseInto implements Comment.
func (*licenseURLComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.License == nil {
		s.Spec.Info.Spec.License = openapi.NewLicenseBuilder().Build()
	}
	s.Spec.Info.Spec.License.Spec.URL = c
	return nil
}

// Tag implements Comment.
func (l *licenseURLComment) Tag() string {
	return "license.url"
}

// Usage implements Comment.
func (l *licenseURLComment) Usage() string {
	return `// @license.url <url>`
}
