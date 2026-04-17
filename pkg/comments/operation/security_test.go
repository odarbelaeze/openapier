package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestSecurityComment_Tag(t *testing.T) {
	comment := operation.NewSecurityComment()
	assert.Equal(t, "security", comment.Tag())
}

func TestSecurityComment_Usage(t *testing.T) {
	comment := operation.NewSecurityComment()
	assert.Equal(t, "@security <name> [scope1] [scope2] ...", comment.Usage())
}

func TestSecurityComment_ParseInto(t *testing.T) {
	tests := []struct {
		name             string
		line             string
		expectedSecurity []openapi.SecurityRequirement
	}{
		{
			name:             "empty line",
			line:             "",
			expectedSecurity: nil,
		},
		{
			name: "single scheme without scopes",
			line: "apiKey",
			expectedSecurity: []openapi.SecurityRequirement{
				{"apiKey": []string(nil)},
			},
		},
		{
			name: "single scheme with scopes",
			line: "oauth2 read:users write:users",
			expectedSecurity: []openapi.SecurityRequirement{
				{"oauth2": []string{"read:users", "write:users"}},
			},
		},
		{
			name: "multiple spaces",
			line: "oauth2   read:users    write:users  ",
			expectedSecurity: []openapi.SecurityRequirement{
				{"oauth2": []string{"read:users", "write:users"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := operation.NewSecurityComment()
			op := operation.NewOperation(nil)
			err := comment.ParseInto(tt.line, nil, op)

			require.NoError(t, err)

			spec := op.Builder.Build()

			if len(tt.expectedSecurity) == 0 {
				assert.Empty(t, spec.Spec.Security)
			} else {
				assert.Equal(t, tt.expectedSecurity, spec.Spec.Security)
			}
		})
	}
}

func TestSecurityComment_Integration(t *testing.T) {
	op := operation.NewOperation(nil)
	err1 := operation.Default().Parse("// @security apiKey", nil, op)
	require.NoError(t, err1)
	err2 := operation.Default().Parse("// @security oauth2 read:users write:users", nil, op)
	require.NoError(t, err2)

	spec := op.Builder.Build()

	expectedSecurity := []openapi.SecurityRequirement{
		{"apiKey": []string(nil)},
		{"oauth2": []string{"read:users", "write:users"}},
	}
	assert.Equal(t, expectedSecurity, spec.Spec.Security)
}
