package operation

import "strings"

func init() {
	Register(NewDescriptionComment())
}

// DescriptionComment is a comment that updates the description of an operation.
type DescriptionComment struct{}

func NewDescriptionComment() *DescriptionComment {
	return &DescriptionComment{}
}

func (d *DescriptionComment) Tag() string {
	return "description"
}

func (d *DescriptionComment) Usage() string {
	return "@description <description>"
}

func (d *DescriptionComment) ParseInto(content string, op *Operation) error {
	op.Builder.Description(strings.TrimSpace(content))
	return nil
}
