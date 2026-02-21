package operation

import (
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

func (c *TagsComment) ParseInto(content string, op *Operation) error {
	fields := strings.Fields(content)
	if len(fields) > 0 {
		op.Builder.AddTags(fields...)
	}
	return nil
}
