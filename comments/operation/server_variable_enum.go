package operation

import (
	"errors"
	"go/ast"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(&ServerVariableEnumComment{})
}

var _ Comment = &ServerVariableEnumComment{}

// ServerVariableEnumComment sets the enum values for a server variable on an operation.
type ServerVariableEnumComment struct{}

// Tag implements Comment.
func (c *ServerVariableEnumComment) Tag() string {
	return "server.variable.enum"
}

// Usage implements Comment.
func (c *ServerVariableEnumComment) Usage() string {
	return "@server.variable.enum <variable> [value1] [value2] ..."
}

// ParseInto implements Comment.
func (c *ServerVariableEnumComment) ParseInto(content string, f *ast.File, op *Operation) error {
	parts := strings.Fields(content)
	if len(parts) < 1 { // Variable name is implicitly required, further values are optional in parser logic but swaggo prefers it
		return nil
	}

	variable := parts[0]
	var enums []string
	if len(parts) > 1 {
		enums = parts[1:]
	}

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

	// We append rather than overwrite to support multiple lines
	lastServer.Spec.Variables[variable].Spec.Enum = append(lastServer.Spec.Variables[variable].Spec.Enum, enums...)

	return nil
}
