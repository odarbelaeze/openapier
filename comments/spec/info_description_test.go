package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoDescriptionComment_Tag(t *testing.T) {
	comment := &spec.InfoDescriptionComment{}
	assert.Equal(t, "info.description", comment.Tag())
}

func TestInfoDescriptionComment_Usage(t *testing.T) {
	comment := &spec.InfoDescriptionComment{}
	assert.Equal(t, "@info.description <description>", comment.Usage())
}

func TestInfoDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		line                string
		expectedDescription string
	}{
		{
			name:                "valid description",
			line:                "This is a sample server for a pet store.",
			expectedDescription: "This is a sample server for a pet store.",
		},
		{
			name:                "empty description",
			line:                "",
			expectedDescription: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoDescriptionComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expectedDescription, o.Spec.Info.Spec.Description)
		})
	}
}
