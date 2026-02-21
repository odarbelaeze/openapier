package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerURLComment_Tag(t *testing.T) {
	comment := &operation.ServerURLComment{}
	assert.Equal(t, "server.url", comment.Tag())
}

func TestServerURLComment_Usage(t *testing.T) {
	comment := &operation.ServerURLComment{}
	assert.Equal(t, "@server.url <url>", comment.Usage())
}

func TestServerURLComment_ParseInto(t *testing.T) {
	tests := []struct {
		name         string
		content      string
		expectedURLs []string
	}{
		{
			name:         "valid url",
			content:      "https://example.com/v1",
			expectedURLs: []string{"https://example.com/v1"},
		},
		{
			name:         "empty line",
			content:      "",
			expectedURLs: nil, // Should not append
		},
		{
			name:         "url with spaces",
			content:      "  https://example.com/v2  ",
			expectedURLs: []string{"https://example.com/v2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &operation.ServerURLComment{}
			op := &operation.Operation{
				Builder: openapi.NewOperationBuilder(),
			}

			err := comment.ParseInto(tt.content, op)
			require.NoError(t, err)

			actualServers := op.Builder.Build().Spec.Servers
			if len(tt.expectedURLs) == 0 {
				assert.Empty(t, actualServers)
			} else {
				require.Len(t, actualServers, len(tt.expectedURLs))
				for i, expectedURL := range tt.expectedURLs {
					assert.Equal(t, expectedURL, actualServers[i].Spec.URL)
				}
			}
		})
	}
}

func TestServerURLComment_ParseInto_Multiple(t *testing.T) {
	comment := &operation.ServerURLComment{}
	op := &operation.Operation{
		Builder: openapi.NewOperationBuilder(),
	}

	err1 := comment.ParseInto("https://dev.example.com", op)
	require.NoError(t, err1)

	err2 := comment.ParseInto("https://prod.example.com", op)
	require.NoError(t, err2)

	actualServers := op.Builder.Build().Spec.Servers
	require.Len(t, actualServers, 2)
	assert.Equal(t, "https://dev.example.com", actualServers[0].Spec.URL)
	assert.Equal(t, "https://prod.example.com", actualServers[1].Spec.URL)
}
