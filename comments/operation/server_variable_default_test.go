package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerVariableDefaultComment_Tag(t *testing.T) {
	comment := &operation.ServerVariableDefaultComment{}
	assert.Equal(t, "server.variable.default", comment.Tag())
}

func TestServerVariableDefaultComment_Usage(t *testing.T) {
	comment := &operation.ServerVariableDefaultComment{}
	assert.Equal(t, "@server.variable.default <variable> <default>", comment.Usage())
}

func TestServerVariableDefaultComment_ParseInto(t *testing.T) {
	tests := []struct {
		name            string
		content         string
		setupServers    bool
		expectedVar     string
		expectedDefault string
		expectError     string
	}{
		{
			name:            "valid default",
			content:         "port 8080",
			setupServers:    true,
			expectedVar:     "port",
			expectedDefault: "8080",
		},
		{
			name:            "default with spaces",
			content:         "path /v1/api",
			setupServers:    true,
			expectedVar:     "path",
			expectedDefault: "/v1/api",
		},
		{
			name:         "missing default",
			content:      "port",
			setupServers: true,
			expectError:  "invalid format",
		},
		{
			name:         "empty content",
			content:      "",
			setupServers: true,
			expectError:  "invalid format",
		},
		{
			name:         "no preceding server",
			content:      "port 8080",
			setupServers: false,
			expectError:  "without a preceding @server.url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &operation.ServerVariableDefaultComment{}
			op := &operation.Operation{
				Builder: openapi.NewOperationBuilder(),
			}

			if tt.setupServers {
				op.Builder.AddServers(openapi.NewServerBuilder().URL("https://example.com:{port}").Build())
			}

			err := comment.ParseInto(tt.content, nil, op)

			if tt.expectError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectError)
			} else {
				require.NoError(t, err)
				if tt.setupServers {
					vars := op.Builder.Build().Spec.Servers[0].Spec.Variables
					require.NotNil(t, vars)
					assert.Contains(t, vars, tt.expectedVar)
					assert.Equal(t, tt.expectedDefault, vars[tt.expectedVar].Spec.Default)
				}
			}
		})
	}
}

func TestServerVariableDefaultComment_ParseInto_PreservesExisting(t *testing.T) {
	comment := &operation.ServerVariableDefaultComment{}
	op := &operation.Operation{
		Builder: openapi.NewOperationBuilder(),
	}

	serverBuilder := openapi.NewServerBuilder().URL("https://example.com:{port}")
	serverBuilder.AddVariable("port", openapi.NewServerVariableBuilder().Description("The port").Build())
	op.Builder.AddServers(serverBuilder.Build())

	err := comment.ParseInto("port 8443", nil, op)
	require.NoError(t, err)

	vars := op.Builder.Build().Spec.Servers[0].Spec.Variables
	assert.Equal(t, "8443", vars["port"].Spec.Default)
	assert.Equal(t, "The port", vars["port"].Spec.Description)
}
