package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoTitleComment{}

func init() {
	Register(&InfoTitleComment{})
}

type InfoTitleComment struct{}

// InfoTitleComment implements [Comment].
func (c *InfoTitleComment) Tag() string {
	return "info.title"
}

// Usage implements [Comment].
func (c *InfoTitleComment) Usage() string {
	return "@info.title <title>"
}

// ParseInto implements [Comment].
func (c *InfoTitleComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.Title = line
	return nil
}
