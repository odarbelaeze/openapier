package operation

import (
	"errors"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&ServerVariableDescriptionComment{})
}

var _ Comment = &ServerVariableDescriptionComment{}

// ServerVariableDescriptionComment sets the description for a server variable on an operation.
type ServerVariableDescriptionComment struct{}

// Tag implements Comment.
func (c *ServerVariableDescriptionComment) Tag() string {
	return "server.variable.description"
}

// Usage implements Comment.
func (c *ServerVariableDescriptionComment) Usage() string {
	return "@server.variable.description <variable> <description>"
}

// ParseInto implements Comment.
func (c *ServerVariableDescriptionComment) ParseInto(content string, op *Operation) error {
	parts := strings.SplitN(strings.TrimSpace(content), " ", 2)
	if len(parts) != 2 {
		return errors.New("invalid format for @server.variable.description, expected: <variable> <description>")
	}

	variable := parts[0]
	desc := strings.TrimSpace(parts[1])

	servers := op.Builder.Build().Spec.Servers
	if len(servers) == 0 {
		return errors.New("cannot add server variable without a preceding @server.url")
	}

	lastServer := servers[len(servers)-1]

	if lastServer.Spec.Variables == nil {
		lastServer.Spec.Variables = make(map[string]*openapi.Extendable[openapi.ServerVariable])
	}

	if _, exists := lastServer.Spec.Variables[variable]; !exists {
		lastServer.Spec.Variables[variable] = openapi.NewServerVariableBuilder().Build()
	}

	lastServer.Spec.Variables[variable].Spec.Description = desc

	return nil
}
