package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestTermsOfService_ParseInto(t *testing.T) {
	tests := []struct {
		name           string
		termsOfService string
		expected       string
	}{
		{
			name:           "simple url",
			termsOfService: "https://example.com/terms",
			expected:       "https://example.com/terms",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := spec.NewTermsOfServiceComment()
			o := openapi.NewOpenAPIBuilder().Build()

			err := comment.ParseInto(tt.termsOfService, o)
			assert.NoError(t, err)

			assert.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expected, o.Spec.Info.Spec.TermsOfService)
		})
	}
}

func TestTermsOfService_Tag(t *testing.T) {
	comment := spec.NewTermsOfServiceComment()
	assert.Equal(t, "termsofservice", comment.Tag())
}

func TestTermsOfService_Usage(t *testing.T) {
	comment := spec.NewTermsOfServiceComment()
	assert.NotEmpty(t, comment.Usage())
}
