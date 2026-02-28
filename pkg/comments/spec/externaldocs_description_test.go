package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestExternalDocsDescriptionComment_Tag(t *testing.T) {
	comment := &spec.ExternalDocsDescriptionComment{}
	assert.Equal(t, "externalDocs.description", comment.Tag())
}

func TestExternalDocsDescriptionComment_Usage(t *testing.T) {
	comment := &spec.ExternalDocsDescriptionComment{}
	assert.Equal(t, "@externalDocs.description <description>", comment.Usage())
}

func TestExternalDocsDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		line                string
		expectedDescription string
	}{
		{
			name:                "valid description",
			line:                "Find more info here",
			expectedDescription: "Find more info here",
		},
		{
			name:                "empty description",
			line:                "",
			expectedDescription: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.ExternalDocsDescriptionComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.ExternalDocs)
			assert.Equal(t, tt.expectedDescription, o.Spec.ExternalDocs.Spec.Description)
		})
	}
}
