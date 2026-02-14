package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestLicenseName_ParseInto(t *testing.T) {
	tests := []struct {
		name     string
		license  string
		expected string
	}{
		{
			name:     "valid license",
			license:  "Apache 2.0",
			expected: "Apache 2.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := spec.NewLicenseNameComment()
			o := openapi.NewOpenAPIBuilder().Build()

			err := comment.ParseInto(tt.license, o)
			assert.NoError(t, err)

			assert.NotNil(t, o.Spec.Info)
			assert.NotNil(t, o.Spec.Info.Spec.License)
			assert.Equal(t, tt.expected, o.Spec.Info.Spec.License.Spec.Name)
		})
	}
}

func TestLicenseName_Tag(t *testing.T) {
	comment := spec.NewLicenseNameComment()
	assert.Equal(t, "license.name", comment.Tag())
}

func TestLicenseName_Usage(t *testing.T) {
	comment := spec.NewLicenseNameComment()
	assert.NotEmpty(t, comment.Usage())
}
