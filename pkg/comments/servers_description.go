package comments

import (
	"errors"

	"github.com/sv-tools/openapi"
)

var _ Comment = &serversDescriptionComment{}

type serversDescriptionComment struct{}

// NewServersDescriptionComment creates a new serversDescriptionComment.
func NewServersDescriptionComment() *serversDescriptionComment {
	return &serversDescriptionComment{}
}

// ParseInto implements Comment.
func (*serversDescriptionComment) ParseInto(c string, s openapi.OpenAPI) error {
	if len(s.Servers) == 0 {
		return errors.New("use @servers.url before you use @servers.description")
	}
	server := s.Servers[len(s.Servers)-1]
	server.Spec.Description = c
	return nil
}

// Tag implements Comment.
func (s *serversDescriptionComment) Tag() string {
	return "servers.description"
}

// Usage implements Comment.
func (s *serversDescriptionComment) Usage() string {
	return `// @servers.description <description>`
}
