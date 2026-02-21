package operation

import (
	"strings"
)

func init() {
	Register(NewDescriptionComment())
}

// DescriptionComment sets the description of an operation.
type DescriptionComment struct{}

func NewDescriptionComment() *DescriptionComment {
	return &DescriptionComment{}
}

func (c *DescriptionComment) Tag() string {
	return "description"
}

func (c *DescriptionComment) Usage() string {
	return "@description <description>"
}

func (c *DescriptionComment) ParseInto(content string, op *Operation) error {
	desc := strings.TrimSpace(content)
	if desc != "" {
		op.Builder.Description(desc)
	}
	return nil
}
