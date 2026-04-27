package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestDescriptionComment(t *testing.T) {
	comment := operation.NewDescriptionComment()

	assert.Equal(t, "description", comment.Tag())
	assert.Equal(t, "@description <description>", comment.Usage())

	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "valid description",
			content:  "This endpoint uploads a new file.",
			expected: "This endpoint uploads a new file.",
		},
		{
			name:     "description with spaces",
			content:  "   Detailed user fetch logic   ",
			expected: "Detailed user fetch logic",
		},
		{
			name:     "empty description",
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

			assert.Equal(t, tt.expected, op.Builder.Build().Spec.Description)
		})
	}
}
