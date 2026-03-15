package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestRequestBodyComment(t *testing.T) {
	comment := operation.NewRequestBodyComment()

	assert.Equal(t, "requestBody", comment.Tag())
	assert.Equal(t, "@requestBody <content_type> <type> [description]", comment.Usage())

	tests := []struct {
		name        string
		content     string
		expectError bool
		validate    func(t *testing.T, op *openapi.Operation)
	}{
		{
			name:        "invalid format",
			content:     "application/json",
			expectError: true,
		},
		{
			name:        "valid form without description",
			content:     "application/json string",
			expectError: false,
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
			name:        "valid form with description",
			content:     "application/json int The identifier",
			expectError: false,
			validate: func(t *testing.T, op *openapi.Operation) {
				require.NotNil(t, op.RequestBody)
				req := op.RequestBody.Spec.Spec
				assert.True(t, req.Required)
				assert.Equal(t, "The identifier", req.Description)
				require.NotNil(t, req.Content["application/json"])
				mediaType := req.Content["application/json"].Spec
				assert.Equal(t, "integer", (*mediaType.Schema.Spec.Type)[0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := operation.NewOperation(schema.NewResolver())
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
