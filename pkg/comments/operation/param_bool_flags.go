package operation

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/sv-tools/openapi"
)

// paramBoolComment provides a common implementation for `@param.<flag>` comments.
type paramBoolComment struct {
	tag    string
	setter func(param *openapi.Parameter)
}

func (c *paramBoolComment) Tag() string {
	return c.tag
}

func (c *paramBoolComment) Usage() string {
	return fmt.Sprintf("@%s <param> [param]...", c.tag)
}

func (c *paramBoolComment) ParseInto(content string, f *ast.File, op *Operation) error {
	fields := strings.Fields(content)
	if len(fields) == 0 {
		return fmt.Errorf("invalid @%s format, expected: %s", c.tag, c.Usage())
	}

	for _, name := range fields {
		found := false
		for _, p := range op.Builder.Build().Spec.Parameters {
			if p.Spec.Spec.Name == name {
				c.setter(p.Spec.Spec)
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("parameter %q not found, use @param %s ... to define it first", name, name)
		}
	}

	return nil
}
