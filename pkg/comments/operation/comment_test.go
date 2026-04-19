package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
)

func TestCommentInterface(t *testing.T) {
	// This test ensures the Comment interface is defined as expected.
	// Since it's an interface, we just check that a variable can be declared.
	var _ operation.Comment
	assert.True(t, true)
}
