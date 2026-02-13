package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestServersVariablesDescription_ParseInto(t *testing.T) {
	tests := []struct {
		name     string
		comment  string
		setup    func(*openapi.Extendable[openapi.OpenAPI])
		expected string
		wantErr  bool
	}{
		{
			name:    "no server",
			comment: "var1 some description",
			setup:   func(o *openapi.Extendable[openapi.OpenAPI]) {},
			wantErr: true,
		},
		{
			name:    "variable not found",
			comment: "var1 some description",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				server := openapi.NewServerBuilder().Build()
				server.Spec.Variables = make(map[string]*openapi.Extendable[openapi.ServerVariable])
				o.Spec.Servers = append(o.Spec.Servers, server)
			},
			wantErr: true,
		},
		{
			name:    "variables are not detected",
			comment: "var1 some description",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				server := openapi.NewServerBuilder().Build()
				o.Spec.Servers = append(o.Spec.Servers, server)
			},
			wantErr: true,
		},
		{
			name:    "regex mismatch",
			comment: "invalid-format",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				server := openapi.NewServerBuilder().
					AddVariable("var1", openapi.NewServerVariableBuilder().Build()).
					Build()
				o.Spec.Servers = append(o.Spec.Servers, server)
			},
			wantErr: false,
		},
		{
			name:    "success",
			comment: "var1 some description",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				server := openapi.NewServerBuilder().
					AddVariable("var1", openapi.NewServerVariableBuilder().Build()).
					Build()
				o.Spec.Servers = append(o.Spec.Servers, server)
			},
			expected: "some description",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := openapi.NewOpenAPIBuilder().Build()
			tt.setup(o)
			comment := spec.NewServersVariablesDescriptionComment()
			err := comment.ParseInto(tt.comment, o)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				server := o.Spec.Servers[len(o.Spec.Servers)-1]
				assert.Equal(t, tt.expected, server.Spec.Variables["var1"].Spec.Description)
			}
		})
	}
}

func TestServersVariablesDescription_Tag(t *testing.T) {
	comment := spec.NewServersVariablesDescriptionComment()
	assert.Equal(t, "servers.variables.description", comment.Tag())
}

func TestServersVariablesDescription_Usage(t *testing.T) {
	comment := spec.NewServersVariablesDescriptionComment()
	assert.Equal(t, "// @servers.variables.description <variable> <description>", comment.Usage())
}
