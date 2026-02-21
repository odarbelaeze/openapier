package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestExternalDocsURLComment(t *testing.T) {
	comment := operation.NewExternalDocsURLComment()

	assert.Equal(t, "externaldocs.url", comment.Tag())
	assert.Equal(t, "@externalDocs.url <url>", comment.Usage())

	tests := []struct {
		name                string
		content             string
		setupDocs           bool
		expectedURL         string
		expectedDescription string
	}{
		{
			name:        "valid url",
			content:     "https://example.com/docs",
			setupDocs:   false,
			expectedURL: "https://example.com/docs",
		},
		{
			name:        "url with spaces",
			content:     "   https://example.com/api-docs   ",
			setupDocs:   false,
			expectedURL: "https://example.com/api-docs",
		},
		{
			name:        "empty content",
			content:     "",
			setupDocs:   false,
			expectedURL: "",
		},
		{
			name:                "existing external docs",
			content:             "https://example.com/new-docs",
			setupDocs:           true,
			expectedURL:         "https://example.com/new-docs",
			expectedDescription: "Existing description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &operation.Operation{
				Builder: openapi.NewOperationBuilder(),
			}

			if tt.setupDocs {
				op.Builder.ExternalDocs(openapi.NewExternalDocsBuilder().Description("Existing description").Build())
			}

			err := comment.ParseInto(tt.content, op)
			require.NoError(t, err)

			actual := op.Builder.Build().Spec.ExternalDocs
			if tt.expectedURL == "" && !tt.setupDocs {
				assert.Nil(t, actual)
			} else {
				require.NotNil(t, actual)
				assert.Equal(t, tt.expectedURL, actual.Spec.URL)
				assert.Equal(t, tt.expectedDescription, actual.Spec.Description)
			}
		})
	}
}
