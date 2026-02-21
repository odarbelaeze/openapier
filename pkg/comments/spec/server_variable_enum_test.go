package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestServerVariableEnumComment_Tag(t *testing.T) {
	comment := &spec.ServerVariableEnumComment{}
	assert.Equal(t, "server.variable.enum", comment.Tag())
}

func TestServerVariableEnumComment_Usage(t *testing.T) {
	comment := &spec.ServerVariableEnumComment{}
	assert.Equal(t, "@server.variable.enum <variable> [value1] [value2] ...", comment.Usage())
}

func TestServerVariableEnumComment_ParseInto(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		setupServers bool
		expectedVar  string
		expectedEnum []string
		expectError  string
	}{
		{
			name:         "single enum",
			line:         "port 8080",
			setupServers: true,
			expectedVar:  "port",
			expectedEnum: []string{"8080"},
		},
		{
			name:         "multiple enums",
			line:         "port 8080 8443",
			setupServers: true,
			expectedVar:  "port",
			expectedEnum: []string{"8080", "8443"},
		},
		{
			name:         "empty line",
			line:         "",
			setupServers: true,
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
			comment := &spec.ServerVariableEnumComment{}
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
			} else if tt.expectedVar != "" {
				require.NoError(t, err)
				if tt.setupServers {
					vars := o.Spec.Servers[0].Spec.Variables
					require.NotNil(t, vars)
					assert.Contains(t, vars, tt.expectedVar)

					if len(tt.expectedEnum) > 0 {
						assert.Equal(t, tt.expectedEnum, vars[tt.expectedVar].Spec.Enum)
					} else {
						assert.Empty(t, vars[tt.expectedVar].Spec.Enum)
					}
				}
			}
		})
	}
}

func TestServerVariableEnumComment_ParseInto_MultipleAppends(t *testing.T) {
	comment := &spec.ServerVariableEnumComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	o.Spec.Servers = []*openapi.Extendable[openapi.Server]{
		openapi.NewServerBuilder().URL("https://example.com:{port}").Build(),
	}

	err1 := comment.ParseInto("port 8080", o)
	require.NoError(t, err1)

	err2 := comment.ParseInto("port 8443", o)
	require.NoError(t, err2)

	vars := o.Spec.Servers[0].Spec.Variables
	assert.Equal(t, []string{"8080", "8443"}, vars["port"].Spec.Enum)
}
