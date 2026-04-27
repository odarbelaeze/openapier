package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerDescriptionComment_Tag(t *testing.T) {
	comment := &operation.ServerDescriptionComment{}
	assert.Equal(t, "server.description", comment.Tag())
}

func TestServerDescriptionComment_Usage(t *testing.T) {
	comment := &operation.ServerDescriptionComment{}
	assert.Equal(t, "@server.description <description>", comment.Usage())
}

func TestServerDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		content             string
		setupServers        bool
		expectedDescription string
		expectError         bool
	}{
		{
			name:                "valid description",
			content:             "Production Server",
			setupServers:        true,
			expectedDescription: "Production Server",
			expectError:         false,
		},
		{
			name:                "empty description",
			content:             "",
			setupServers:        true,
			expectedDescription: "",
			expectError:         false,
		},
		{
			name:                "description with spaces",
			content:             "  Development Server  ",
			setupServers:        true,
			expectedDescription: "Development Server",
			expectError:         false,
		},
		{
			name:                "no preceding server",
			content:             "Production Server",
			setupServers:        false,
			expectedDescription: "",
			expectError:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &operation.ServerDescriptionComment{}
			op := &operation.Operation{
				Builder: openapi.NewOperationBuilder(),
			}

			if tt.setupServers {
				op.Builder.AddServers(openapi.NewServerBuilder().URL("https://example.com").Build())
			}

			err := comment.ParseInto(tt.content, nil, op)

			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "without a preceding @server.url")
			} else {
				require.NoError(t, err)
				if tt.setupServers {
					assert.Equal(t, tt.expectedDescription, op.Builder.Build().Spec.Servers[0].Spec.Description)
				}
			}
		})
	}
}

func TestServerDescriptionComment_ParseInto_AppliesToLastServer(t *testing.T) {
	comment := &operation.ServerDescriptionComment{}
	op := &operation.Operation{
		Builder: openapi.NewOperationBuilder(),
	}

	op.Builder.AddServers(
		openapi.NewServerBuilder().URL("https://dev.example.com").Build(),
		openapi.NewServerBuilder().URL("https://prod.example.com").Build(),
	)

	err := comment.ParseInto("Production API", nil, op)
	require.NoError(t, err)

	actualServers := op.Builder.Build().Spec.Servers
	assert.Empty(t, actualServers[0].Spec.Description)
	assert.Equal(t, "Production API", actualServers[1].Spec.Description)
}
