package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestResponseComment(t *testing.T) {
	comment := operation.NewResponseComment()

	assert.Equal(t, "response", comment.Tag())
	assert.Equal(t, "@response <status_code> <content_type> <type> [description]", comment.Usage())

	tests := []struct {
		name        string
		content     string
		expectError bool
		validate    func(t *testing.T, responses *openapi.Responses)
	}{
		{
			name:        "invalid format",
			content:     "200 application/json",
			expectError: true,
		},
		{
			name:        "valid form without description",
			content:     "200 application/json string",
			expectError: false,
			validate: func(t *testing.T, responses *openapi.Responses) {
				require.NotNil(t, responses.Response["200"])
				resp := responses.Response["200"].Spec.Spec
				assert.Empty(t, resp.Description)
				require.NotNil(t, resp.Content["application/json"])
				mediaType := resp.Content["application/json"].Spec
				assert.Equal(t, "string", (*mediaType.Schema.Spec.Type)[0])
			},
		},
		{
			name:        "valid form with description",
			content:     "500 application/json int Error details",
			expectError: false,
			validate: func(t *testing.T, responses *openapi.Responses) {
				require.NotNil(t, responses.Response["500"])
				resp := responses.Response["500"].Spec.Spec
				assert.Equal(t, "Error details", resp.Description)
				require.NotNil(t, resp.Content["application/json"])
				mediaType := resp.Content["application/json"].Spec
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
					tt.validate(t, op.ResponsesBuilder.Build().Spec.Spec)
				}
			}
		})
	}
}
