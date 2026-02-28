package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoTermsOfServiceComment{}

func init() {
	Register(&InfoTermsOfServiceComment{})
}

type InfoTermsOfServiceComment struct{}

// Tag implements [Comment].
func (c *InfoTermsOfServiceComment) Tag() string {
	return "info.termsOfService"
}

// Usage implements [Comment].
func (c *InfoTermsOfServiceComment) Usage() string {
	return "@info.termsOfService <url>"
}

// ParseInto implements [Comment].
func (c *InfoTermsOfServiceComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.TermsOfService = line
	return nil
}
