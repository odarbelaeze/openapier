package spec

import "github.com/sv-tools/openapi"

var _ Comment = &InfoSummaryComment{}

func init() {
	Register(&InfoSummaryComment{})
}

type InfoSummaryComment struct{}

// Tag implements [Comment].
func (c *InfoSummaryComment) Tag() string {
	return "info.summary"
}

// Usage implements [Comment].
func (c *InfoSummaryComment) Usage() string {
	return "@info.summary <summary>"
}

// ParseInto implements [Comment].
func (c *InfoSummaryComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.Summary = line
	return nil
}
