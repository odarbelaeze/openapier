package spec

import (
	"errors"
	"fmt"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewServersVariablesDescriptionComment())
}

type serversVariablesDescriptionComment struct{}

// NewServersVariablesDescriptionComment creates a new serversVariablesDescriptionComment.
func NewServersVariablesDescriptionComment() *serversVariablesDescriptionComment {
	return &serversVariablesDescriptionComment{}
}

// ParseInto implements Comment.
func (s *serversVariablesDescriptionComment) ParseInto(c string, o *openapi.Extendable[openapi.OpenAPI]) error {
	if len(o.Spec.Servers) == 0 {
		return errors.New("use @servers.url before you use @servers.variables.description")
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
		variable.Spec.Description = matches[2]
	}
	return nil
}

// Tag implements Comment.
func (s *serversVariablesDescriptionComment) Tag() string {
	return "servers.variables.description"
}

// Usage implements Comment.
func (s *serversVariablesDescriptionComment) Usage() string {
	return `// @servers.variables.description <variable> <description>`
}
