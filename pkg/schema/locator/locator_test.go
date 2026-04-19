package locator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/locator"
	"github.com/stretchr/testify/assert"
)

func TestLocator_Prefix(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "no special characters",
			path:     "github.com/odarbelaeze/openapier",
			expected: "github_com_odarbelaeze_openapier",
		},
		{
			name:     "with slashes",
			path:     "pkg/schema/locator",
			expected: "pkg_schema_locator",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := locator.Locator{Path: tt.path}
			assert.Equal(t, tt.expected, l.Prefix())
		})
	}
}

func TestLocator_Namespace(t *testing.T) {
	l := locator.Locator{Package: "schema"}
	assert.Equal(t, "schema", l.Namespace())
}

func TestLocator_TypeName(t *testing.T) {
	tests := []struct {
		name       string
		typeName   string
		typeParams []string
		expected   string
	}{
		{
			name:     "simple type",
			typeName: "Locator",
			expected: "Locator",
		},
		{
			name:       "generic type",
			typeName:   "Response",
			typeParams: []string{"T", "K"},
			expected:   "Response[T,K]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := locator.Locator{
				Name:       tt.typeName,
				TypeParams: tt.typeParams,
			}
			assert.Equal(t, tt.expected, l.TypeName())
		})
	}
}

func TestLocator_String(t *testing.T) {
	l := locator.Locator{
		Path:    "github.com/odarbelaeze/openapier/pkg/schema",
		Package: "schema",
		Name:    "Locator",
	}
	expected := "github_com_odarbelaeze_openapier_pkg_schema:schema.Locator"
	assert.Equal(t, expected, l.String())
}
