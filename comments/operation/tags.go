package operation

import (
	"fmt"
	"go/ast"
	"strings"
)

func init() {
	Register(NewTagsComment())
}

// TagsComment adds tags to an operation.
type TagsComment struct{}

func NewTagsComment() *TagsComment {
	return &TagsComment{}
}

func (c *TagsComment) Tag() string {
	return "tags"
}

func (c *TagsComment) Usage() string {
	return "@tags <tag1> [tag2]..."
}

func (c *TagsComment) ParseInto(content string, f *ast.File, op *Operation) error {
	fields := strings.Fields(content)
	if len(fields) == 0 {
		return fmt.Errorf("invalid format for @%s, expected: %s", c.Tag(), c.Usage())
	}
	op.Builder.AddTags(fields...)
	return nil
}
