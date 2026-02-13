package spec

import (
	"errors"
	"fmt"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewServersVariablesDefaultComment())
}

type serversVariablesDefaultComment struct{}

// NewServersVariablesDefaultComment creates a new serversVariablesDefaultComment.
func NewServersVariablesDefaultComment() *serversVariablesDefaultComment {
	return &serversVariablesDefaultComment{}
}

// ParseInto implements Comment.
func (s *serversVariablesDefaultComment) ParseInto(c string, o *openapi.Extendable[openapi.OpenAPI]) error {
	if len(o.Spec.Servers) == 0 {
		return errors.New("use @servers.url before you use @servers.variables.default")
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
		variable.Spec.Default = matches[2]
	}
	return nil
}

// Tag implements Comment.
func (s *serversVariablesDefaultComment) Tag() string {
	return "servers.variables.default"
}

// Usage implements Comment.
func (s *serversVariablesDefaultComment) Usage() string {
	return `// @servers.variables.default <variable> <default>`
}
