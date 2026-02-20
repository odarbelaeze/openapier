package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoLicenseURLComment{}

func init() {
	Register(&InfoLicenseURLComment{})
}

type InfoLicenseURLComment struct{}

// Tag implements [Comment].
func (c *InfoLicenseURLComment) Tag() string {
	return "info.license.url"
}

// Usage implements [Comment].
func (c *InfoLicenseURLComment) Usage() string {
	return "@info.license.url <url>"
}

// ParseInto implements [Comment].
func (c *InfoLicenseURLComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.License == nil {
		s.Spec.Info.Spec.License = openapi.NewLicenseBuilder().Build()
	}
	s.Spec.Info.Spec.License.Spec.URL = line
	return nil
}
