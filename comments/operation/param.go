package operation

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewParamComment())
}

// ParamComment adds a parameter to an operation.
type ParamComment struct{}

func NewParamComment() *ParamComment {
	return &ParamComment{}
}

func (c *ParamComment) Tag() string {
	return "param"
}

func (c *ParamComment) Usage() string {
	return "@param <name> <type> <in> [description...]"
}

func (c *ParamComment) ParseInto(content string, f *ast.File, op *Operation) error {
	fields := strings.Fields(content)
	if len(fields) < 3 {
		return fmt.Errorf("invalid @param format, expected: %s", c.Usage())
	}
	name := fields[0]
	typ := fields[1]
	in := fields[2]
	description := strings.Join(fields[3:], " ")

	paramSchema, err := op.Resolver.Resolve(typ)
	if err != nil {
		return fmt.Errorf("failed to resolve type %q: %w", typ, err)
	}

	builder := openapi.NewParameterBuilder().
		Name(name).
		In(in).
		Schema(paramSchema)

	if description != "" {
		builder.Description(description)
	}

	if in == openapi.InPath {
		builder.Required(true)
	}

	op.Builder.AddParameters(builder.Build())

	return nil
}
