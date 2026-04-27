package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestIDComment(t *testing.T) {
	comment := operation.NewIDComment()

	assert.Equal(t, "id", comment.Tag())
	assert.Equal(t, "@id <operationId>", comment.Usage())

	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "valid id",
			content:  "UploadFile",
			expected: "UploadFile",
		},
		{
			name:     "id with spaces",
			content:  "   GetUsers   ",
			expected: "GetUsers",
		},
		{
			name:     "empty id",
			content:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &operation.Operation{
				Builder: openapi.NewOperationBuilder(),
			}

			err := comment.ParseInto(tt.content, nil, op)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, op.Builder.Build().Spec.OperationID)
		})
	}
}
