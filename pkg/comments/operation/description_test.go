package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		comment             string
		expectedDescription string
		expectedError       string
	}{
		{
			name:                "valid description",
			comment:             "My description",
			expectedDescription: "My description",
		},
		{
			name:                "empty description",
			comment:             "",
			expectedDescription: "",
		},
		{
			name:                "description with spaces",
			comment:             "  My description  ",
			expectedDescription: "My description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := operation.NewOperation()
			comment := operation.NewDescriptionComment()

			err := comment.ParseInto(tt.comment, op)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				spec := op.Builder.Build()
				assert.Equal(t, tt.expectedDescription, spec.Spec.Description)
			}
		})
	}
}

func TestDescriptionComment_Tag(t *testing.T) {
	comment := operation.NewDescriptionComment()
	assert.Equal(t, "description", comment.Tag())
}

func TestDescriptionComment_Usage(t *testing.T) {
	comment := operation.NewDescriptionComment()
	assert.Equal(t, "@description <description>", comment.Usage())
}

func TestDescriptionComment_Integration(t *testing.T) {
	op := operation.NewOperation()
	err := operation.DefaultRegistry.Parse("// @description Integrated Description", op)
	require.NoError(t, err)

	spec := op.Builder.Build()
	assert.Equal(t, "Integrated Description", spec.Spec.Description)
}
