package operation

import (
	"errors"
	"go/ast"
	"strings"
)

func init() {
	Register(&ServerDescriptionComment{})
}

var _ Comment = &ServerDescriptionComment{}

// ServerDescriptionComment sets the description for the last added server of an operation.
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
func (c *ServerDescriptionComment) ParseInto(content string, f *ast.File, op *Operation) error {
	desc := strings.TrimSpace(content)
	if desc == "" {
		return nil
	}

	servers := op.Builder.Build().Spec.Servers
	if len(servers) == 0 {
		return errors.New("cannot add server description without a preceding @server.url")
	}

	lastServer := servers[len(servers)-1]
	lastServer.Spec.Description = desc

	return nil
}
