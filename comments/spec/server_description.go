package spec

import (
	"errors"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&ServerDescriptionComment{})
}

var _ Comment = &ServerDescriptionComment{}

// ServerDescriptionComment is a comment that updates the description of the last added server.
type ServerDescriptionComment struct{}

// Tag implements Comment.
func (c *ServerDescriptionComment) Tag() string {
	return "server.description"
}

// Usage implements Comment.
func (c *ServerDescriptionComment) Usage() string {
	return "@server.description <description>"
}

// ParseInto implements Comment.
func (c *ServerDescriptionComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	desc := strings.TrimSpace(line)
	if desc == "" {
		return nil
	}

	if len(s.Spec.Servers) == 0 {
		return errors.New("cannot add server description without a preceding @server.url")
	}

	lastServer := s.Spec.Servers[len(s.Spec.Servers)-1]
	lastServer.Spec.Description = desc

	return nil
}
