package operation

import (
	"strings"
)

func init() {
	Register(NewIDComment())
}

// IDComment sets the operationId of an operation.
type IDComment struct{}

func NewIDComment() *IDComment {
	return &IDComment{}
}

func (c *IDComment) Tag() string {
	return "id"
}

func (c *IDComment) Usage() string {
	return "@id <operationId>"
}

func (c *IDComment) ParseInto(content string, op *Operation) error {
	id := strings.TrimSpace(content)
	if id != "" {
		op.Builder.OperationID(id)
	}
	return nil
}
