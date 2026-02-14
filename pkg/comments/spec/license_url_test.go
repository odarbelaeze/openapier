package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestLicenseURL_ParseInto(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "valid license url",
			url:      "https://www.apache.org/licenses/LICENSE-2.0.html",
			expected: "https://www.apache.org/licenses/LICENSE-2.0.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := spec.NewLicenseURLComment()
			o := openapi.NewOpenAPIBuilder().Build()

			err := comment.ParseInto(tt.url, o)
			assert.NoError(t, err)

			assert.NotNil(t, o.Spec.Info)
			assert.NotNil(t, o.Spec.Info.Spec.License)
			assert.Equal(t, tt.expected, o.Spec.Info.Spec.License.Spec.URL)
		})
	}
}

func TestLicenseURL_Tag(t *testing.T) {
	comment := spec.NewLicenseURLComment()
	assert.Equal(t, "license.url", comment.Tag())
}

func TestLicenseURL_Usage(t *testing.T) {
	comment := spec.NewLicenseURLComment()
	assert.NotEmpty(t, comment.Usage())
}
