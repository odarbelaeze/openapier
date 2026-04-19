package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoLicenseIdentifierComment_Tag(t *testing.T) {
	comment := &spec.InfoLicenseIdentifierComment{}
	assert.Equal(t, "info.license.identifier", comment.Tag())
}

func TestInfoLicenseIdentifierComment_Usage(t *testing.T) {
	comment := &spec.InfoLicenseIdentifierComment{}
	assert.Equal(t, "@info.license.identifier <identifier>", comment.Usage())
}

func TestInfoLicenseIdentifierComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                      string
		line                      string
		expectedLicenseIdentifier string
	}{
		{
			name:                      "valid identifier",
			line:                      "Apache-2.0",
			expectedLicenseIdentifier: "Apache-2.0",
		},
		{
			name:                      "empty identifier",
			line:                      "",
			expectedLicenseIdentifier: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoLicenseIdentifierComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			require.NotNil(t, o.Spec.Info.Spec.License)
			assert.Equal(t, tt.expectedLicenseIdentifier, o.Spec.Info.Spec.License.Spec.Identifier)
		})
	}
}
