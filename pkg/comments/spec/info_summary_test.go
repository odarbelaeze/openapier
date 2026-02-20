package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoSummaryComment_Tag(t *testing.T) {
	comment := &spec.InfoSummaryComment{}
	assert.Equal(t, "info.summary", comment.Tag())
}

func TestInfoSummaryComment_Usage(t *testing.T) {
	comment := &spec.InfoSummaryComment{}
	assert.Equal(t, "@info.summary <summary>", comment.Usage())
}

func TestInfoSummaryComment_ParseInto(t *testing.T) {
	tests := []struct {
		name            string
		line            string
		expectedSummary string
	}{
		{
			name:            "valid summary",
			line:            "A pet store manager.",
			expectedSummary: "A pet store manager.",
		},
		{
			name:            "empty summary",
			line:            "",
			expectedSummary: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoSummaryComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expectedSummary, o.Spec.Info.Spec.Summary)
		})
	}
}
