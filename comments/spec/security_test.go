package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestSecurityComment_Tag(t *testing.T) {
	comment := &spec.SecurityComment{}
	assert.Equal(t, "security", comment.Tag())
}

func TestSecurityComment_Usage(t *testing.T) {
	comment := &spec.SecurityComment{}
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
			comment := &spec.SecurityComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)

			if len(tt.expectedSecurity) == 0 {
				assert.Empty(t, o.Spec.Security)
			} else {
				assert.Equal(t, tt.expectedSecurity, o.Spec.Security)
			}
		})
	}
}

func TestSecurityComment_ParseInto_Multiple(t *testing.T) {
	comment := &spec.SecurityComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	err1 := comment.ParseInto("apiKey", o)
	require.NoError(t, err1)

	err2 := comment.ParseInto("oauth2 read:users write:users", o)
	require.NoError(t, err2)

	expectedSecurity := []openapi.SecurityRequirement{
		{"apiKey": []string(nil)},
		{"oauth2": []string{"read:users", "write:users"}},
	}
	assert.Equal(t, expectedSecurity, o.Spec.Security)
}
