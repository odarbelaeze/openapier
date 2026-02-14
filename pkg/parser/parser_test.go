package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
	"go.yaml.in/yaml/v4"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name string
		root string
		main string
	}{
		{
			name: "should parse a simple spec",
			root: "./testdata/simple",
			main: "main.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := parser.NewParser()
			spec, err := p.Parse(tt.root, tt.main)
			require.NoError(t, err)
			require.NotNil(t, spec)

			expected, err := os.ReadFile(filepath.Join(tt.root, "expected.yaml"))
			require.NoError(t, err)
			var expectedSpec *openapi.Extendable[openapi.OpenAPI]
			err = yaml.Unmarshal(expected, &expectedSpec)
			require.NoError(t, err)

			assert.Equal(t, expectedSpec, spec)
		})
	}
}
