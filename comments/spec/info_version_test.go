package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoVersionComment_Tag(t *testing.T) {
	comment := &spec.InfoVersionComment{}
	assert.Equal(t, "info.version", comment.Tag())
}

func TestInfoVersionComment_Usage(t *testing.T) {
	comment := &spec.InfoVersionComment{}
	assert.Equal(t, "@info.version <version>", comment.Usage())
}

func TestInfoVersionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name            string
		line            string
		expectedVersion string
	}{
		{
			name:            "valid version",
			line:            "1.0.1",
			expectedVersion: "1.0.1",
		},
		{
			name:            "empty version",
			line:            "",
			expectedVersion: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoVersionComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expectedVersion, o.Spec.Info.Spec.Version)
		})
	}
}
