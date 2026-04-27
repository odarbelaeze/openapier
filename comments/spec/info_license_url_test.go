package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoLicenseURLComment_Tag(t *testing.T) {
	comment := &spec.InfoLicenseURLComment{}
	assert.Equal(t, "info.license.url", comment.Tag())
}

func TestInfoLicenseURLComment_Usage(t *testing.T) {
	comment := &spec.InfoLicenseURLComment{}
	assert.Equal(t, "@info.license.url <url>", comment.Usage())
}

func TestInfoLicenseURLComment_ParseInto(t *testing.T) {
	tests := []struct {
		name               string
		line               string
		expectedLicenseURL string
	}{
		{
			name:               "valid url",
			line:               "https://www.apache.org/licenses/LICENSE-2.0.html",
			expectedLicenseURL: "https://www.apache.org/licenses/LICENSE-2.0.html",
		},
		{
			name:               "empty url",
			line:               "",
			expectedLicenseURL: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoLicenseURLComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			require.NotNil(t, o.Spec.Info.Spec.License)
			assert.Equal(t, tt.expectedLicenseURL, o.Spec.Info.Spec.License.Spec.URL)
		})
	}
}
