package comments_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestServersVariablesEnum_ParseInto(t *testing.T) {
	tests := []struct {
		name     string
		comment  string
		setup    func(*openapi.Extendable[openapi.OpenAPI])
		expected []string
		wantErr  bool
	}{
		{
			name:    "no server",
			comment: "var1 value1",
			setup:   func(o *openapi.Extendable[openapi.OpenAPI]) {},
			wantErr: true,
		},
		{
			name:    "variable not found",
			comment: "var1 value1",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				o.Spec.Servers = append(o.Spec.Servers, openapi.NewServerBuilder().Build())
			},
			wantErr: true,
		},
		{
			name:    "success",
			comment: "var1 value1",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				server := openapi.NewServerBuilder().
					AddVariable("var1", openapi.NewServerVariableBuilder().Build()).
					Build()
				o.Spec.Servers = append(o.Spec.Servers, server)
			},
			expected: []string{"value1"},
			wantErr:  false,
		},
		{
			name:    "multiple enums",
			comment: "var1 value2",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				server := openapi.NewServerBuilder().
					AddVariable("var1", openapi.NewServerVariableBuilder().Enum("value1").Build()).
					Build()
				o.Spec.Servers = append(o.Spec.Servers, server)
			},
			expected: []string{"value1", "value2"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := openapi.NewOpenAPIBuilder().Build()
			tt.setup(o)
			comment := comments.NewServersVariablesEnumComment()
			err := comment.ParseInto(tt.comment, o)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				server := o.Spec.Servers[len(o.Spec.Servers)-1]
				assert.Equal(t, tt.expected, server.Spec.Variables["var1"].Spec.Enum)
			}
		})
	}
}

func TestServersVariablesEnum_Tag(t *testing.T) {
	comment := comments.NewServersVariablesEnumComment()
	assert.Equal(t, "servers.variables.enum", comment.Tag())
}

func TestServersVariablesEnum_Usage(t *testing.T) {
	comment := comments.NewServersVariablesEnumComment()
	assert.Equal(t, "// @servers.variables.enum <variable> <value>", comment.Usage())
}
