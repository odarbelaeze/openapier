package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestVersion_ParseInto(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected string
	}{
		{
			name:     "valid version",
			version:  "1.0.12",
			expected: "1.0.12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := spec.NewVersionComment()
			o := openapi.NewOpenAPIBuilder().Build()

			err := comment.ParseInto(tt.version, o)
			assert.NoError(t, err)

			assert.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expected, o.Spec.Info.Spec.Version)
		})
	}
}

func TestVersion_Tag(t *testing.T) {
	comment := spec.NewVersionComment()
	assert.Equal(t, "version", comment.Tag())
}

func TestVersion_Usage(t *testing.T) {
	comment := spec.NewVersionComment()
	assert.NotEmpty(t, comment.Usage())
}
