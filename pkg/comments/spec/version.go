package spec

import (
	"github.com/sv-tools/openapi"
)

var _ Comment = &versionComment{}

func init() {
	Register(NewVersionComment())
}

type versionComment struct{}

// NewVersionComment creates a new versionComment.
func NewVersionComment() *versionComment {
	return &versionComment{}
}

// ParseInto implements Comment.
func (*versionComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if s.Spec.Info == nil {
		s.Spec.Info = openapi.NewInfoBuilder().Build()
	}
	s.Spec.Info.Spec.Version = c
	return nil
}

// Tag implements Comment.
func (v *versionComment) Tag() string {
	return "version"
}

// Usage implements Comment.
func (v *versionComment) Usage() string {
	return `// @version <version>`
}
