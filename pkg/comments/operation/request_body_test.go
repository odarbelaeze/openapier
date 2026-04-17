package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/schema/resolver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestRequestBodyComment_Success(t *testing.T) {
	comment := operation.NewRequestBodyComment()

	assert.Equal(t, "requestBody", comment.Tag())
	assert.Equal(t, "@requestBody <content_type> <type> [description]", comment.Usage())

	tests := []struct {
		name     string
		content  string
		validate func(t *testing.T, op *openapi.Operation)
	}{
		{
			name:    "valid form without description",
			content: "application/json string",
			validate: func(t *testing.T, op *openapi.Operation) {
				require.NotNil(t, op.RequestBody)
				req := op.RequestBody.Spec.Spec
				assert.True(t, req.Required)
				assert.Empty(t, req.Description)
				require.NotNil(t, req.Content["application/json"])
				mediaType := req.Content["application/json"].Spec
				assert.Equal(t, "string", (*mediaType.Schema.Spec.Type)[0])
			},
		},
		{
			name:    "valid form with description",
			content: "application/json string The identifier",
			validate: func(t *testing.T, op *openapi.Operation) {
				require.NotNil(t, op.RequestBody)
				req := op.RequestBody.Spec.Spec
				assert.True(t, req.Required)
				assert.Equal(t, "The identifier", req.Description)
				require.NotNil(t, req.Content["application/json"])
				mediaType := req.Content["application/json"].Spec
				assert.Equal(t, "string", (*mediaType.Schema.Spec.Type)[0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := resolver.NewMockResolver(t)
			resolver.EXPECT().Resolve("string").Return(openapi.NewSchemaBuilder().Type("string").Build(), nil)
			op := operation.NewOperation(resolver)
			err := comment.ParseInto(tt.content, nil, op)
			require.NoError(t, err)
			if tt.validate != nil {
				tt.validate(t, op.Builder.Build().Spec)
			}
		})
	}
}

func TestRequestBodyComment_Failure(t *testing.T) {
	comment := operation.NewRequestBodyComment()
	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "invalid format",
			content: "application/json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := operation.NewOperation(nil)
			err := comment.ParseInto(tt.content, nil, op)
			assert.Error(t, err)
		})
	}
}
