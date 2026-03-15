package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestExternalDocsDescriptionComment(t *testing.T) {
	comment := operation.NewExternalDocsDescriptionComment()

	assert.Equal(t, "externaldocs.description", comment.Tag())
	assert.Equal(t, "@externalDocs.description <description>", comment.Usage())

	tests := []struct {
		name                string
		content             string
		setupDocs           bool
		expectedDescription string
		expectedURL         string
	}{
		{
			name:                "valid description",
			content:             "Find more info here",
			setupDocs:           false,
			expectedDescription: "Find more info here",
		},
		{
			name:                "description with spaces",
			content:             "   More detailed docs   ",
			setupDocs:           false,
			expectedDescription: "More detailed docs",
		},
		{
			name:                "empty content",
			content:             "",
			setupDocs:           false,
			expectedDescription: "",
		},
		{
			name:                "existing external docs",
			content:             "Updated explanation",
			setupDocs:           true,
			expectedDescription: "Updated explanation",
			expectedURL:         "https://example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &operation.Operation{
				Builder: openapi.NewOperationBuilder(),
			}

			if tt.setupDocs {
				op.Builder.ExternalDocs(openapi.NewExternalDocsBuilder().URL("https://example.com").Build())
			}

			err := comment.ParseInto(tt.content, nil, op)
			require.NoError(t, err)

			actual := op.Builder.Build().Spec.ExternalDocs
			if tt.expectedDescription == "" && !tt.setupDocs {
				assert.Nil(t, actual)
			} else {
				require.NotNil(t, actual)
				assert.Equal(t, tt.expectedDescription, actual.Spec.Description)
				assert.Equal(t, tt.expectedURL, actual.Spec.URL)
			}
		})
	}
}
