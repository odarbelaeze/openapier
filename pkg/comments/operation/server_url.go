package operation

import (
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&ServerURLComment{})
}

var _ Comment = &ServerURLComment{}

// ServerURLComment is a comment that adds a new server to the operation.
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
func (c *ServerURLComment) ParseInto(content string, op *Operation) error {
	url := strings.TrimSpace(content)
	if url == "" {
		return nil
	}

	server := openapi.NewServerBuilder().URL(url).Build()
	op.Builder.AddServers(server)

	return nil
}
