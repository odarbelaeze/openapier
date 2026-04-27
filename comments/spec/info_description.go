package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoDescriptionComment{}

func init() {
	Register(&InfoDescriptionComment{})
}

type InfoDescriptionComment struct{}

// Tag implements [Comment].
func (c *InfoDescriptionComment) Tag() string {
	return "info.description"
}

// Usage implements [Comment].
func (c *InfoDescriptionComment) Usage() string {
	return "@info.description <description>"
}

// ParseInto implements [Comment].
func (c *InfoDescriptionComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.Description = line
	return nil
}
