package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestParamAllowReservedComment(t *testing.T) {
	comment := operation.NewParamAllowReservedComment()
	assert.Equal(t, "param.allowReserved", comment.Tag())
	assert.Equal(t, "@param.allowReserved <param> [param]...", comment.Usage())

	op := operation.NewOperation(nil)
	op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
	op.Builder.AddParameters(openapi.NewParameterBuilder().Name("other").Build())

	err := comment.ParseInto("id other", nil, op)
	require.NoError(t, err)

	spec := op.Builder.Build().Spec
	require.Len(t, spec.Parameters, 2)
	assert.True(t, spec.Parameters[0].Spec.Spec.AllowReserved)
	assert.True(t, spec.Parameters[1].Spec.Spec.AllowReserved)
}
