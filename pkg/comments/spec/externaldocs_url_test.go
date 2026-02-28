package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestExternalDocsURLComment_Tag(t *testing.T) {
	comment := &spec.ExternalDocsURLComment{}
	assert.Equal(t, "externalDocs.url", comment.Tag())
}

func TestExternalDocsURLComment_Usage(t *testing.T) {
	comment := &spec.ExternalDocsURLComment{}
	assert.Equal(t, "@externalDocs.url <url>", comment.Usage())
}

func TestExternalDocsURLComment_ParseInto(t *testing.T) {
	tests := []struct {
		name        string
		line        string
		expectedURL string
	}{
		{
			name:        "valid url",
			line:        "https://example.com/docs",
			expectedURL: "https://example.com/docs",
		},
		{
			name:        "empty url",
			line:        "",
			expectedURL: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.ExternalDocsURLComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.ExternalDocs)
			assert.Equal(t, tt.expectedURL, o.Spec.ExternalDocs.Spec.URL)
		})
	}
}
