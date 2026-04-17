package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/schema/resolver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
	"go.uber.org/mock/gomock"
)

func TestParamComment_Success(t *testing.T) {
	comment := operation.NewParamComment()

	assert.Equal(t, "param", comment.Tag())
	assert.Equal(t, "@param <name> <type> <in> [description...]", comment.Usage())

	tests := []struct {
		name     string
		content  string
		validate func(t *testing.T, op *openapi.Operation)
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			resolver := resolver.NewMockResolver(controller)
			resolver.EXPECT().Resolve("int").Return(openapi.NewSchemaBuilder().Type("number").Build(), nil)
			op := operation.NewOperation(resolver)
			err := comment.ParseInto(tt.content, nil, op)
			require.NoError(t, err)
			if tt.validate != nil {
				tt.validate(t, op.Builder.Build().Spec)
			}
		})
	}
}

func TestParamComment_InvalidFormat(t *testing.T) {
	comment := operation.NewParamComment()
	op := operation.NewOperation(nil)
	err := comment.ParseInto("id int", nil, op)
	assert.Error(t, err)
}
