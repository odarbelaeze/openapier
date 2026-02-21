package spec

import (
	"errors"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&ServerVariableDefaultComment{})
}

var _ Comment = &ServerVariableDefaultComment{}

// ServerVariableDefaultComment sets the default value for a server variable.
type ServerVariableDefaultComment struct{}

// Tag implements Comment.
func (c *ServerVariableDefaultComment) Tag() string {
	return "server.variable.default"
}

// Usage implements Comment.
func (c *ServerVariableDefaultComment) Usage() string {
	return "@server.variable.default <variable> <default>"
}

// ParseInto implements Comment.
func (c *ServerVariableDefaultComment) ParseInto(line string, s *openapi.Extendable[openapi.OpenAPI]) error {
	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) != 2 {
		return errors.New("invalid format for @server.variable.default, expected: <variable> <default>")
	}

	variable := parts[0]
	defaultValue := strings.TrimSpace(parts[1])

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

	lastServer.Spec.Variables[variable].Spec.Default = defaultValue

	return nil
}
