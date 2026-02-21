package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerVariableDescriptionComment_Tag(t *testing.T) {
	comment := &spec.ServerVariableDescriptionComment{}
	assert.Equal(t, "server.variable.description", comment.Tag())
}

func TestServerVariableDescriptionComment_Usage(t *testing.T) {
	comment := &spec.ServerVariableDescriptionComment{}
	assert.Equal(t, "@server.variable.description <variable> <description>", comment.Usage())
}

func TestServerVariableDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		line                string
		setupServers        bool
		expectedVar         string
		expectedDescription string
		expectError         string
	}{
		{
			name:                "valid description",
			line:                "port The port number",
			setupServers:        true,
			expectedVar:         "port",
			expectedDescription: "The port number",
		},
		{
			name:                "description with spaces",
			line:                "path   API base path  ",
			setupServers:        true,
			expectedVar:         "path",
			expectedDescription: "API base path",
		},
		{
			name:         "missing description",
			line:         "port",
			setupServers: true,
			expectError:  "invalid format",
		},
		{
			name:         "empty line",
			line:         "",
			setupServers: true,
			expectError:  "invalid format",
		},
		{
			name:         "no preceding server",
			line:         "port The port",
			setupServers: false,
			expectError:  "without a preceding @server.url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.ServerVariableDescriptionComment{}
			o := openapi.NewOpenAPIBuilder().Build()

			if tt.setupServers {
				o.Spec.Servers = []*openapi.Extendable[openapi.Server]{
					openapi.NewServerBuilder().URL("https://example.com:{port}").Build(),
				}
			}

			err := comment.ParseInto(tt.line, o)

			if tt.expectError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectError)
			} else {
				require.NoError(t, err)
				if tt.setupServers {
					vars := o.Spec.Servers[0].Spec.Variables
					require.NotNil(t, vars)
					assert.Contains(t, vars, tt.expectedVar)
					assert.Equal(t, tt.expectedDescription, vars[tt.expectedVar].Spec.Description)
				}
			}
		})
	}
}

func TestServerVariableDescriptionComment_ParseInto_PreservesExisting(t *testing.T) {
	comment := &spec.ServerVariableDescriptionComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	o.Spec.Servers = []*openapi.Extendable[openapi.Server]{
		openapi.NewServerBuilder().URL("https://example.com:{port}").Build(),
	}

	o.Spec.Servers[0].Spec.Variables = map[string]*openapi.Extendable[openapi.ServerVariable]{
		"port": openapi.NewServerVariableBuilder().Default("8080").Build(),
	}

	err := comment.ParseInto("port API Port", o)
	require.NoError(t, err)

	vars := o.Spec.Servers[0].Spec.Variables
	assert.Equal(t, "8080", vars["port"].Spec.Default)
	assert.Equal(t, "API Port", vars["port"].Spec.Description)
}
