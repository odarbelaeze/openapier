package comments

import (
	"errors"

	"github.com/sv-tools/openapi"
)

var _ Comment = &hostComment{}

func init() {
	Register(NewHostComment())
}

type hostComment struct{}

// NewHostComment creates a new hostComment.
func NewHostComment() *hostComment {
	return &hostComment{}
}

// ParseInto implements Comment.
func (h *hostComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	return errors.New("@host is not supported use @servers.url instead")
}

// Tag implements Comment.
func (h *hostComment) Tag() string {
	return "host"
}

// Usage implements Comment.
func (h *hostComment) Usage() string {
	return `// @host <host>`
}
