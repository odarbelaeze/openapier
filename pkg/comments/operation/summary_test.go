package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestSummaryComment(t *testing.T) {
	comment := operation.NewSummaryComment()

	assert.Equal(t, "summary", comment.Tag())
	assert.Equal(t, "@summary <summary>", comment.Usage())

	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "valid summary",
			content:  "Uploads a new file",
			expected: "Uploads a new file",
		},
		{
			name:     "summary with spaces",
			content:  "   Get all users   ",
			expected: "Get all users",
		},
		{
			name:     "empty summary",
			content:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &operation.Operation{
				Builder: openapi.NewOperationBuilder(),
			}

			err := comment.ParseInto(tt.content, op)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, op.Builder.Build().Spec.Summary)
		})
	}
}
