package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestParamDeprecatedComment(t *testing.T) {
	comment := operation.NewParamDeprecatedComment()
	assert.Equal(t, "param.deprecated", comment.Tag())
	assert.Equal(t, "@param.deprecated <param> [param]...", comment.Usage())

	op := operation.NewOperation(nil)
	op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())

	err := comment.ParseInto("id", nil, op)
	require.NoError(t, err)

	spec := op.Builder.Build().Spec
	require.Len(t, spec.Parameters, 1)
	assert.True(t, spec.Parameters[0].Spec.Spec.Deprecated)
}
