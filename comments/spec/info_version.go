package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoVersionComment{}

func init() {
	Register(&InfoVersionComment{})
}

type InfoVersionComment struct{}

// Tag implements [Comment].
func (c *InfoVersionComment) Tag() string {
	return "info.version"
}

// Usage implements [Comment].
func (c *InfoVersionComment) Usage() string {
	return "@info.version <version>"
}

// ParseInto implements [Comment].
func (c *InfoVersionComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.Version = line
	return nil
}
