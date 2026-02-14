package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestContactEmail_ParseInto(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected string
	}{
		{
			name:     "valid email",
			email:    "apiteam@swagger.io",
			expected: "apiteam@swagger.io",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := spec.NewContactEmailComment()
			o := openapi.NewOpenAPIBuilder().Build()

			err := comment.ParseInto(tt.email, o)
			assert.NoError(t, err)

			assert.NotNil(t, o.Spec.Info)
			assert.NotNil(t, o.Spec.Info.Spec.Contact)
			assert.Equal(t, tt.expected, o.Spec.Info.Spec.Contact.Spec.Email)
		})
	}
}

func TestContactEmail_Tag(t *testing.T) {
	comment := spec.NewContactEmailComment()
	assert.Equal(t, "contact.email", comment.Tag())
}

func TestContactEmail_Usage(t *testing.T) {
	comment := spec.NewContactEmailComment()
	assert.NotEmpty(t, comment.Usage())
}
