package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/odarbelaeze/openapier/pkg/parser"
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
		{
			name: "should parse types in a spec",
			root: "./testdata/types",
			main: "main.go",
		},
		{
			name: "should parse generics",
			root: "./testdata/generics",
			main: "main.go",
		},
		{
			name: "should parse validation tags",
			root: "./testdata/validator",
			main: "main.go",
		},
		{
			name: "should parse external dependencies",
			root: "./testdata/external",
			main: "main.go",
		},
		{
			name: "should parse embed structs",
			root: "./testdata/embed",
			main: "main.go",
		},
		{
			name: "should support many a format",
			root: "./testdata/formats",
			main: "main.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := parser.NewParser(tt.root, tt.main)
			spec, err := p.Parse()
			require.NoError(t, err)
			require.NotNil(t, spec)

			expectedString, err := os.ReadFile(filepath.Join(tt.root, "expected.yaml"))
			require.NoError(t, err)

			expectedSpec := &openapi.Extendable[openapi.OpenAPI]{}
			err = yaml.Unmarshal(expectedString, expectedSpec)
			require.NoError(t, err)

			if diff := cmp.Diff(expectedSpec, spec, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("spec mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
