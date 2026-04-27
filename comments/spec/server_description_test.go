package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerDescriptionComment_Tag(t *testing.T) {
	comment := &spec.ServerDescriptionComment{}
	assert.Equal(t, "server.description", comment.Tag())
}

func TestServerDescriptionComment_Usage(t *testing.T) {
	comment := &spec.ServerDescriptionComment{}
	assert.Equal(t, "@server.description <description>", comment.Usage())
}

func TestServerDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		line                string
		setupServers        bool
		expectedDescription string
		expectError         bool
	}{
		{
			name:                "valid description",
			line:                "Production Server",
			setupServers:        true,
			expectedDescription: "Production Server",
			expectError:         false,
		},
		{
			name:                "empty description",
			line:                "",
			setupServers:        true,
			expectedDescription: "",
			expectError:         false,
		},
		{
			name:                "description with spaces",
			line:                "  Development Server  ",
			setupServers:        true,
			expectedDescription: "Development Server",
			expectError:         false,
		},
		{
			name:                "no preceding server",
			line:                "Production Server",
			setupServers:        false,
			expectedDescription: "",
			expectError:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.ServerDescriptionComment{}
			o := openapi.NewOpenAPIBuilder().Build()

			if tt.setupServers {
				o.Spec.Servers = []*openapi.Extendable[openapi.Server]{
					openapi.NewServerBuilder().URL("https://example.com").Build(),
				}
			}

			err := comment.ParseInto(tt.line, o)

			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "without a preceding @server.url")
			} else {
				require.NoError(t, err)
				if tt.setupServers {
					assert.Equal(t, tt.expectedDescription, o.Spec.Servers[0].Spec.Description)
				}
			}
		})
	}
}

func TestServerDescriptionComment_ParseInto_AppliesToLastServer(t *testing.T) {
	comment := &spec.ServerDescriptionComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	o.Spec.Servers = []*openapi.Extendable[openapi.Server]{
		openapi.NewServerBuilder().URL("https://dev.example.com").Build(),
		openapi.NewServerBuilder().URL("https://prod.example.com").Build(),
	}

	err := comment.ParseInto("Production API", o)
	require.NoError(t, err)

	assert.Empty(t, o.Spec.Servers[0].Spec.Description)
	assert.Equal(t, "Production API", o.Spec.Servers[1].Spec.Description)
}
