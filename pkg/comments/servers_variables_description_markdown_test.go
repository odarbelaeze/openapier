package comments_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestServersVariablesDescriptionMarkdown_ParseInto(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "openapier-test-*")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	markdownFile := "test.md"
	markdownContent := "This is a markdown description."
	err = os.WriteFile(filepath.Join(tempDir, markdownFile), []byte(markdownContent), 0644)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		comment  string
		setup    func(*openapi.Extendable[openapi.OpenAPI])
		expected string
		wantErr  bool
	}{
		{
			name:    "no server",
			comment: "var1 test.md",
			setup:   func(o *openapi.Extendable[openapi.OpenAPI]) {},
			wantErr: true,
		},
		{
			name:    "variable not found",
			comment: "var1 test.md",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				o.Spec.Servers = append(o.Spec.Servers, openapi.NewServerBuilder().Build())
			},
			wantErr: true,
		},
		{
			name:    "success",
			comment: "var1 test.md",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				server := openapi.NewServerBuilder().
					AddVariable("var1", openapi.NewServerVariableBuilder().Build()).
					Build()
				o.Spec.Servers = append(o.Spec.Servers, server)
			},
			expected: markdownContent,
			wantErr:  false,
		},
		{
			name:    "file not found",
			comment: "var1 non-existent.md",
			setup: func(o *openapi.Extendable[openapi.OpenAPI]) {
				server := openapi.NewServerBuilder().
					AddVariable("var1", openapi.NewServerVariableBuilder().Build()).
					Build()
				o.Spec.Servers = append(o.Spec.Servers, server)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := openapi.NewOpenAPIBuilder().Build()
			tt.setup(o)
			comment := comments.NewServersVariablesDescriptionMarkdownComment(tempDir)
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

func TestServersVariablesDescriptionMarkdown_Tag(t *testing.T) {
	comment := comments.NewServersVariablesDescriptionMarkdownComment(".")
	assert.Equal(t, "servers.variables.description.markdown", comment.Tag())
}

func TestServersVariablesDescriptionMarkdown_Usage(t *testing.T) {
	comment := comments.NewServersVariablesDescriptionMarkdownComment(".")
	assert.Equal(t, "// @servers.variables.description.markdown <variable> <filename>", comment.Usage())
}
