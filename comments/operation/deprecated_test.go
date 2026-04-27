package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestDeprecatedComment(t *testing.T) {
	comment := operation.NewDeprecatedComment()

	assert.Equal(t, "deprecated", comment.Tag())
	assert.Equal(t, "@deprecated", comment.Usage())

	op := &operation.Operation{
		Builder: openapi.NewOperationBuilder(),
	}

	assert.False(t, op.Builder.Build().Spec.Deprecated)

	err := comment.ParseInto("", nil, op)
	require.NoError(t, err)

	assert.True(t, op.Builder.Build().Spec.Deprecated)
}
