package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoLicenseIdentifierComment{}

func init() {
	Register(&InfoLicenseIdentifierComment{})
}

type InfoLicenseIdentifierComment struct{}

// Tag implements [Comment].
func (c *InfoLicenseIdentifierComment) Tag() string {
	return "info.license.identifier"
}

// Usage implements [Comment].
func (c *InfoLicenseIdentifierComment) Usage() string {
	return "@info.license.identifier <identifier>"
}

// ParseInto implements [Comment].
func (c *InfoLicenseIdentifierComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	if s.Spec.Info.Spec.License == nil {
		s.Spec.Info.Spec.License = openapi.NewLicenseBuilder().Build()
	}
	s.Spec.Info.Spec.License.Spec.Identifier = line
	return nil
}
