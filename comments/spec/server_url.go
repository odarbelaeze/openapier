package spec

import (
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&ServerURLComment{})
}

var _ Comment = &ServerURLComment{}

// ServerURLComment is a comment that adds a new server to the global specification.
type ServerURLComment struct{}

// Tag implements Comment.
func (c *ServerURLComment) Tag() string {
	return "server.url"
}

// Usage implements Comment.
func (c *ServerURLComment) Usage() string {
	return "@server.url <url>"
}

// ParseInto implements Comment.
func (c *ServerURLComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	url := strings.TrimSpace(line)
	if url == "" {
		return nil
	}

	server := openapi.NewServerBuilder().URL(url).Build()

	if s.Spec.Servers == nil {
		s.Spec.Servers = make([]*openapi.Extendable[openapi.Server], 0)
	}

	s.Spec.Servers = append(s.Spec.Servers, server)

	return nil
}
