package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestParamComment(t *testing.T) {
	comment := operation.NewParamComment()

	assert.Equal(t, "param", comment.Tag())
	assert.Equal(t, "@param <name> <type> <in> [description...]", comment.Usage())

	tests := []struct {
		name        string
		content     string
		expectError bool
		validate    func(t *testing.T, op *openapi.Operation)
	}{
		{
			name:    "a path parameter",
			content: "id int path",
			validate: func(t *testing.T, op *openapi.Operation) {
				require.Len(t, op.Parameters, 1)
				param := op.Parameters[0].Spec.Spec
				assert.Equal(t, "id", param.Name)
				assert.Equal(t, "path", param.In)
				assert.True(t, param.Required)
				require.NotNil(t, param.Schema)
			},
		},
		{
			name:    "a query parameter",
			content: "limit int query",
			validate: func(t *testing.T, op *openapi.Operation) {
				require.Len(t, op.Parameters, 1)
				param := op.Parameters[0].Spec.Spec
				assert.Equal(t, "limit", param.Name)
				assert.Equal(t, "query", param.In)
				assert.False(t, param.Required)
				require.NotNil(t, param.Schema)
			},
		},
		{
			name:        "invalid format",
			content:     "id int",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := operation.NewOperation(schema.NewResolver(nil))

			err := comment.ParseInto(tt.content, nil, op)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, op.Builder.Build().Spec)
				}
			}
		})
	}
}
