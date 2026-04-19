package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestSecuritySchemeComment_Tag(t *testing.T) {
	comment := &spec.SecuritySchemeComment{}
	assert.Equal(t, "securityScheme", comment.Tag())
}

func TestSecuritySchemeComment_Usage(t *testing.T) {
	comment := &spec.SecuritySchemeComment{}
	assert.Equal(t, "@securityScheme <name> <type> [<description>...]", comment.Usage())
}

func TestSecuritySchemeComment_ParseInto(t *testing.T) {
	tests := []struct {
		name           string
		line           string
		expectedName   string
		expectedType   string
		expectedDesc   string
		expectedError  bool
	}{
		{
			name:          "too few fields",
			line:          "name",
			expectedError: true,
		},
		{
			name:          "basic scheme",
			line:          "myAuth apiKey",
			expectedName:  "myAuth",
			expectedType:  "apiKey",
			expectedError: false,
		},
		{
			name:          "scheme with description",
			line:          "myAuth oauth2 This is a description",
			expectedName:  "myAuth",
			expectedType:  "oauth2",
			expectedDesc:  "This is a description",
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.SecuritySchemeComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			if tt.expectedError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Components)
			require.NotNil(t, o.Spec.Components.Spec.SecuritySchemes)

			scheme, ok := o.Spec.Components.Spec.SecuritySchemes[tt.expectedName]
			require.True(t, ok)
			assert.Equal(t, tt.expectedType, scheme.Spec.Spec.Type)
			assert.Equal(t, tt.expectedDesc, scheme.Spec.Spec.Description)
		})
	}
}

func TestSecuritySchemeComment_ParseInto_Multiple(t *testing.T) {
	comment := &spec.SecuritySchemeComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	err1 := comment.ParseInto("auth1 apiKey description 1", o)
	require.NoError(t, err1)

	err2 := comment.ParseInto("auth2 http description 2", o)
	require.NoError(t, err2)

	require.NotNil(t, o.Spec.Components)
	require.NotNil(t, o.Spec.Components.Spec.SecuritySchemes)
	assert.Len(t, o.Spec.Components.Spec.SecuritySchemes, 2)

	assert.Equal(t, "apiKey", o.Spec.Components.Spec.SecuritySchemes["auth1"].Spec.Spec.Type)
	assert.Equal(t, "description 1", o.Spec.Components.Spec.SecuritySchemes["auth1"].Spec.Spec.Description)
	assert.Equal(t, "http", o.Spec.Components.Spec.SecuritySchemes["auth2"].Spec.Spec.Type)
	assert.Equal(t, "description 2", o.Spec.Components.Spec.SecuritySchemes["auth2"].Spec.Spec.Description)
}
