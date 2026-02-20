package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoLicenseNameComment{}

func init() {
	Register(&InfoLicenseNameComment{})
}

type InfoLicenseNameComment struct{}

// Tag implements [Comment].
func (c *InfoLicenseNameComment) Tag() string {
	return "info.license.name"
}

// Usage implements [Comment].
func (c *InfoLicenseNameComment) Usage() string {
	return "@info.license.name <name>"
}

// ParseInto implements [Comment].
func (c *InfoLicenseNameComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.License == nil {
		s.Spec.Info.Spec.License = openapi.NewLicenseBuilder().Build()
	}
	s.Spec.Info.Spec.License.Spec.Name = line
	return nil
}
