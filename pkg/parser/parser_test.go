package parser_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/parser"
	"github.com/stretchr/testify/require"
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
		})
	}
}
