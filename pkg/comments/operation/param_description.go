package operation

import (
	"fmt"
	"go/ast"
	"strings"
)

func init() {
	Register(NewParamDescriptionComment())
}

// ParamDescriptionComment sets the description for a parameter in an operation.
type ParamDescriptionComment struct{}

func NewParamDescriptionComment() *ParamDescriptionComment {
	return &ParamDescriptionComment{}
}

func (c *ParamDescriptionComment) Tag() string {
	return "param.description"
}

func (c *ParamDescriptionComment) Usage() string {
	return "@param.description <name> <description>"
}

func (c *ParamDescriptionComment) ParseInto(content string, f *ast.File, op *Operation) error {
	fields := strings.Fields(content)
	if len(fields) < 2 {
		return fmt.Errorf("invalid @param.description format, expected: %s", c.Usage())
	}

	name := fields[0]
	description := strings.Join(fields[1:], " ")

	for _, p := range op.Builder.Build().Spec.Parameters {
		if p.Spec.Spec.Name == name {
			p.Spec.Spec.Description = description
			return nil
		}
	}

	return fmt.Errorf("parameter %q not found, use @param %s ... to define it first", name, name)
}
