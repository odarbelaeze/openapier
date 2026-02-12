package comments_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestServersURL_ParseInto(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
		vars     []string
	}{
		{
			name:     "simple url",
			url:      "https://example.com/api",
			expected: "https://example.com/api",
			vars:     nil,
		},
		{
			name:     "url with variables",
			url:      "https://{subdomain}.example.com:{port}/v1",
			expected: "https://{subdomain}.example.com:{port}/v1",
			vars:     []string{"subdomain", "port"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := comments.NewServersURLComment()
			o := openapi.NewOpenAPIBuilder().Build()

			err := comment.ParseInto(tt.url, o)
			assert.NoError(t, err)

			// We need to check if the server was added to o.Spec.Servers
			if len(o.Spec.Servers) > 0 {
				server := o.Spec.Servers[len(o.Spec.Servers)-1]
				assert.Equal(t, tt.expected, server.Spec.URL)

				if len(tt.vars) > 0 {
					assert.NotNil(t, server.Spec.Variables)
					for _, v := range tt.vars {
						_, ok := server.Spec.Variables[v]
						assert.True(t, ok, "variable %s not found", v)
					}
				}
			} else {
				// If we expected changes but got none, fail (unless we expect failure)
				assert.Fail(t, "Server was not added to OpenAPI spec")
			}
		})
	}
}

func TestServersURL_Tag(t *testing.T) {
	comment := comments.NewServersURLComment()
	assert.Equal(t, "servers.url", comment.Tag())
}

func TestServersURL_Usage(t *testing.T) {
	comment := comments.NewServersURLComment()
	assert.NotEmpty(t, comment.Usage())
}
