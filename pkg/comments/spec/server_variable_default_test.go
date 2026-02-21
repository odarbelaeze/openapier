package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerVariableDefaultComment_Tag(t *testing.T) {
	comment := &spec.ServerVariableDefaultComment{}
	assert.Equal(t, "server.variable.default", comment.Tag())
}

func TestServerVariableDefaultComment_Usage(t *testing.T) {
	comment := &spec.ServerVariableDefaultComment{}
	assert.Equal(t, "@server.variable.default <variable> <default>", comment.Usage())
}

func TestServerVariableDefaultComment_ParseInto(t *testing.T) {
	tests := []struct {
		name            string
		line            string
		setupServers    bool
		expectedVar     string
		expectedDefault string
		expectError     string
	}{
		{
			name:            "valid default",
			line:            "port 8080",
			setupServers:    true,
			expectedVar:     "port",
			expectedDefault: "8080",
		},
		{
			name:            "default with spaces",
			line:            "path /v1/api",
			setupServers:    true,
			expectedVar:     "path",
			expectedDefault: "/v1/api",
		},
		{
			name:         "missing default",
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
			line:         "port 8080",
			setupServers: false,
			expectError:  "without a preceding @server.url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.ServerVariableDefaultComment{}
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
					assert.Equal(t, tt.expectedDefault, vars[tt.expectedVar].Spec.Default)
				}
			}
		})
	}
}

func TestServerVariableDefaultComment_ParseInto_PreservesExisting(t *testing.T) {
	comment := &spec.ServerVariableDefaultComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	o.Spec.Servers = []*openapi.Extendable[openapi.Server]{
		openapi.NewServerBuilder().URL("https://example.com:{port}").Build(),
	}

	o.Spec.Servers[0].Spec.Variables = map[string]*openapi.Extendable[openapi.ServerVariable]{
		"port": openapi.NewServerVariableBuilder().Description("The port").Build(),
	}

	err := comment.ParseInto("port 8443", o)
	require.NoError(t, err)

	vars := o.Spec.Servers[0].Spec.Variables
	assert.Equal(t, "8443", vars["port"].Spec.Default)
	assert.Equal(t, "The port", vars["port"].Spec.Description)
}
