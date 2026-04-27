package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestParamAllowEmptyValueComment(t *testing.T) {
	comment := operation.NewParamAllowEmptyValueComment()
	assert.Equal(t, "param.allowEmptyValue", comment.Tag())
	assert.Equal(t, "@param.allowEmptyValue <param> [param]...", comment.Usage())

	op := operation.NewOperation(nil)
	op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())

	err := comment.ParseInto("id", nil, op)
	require.NoError(t, err)

	spec := op.Builder.Build().Spec
	require.Len(t, spec.Parameters, 1)
	assert.True(t, spec.Parameters[0].Spec.Spec.AllowEmptyValue)
}
