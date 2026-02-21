package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerURLComment_Tag(t *testing.T) {
	comment := &spec.ServerURLComment{}
	assert.Equal(t, "server.url", comment.Tag())
}

func TestServerURLComment_Usage(t *testing.T) {
	comment := &spec.ServerURLComment{}
	assert.Equal(t, "@server.url <url>", comment.Usage())
}

func TestServerURLComment_ParseInto(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		expectedURLs []string
	}{
		{
			name:         "valid url",
			line:         "https://example.com/v1",
			expectedURLs: []string{"https://example.com/v1"},
		},
		{
			name:         "empty line",
			line:         "",
			expectedURLs: nil, // Should not append
		},
		{
			name:         "url with spaces",
			line:         "  https://example.com/v2  ",
			expectedURLs: []string{"https://example.com/v2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.ServerURLComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)

			if len(tt.expectedURLs) == 0 {
				assert.Empty(t, o.Spec.Servers)
			} else {
				require.Len(t, o.Spec.Servers, len(tt.expectedURLs))
				for i, expectedURL := range tt.expectedURLs {
					assert.Equal(t, expectedURL, o.Spec.Servers[i].Spec.URL)
				}
			}
		})
	}
}

func TestServerURLComment_ParseInto_Multiple(t *testing.T) {
	comment := &spec.ServerURLComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	err1 := comment.ParseInto("https://dev.example.com", o)
	require.NoError(t, err1)

	err2 := comment.ParseInto("https://prod.example.com", o)
	require.NoError(t, err2)

	require.Len(t, o.Spec.Servers, 2)
	assert.Equal(t, "https://dev.example.com", o.Spec.Servers[0].Spec.URL)
	assert.Equal(t, "https://prod.example.com", o.Spec.Servers[1].Spec.URL)
}
