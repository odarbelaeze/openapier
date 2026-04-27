package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoLicenseNameComment_Tag(t *testing.T) {
	comment := &spec.InfoLicenseNameComment{}
	assert.Equal(t, "info.license.name", comment.Tag())
}

func TestInfoLicenseNameComment_Usage(t *testing.T) {
	comment := &spec.InfoLicenseNameComment{}
	assert.Equal(t, "@info.license.name <name>", comment.Usage())
}

func TestInfoLicenseNameComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		line                string
		expectedLicenseName string
	}{
		{
			name:                "valid name",
			line:                "Apache 2.0",
			expectedLicenseName: "Apache 2.0",
		},
		{
			name:                "empty name",
			line:                "",
			expectedLicenseName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoLicenseNameComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			require.NotNil(t, o.Spec.Info.Spec.License)
			assert.Equal(t, tt.expectedLicenseName, o.Spec.Info.Spec.License.Spec.Name)
		})
	}
}
