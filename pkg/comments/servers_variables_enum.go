package comments

import (
	"errors"
	"fmt"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewServersVariablesEnumComment())
}

type serversVariablesEnumComment struct{}

// NewServersVariablesEnumComment creates a new serversVariablesEnumComment.
func NewServersVariablesEnumComment() *serversVariablesEnumComment {
	return &serversVariablesEnumComment{}
}

// ParseInto implements Comment.
func (s *serversVariablesEnumComment) ParseInto(c string, o *openapi.Extendable[openapi.OpenAPI]) error {
	if len(o.Spec.Servers) == 0 {
		return errors.New("use @servers.url before you use @servers.variables.enum")
	}
	server := o.Spec.Servers[len(o.Spec.Servers)-1]
	matches := serversVariablesPattern.FindStringSubmatch(c)
	if len(matches) > 0 {
		if server.Spec.Variables == nil {
			return errors.New("variables are not detected")
		}
		variable, ok := server.Spec.Variables[matches[1]]
		if !ok {
			return fmt.Errorf("variable %q is not detected", matches[1])
		}
		variable.Spec.Enum = append(variable.Spec.Enum, matches[2])
	}
	return nil
}

// Tag implements Comment.
func (s *serversVariablesEnumComment) Tag() string {
	return "servers.variables.enum"
}

// Usage implements Comment.
func (s *serversVariablesEnumComment) Usage() string {
	return `// @servers.variables.enum <variable> <value>`
}
