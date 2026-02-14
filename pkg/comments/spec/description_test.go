package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestDescription_ParseInto(t *testing.T) {
	tests := []struct {
		name        string
		description string
		expected    string
	}{
		{
			name:        "simple description",
			description: "My API Description",
			expected:    "My API Description",
		},
		{
			name:        "description with markdown",
			description: "My **API** Description",
			expected:    "My **API** Description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := spec.NewDescriptionComment()
			o := openapi.NewOpenAPIBuilder().Build()

			err := comment.ParseInto(tt.description, o)
			assert.NoError(t, err)

			assert.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expected, o.Spec.Info.Spec.Description)
		})
	}
}

func TestDescription_Tag(t *testing.T) {
	comment := spec.NewDescriptionComment()
	assert.Equal(t, "description", comment.Tag())
}

func TestDescription_Usage(t *testing.T) {
	comment := spec.NewDescriptionComment()
	assert.NotEmpty(t, comment.Usage())
}
