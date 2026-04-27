package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerVariableDescriptionComment_Tag(t *testing.T) {
	comment := &operation.ServerVariableDescriptionComment{}
	assert.Equal(t, "server.variable.description", comment.Tag())
}

func TestServerVariableDescriptionComment_Usage(t *testing.T) {
	comment := &operation.ServerVariableDescriptionComment{}
	assert.Equal(t, "@server.variable.description <variable> <description>", comment.Usage())
}

func TestServerVariableDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		content             string
		setupServers        bool
		expectedVar         string
		expectedDescription string
		expectError         string
	}{
		{
			name:                "valid description",
			content:             "port The port number",
			setupServers:        true,
			expectedVar:         "port",
			expectedDescription: "The port number",
		},
		{
			name:                "description with spaces",
			content:             "path   API base path  ",
			setupServers:        true,
			expectedVar:         "path",
			expectedDescription: "API base path",
		},
		{
			name:         "missing description",
			content:      "port",
			setupServers: true,
			expectError:  "invalid format",
		},
		{
			name:         "empty line",
			content:      "",
			setupServers: true,
			expectError:  "invalid format",
		},
		{
			name:         "no preceding server",
			content:      "port The port",
			setupServers: false,
			expectError:  "without a preceding @server.url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &operation.ServerVariableDescriptionComment{}
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
					assert.Equal(t, tt.expectedDescription, vars[tt.expectedVar].Spec.Description)
				}
			}
		})
	}
}

func TestServerVariableDescriptionComment_ParseInto_PreservesExisting(t *testing.T) {
	comment := &operation.ServerVariableDescriptionComment{}
	op := &operation.Operation{
		Builder: openapi.NewOperationBuilder(),
	}

	serverBuilder := openapi.NewServerBuilder().URL("https://example.com:{port}")
	serverBuilder.AddVariable("port", openapi.NewServerVariableBuilder().Default("8080").Build())
	op.Builder.AddServers(serverBuilder.Build())

	err := comment.ParseInto("port API Port", nil, op)
	require.NoError(t, err)

	vars := op.Builder.Build().Spec.Servers[0].Spec.Variables
	assert.Equal(t, "8080", vars["port"].Spec.Default)
	assert.Equal(t, "API Port", vars["port"].Spec.Description)
}
