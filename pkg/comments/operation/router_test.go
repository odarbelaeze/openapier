package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouterComment_ParseInto(t *testing.T) {
	tests := []struct {
		name          string
		comment       string
		expectedRoute operation.Route
		expectedError string
	}{
		{
			name:    "valid router comment",
			comment: "/users [get]",
			expectedRoute: operation.Route{
				Path:   "/users",
				Method: "get",
			},
		},
		{
			name:    "valid router comment with complex path",
			comment: "/api/v1/users/{userId}/posts [post]",
			expectedRoute: operation.Route{
				Path:   "/api/v1/users/{userId}/posts",
				Method: "post",
			},
		},
		{
			name:          "invalid router comment format - missing method",
			comment:       "/users",
			expectedError: "invalid router comment format: /users",
		},
		{
			name:          "invalid router comment format - missing path",
			comment:       "[get]",
			expectedError: "invalid router comment format: [get]",
		},
		{
			name:          "invalid router comment format - empty",
			comment:       "",
			expectedError: "invalid router comment format: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := operation.NewOperation(schema.NewResolver(nil))
			comment := operation.NewRouterComment()

			err := comment.ParseInto(tt.comment, nil, op)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.Empty(t, op.Routes) // Ensure no route is added on error
			} else {
				require.NoError(t, err)
				require.Len(t, op.Routes, 1)
				assert.Equal(t, tt.expectedRoute, op.Routes[0])
			}
		})
	}
}

func TestRouterComment_Tag(t *testing.T) {
	comment := operation.NewRouterComment()
	assert.Equal(t, "router", comment.Tag())
}

func TestRouterComment_Usage(t *testing.T) {
	comment := operation.NewRouterComment()
	assert.Equal(t, "@router <path> [<method>]", comment.Usage())
}

func TestRouterComment_Integration(t *testing.T) {
	op := operation.NewOperation(schema.NewResolver(nil))
	err := operation.DefaultRegistry.Parse("// @router /user [get]", nil, op)
	require.NoError(t, err)

	require.Len(t, op.Routes, 1)
	assert.Equal(t, "/user", op.Routes[0].Path)
	assert.Equal(t, "get", op.Routes[0].Method)
}
