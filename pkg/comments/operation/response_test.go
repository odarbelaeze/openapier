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

func TestResponseComment_Success(t *testing.T) {
	comment := operation.NewResponseComment()

	assert.Equal(t, "response", comment.Tag())
	assert.Equal(t, "@response <status_code> <content_type> <type> [description]", comment.Usage())

	tests := []struct {
		name     string
		content  string
		validate func(t *testing.T, responses *openapi.Responses)
	}{
		{
			name:    "valid form without description",
			content: "200 application/json string",
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
			name:    "valid form with description",
			content: "500 application/json string Error details",
			validate: func(t *testing.T, responses *openapi.Responses) {
				require.NotNil(t, responses.Response["500"])
				resp := responses.Response["500"].Spec.Spec
				assert.Equal(t, "Error details", resp.Description)
				require.NotNil(t, resp.Content["application/json"])
				mediaType := resp.Content["application/json"].Spec
				assert.Equal(t, "string", (*mediaType.Schema.Spec.Type)[0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Mock resolver
			controller := gomock.NewController(t)
			resolver := resolver.NewMockResolver(controller)
			resolver.EXPECT().Resolve("string").Return(openapi.NewSchemaBuilder().Type("string").Build(), nil)
			op := operation.NewOperation(resolver)
			err := comment.ParseInto(tt.content, nil, op)
			require.NoError(t, err)
			if tt.validate != nil {
				tt.validate(t, op.ResponsesBuilder.Build().Spec.Spec)
			}
		})
	}
}

func TestResponseComment_Failure(t *testing.T) {
	comment := operation.NewResponseComment()
	op := operation.NewOperation(nil)
	err := comment.ParseInto("invalid", nil, op)
	assert.Error(t, err)
}
