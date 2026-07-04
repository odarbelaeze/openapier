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
	names := make(map[string]struct{})
	for _, name := range fields {
		names[name] = struct{}{}
	}
	for _, p := range op.Builder.Build().Spec.Parameters {
		name := p.Spec.Spec.Name
		if _, ok := names[name]; ok {
			c.setter(p.Spec.Spec)
			delete(names, name)
		}
	}
	if len(names) > 0 {
		var missing []string
		for name := range names {
			missing = append(missing, name)
		}
		return fmt.Errorf("parameters not found for @%s: %v, use @param <name> ... to define them first", c.tag, missing)
	}
	return nil
}
