package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoTermsOfServiceComment_Tag(t *testing.T) {
	comment := &spec.InfoTermsOfServiceComment{}
	assert.Equal(t, "info.termsofservice", comment.Tag())
}

func TestInfoTermsOfServiceComment_Usage(t *testing.T) {
	comment := &spec.InfoTermsOfServiceComment{}
	assert.Equal(t, "@info.termsOfService <url>", comment.Usage())
}

func TestInfoTermsOfServiceComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                   string
		line                   string
		expectedTermsOfService string
	}{
		{
			name:                   "valid url",
			line:                   "https://example.com/terms/",
			expectedTermsOfService: "https://example.com/terms/",
		},
		{
			name:                   "empty url",
			line:                   "",
			expectedTermsOfService: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoTermsOfServiceComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expectedTermsOfService, o.Spec.Info.Spec.TermsOfService)
		})
	}
}
