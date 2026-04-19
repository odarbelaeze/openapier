package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecurityNoneComment_Tag(t *testing.T) {
	comment := operation.NewSecurityNoneComment()
	assert.Equal(t, "security.none", comment.Tag())
}

func TestSecurityNoneComment_Usage(t *testing.T) {
	comment := operation.NewSecurityNoneComment()
	assert.Equal(t, "@security.none", comment.Usage())
}

func TestSecurityNoneComment_ParseInto(t *testing.T) {
	comment := operation.NewSecurityNoneComment()
	op := operation.NewOperation(nil)

	// Add some security first
	secComment := operation.NewSecurityComment()
	require.NoError(t, secComment.ParseInto("apiKey", nil, op))
	spec := op.Builder.Build()
	assert.Len(t, spec.Spec.Security, 1)

	// Clear it
	err := comment.ParseInto("", nil, op)
	require.NoError(t, err)

	spec = op.Builder.Build()
	assert.Empty(t, spec.Spec.Security)
}
