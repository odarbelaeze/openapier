package spec

import (
	"errors"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&ServerVariableDescriptionComment{})
}

var _ Comment = &ServerVariableDescriptionComment{}

// ServerVariableDescriptionComment sets the description for a server variable.
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
func (c *ServerVariableDescriptionComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) != 2 {
		return errors.New("invalid format for @server.variable.description, expected: <variable> <description>")
	}

	variable := parts[0]
	desc := strings.TrimSpace(parts[1])

	if len(s.Spec.Servers) == 0 {
		return errors.New("cannot add server variable without a preceding @server.url")
	}

	lastServer := s.Spec.Servers[len(s.Spec.Servers)-1]

	if lastServer.Spec.Variables == nil {
		lastServer.Spec.Variables = make(map[string]*openapi.Extendable[openapi.ServerVariable])
	}

	if _, exists := lastServer.Spec.Variables[variable]; !exists {
		lastServer.Spec.Variables[variable] = openapi.NewServerVariableBuilder().Build()
	}

	lastServer.Spec.Variables[variable].Spec.Description = desc

	return nil
}
