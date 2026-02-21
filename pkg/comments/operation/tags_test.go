package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagsComment(t *testing.T) {
	comment := operation.NewTagsComment()

	assert.Equal(t, "tags", comment.Tag())
	assert.Equal(t, "@tags <tag1> [tag2]...", comment.Usage())

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "single tag",
			content:  "users",
			expected: []string{"users"},
		},
		{
			name:     "multiple tags space-separated",
			content:  "users admins auth",
			expected: []string{"users", "admins", "auth"},
		},
		{
			name:     "multiple tags with irregular spaces",
			content:  "   roles   permissions  ",
			expected: []string{"roles", "permissions"},
		},
		{
			name:     "empty content",
			content:  "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &operation.Operation{
				Builder: openapi.NewOperationBuilder(),
			}

			err := comment.ParseInto(tt.content, op)
			require.NoError(t, err)

			actual := op.Builder.Build().Spec.Tags
			if len(tt.expected) == 0 {
				assert.Empty(t, actual)
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestTagsComment_MultipleAppends(t *testing.T) {
	comment := operation.NewTagsComment()

	op := &operation.Operation{
		Builder: openapi.NewOperationBuilder(),
	}

	err1 := comment.ParseInto("users", op)
	require.NoError(t, err1)

	err2 := comment.ParseInto("admins auth", op)
	require.NoError(t, err2)

	assert.Equal(t, []string{"users", "admins", "auth"}, op.Builder.Build().Spec.Tags)
}
