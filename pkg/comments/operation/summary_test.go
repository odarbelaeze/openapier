package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSummaryComment_ParseInto(t *testing.T) {
	tests := []struct {
		name            string
		comment         string
		expectedSummary string
		expectedError   string
	}{
		{
			name:            "valid summary",
			comment:         "My summary",
			expectedSummary: "My summary",
		},
		{
			name:            "empty summary",
			comment:         "",
			expectedSummary: "",
		},
		{
			name:            "summary with spaces",
			comment:         "  My summary  ",
			expectedSummary: "My summary",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := operation.NewOperation()
			comment := operation.NewSummaryComment()

			err := comment.ParseInto(tt.comment, op)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				spec := op.Builder.Build()
				assert.Equal(t, tt.expectedSummary, spec.Spec.Summary)
			}
		})
	}
}

func TestSummaryComment_Tag(t *testing.T) {
	comment := operation.NewSummaryComment()
	assert.Equal(t, "summary", comment.Tag())
}

func TestSummaryComment_Usage(t *testing.T) {
	comment := operation.NewSummaryComment()
	assert.Equal(t, "@summary <summary>", comment.Usage())
}

func TestSummaryComment_Integration(t *testing.T) {
	op := operation.NewOperation()
	err := operation.DefaultRegistry.Parse("// @summary Integrated Summary", op)
	require.NoError(t, err)

	spec := op.Builder.Build()
	assert.Equal(t, "Integrated Summary", spec.Spec.Summary)
}
