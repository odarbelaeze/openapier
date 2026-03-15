package schema

import (
	"reflect"
	"testing"

	"github.com/sv-tools/openapi"
)

func TestParseType(t *testing.T) {
	tests := []struct {
		name     string
		typeStr  string
		expected *openapi.RefOrSpec[openapi.Schema]
	}{
		{
			name:     "int type",
			typeStr:  "int",
			expected: openapi.NewSchemaBuilder().AddType("integer").Format("int32").Build(),
		},
		{
			name:     "int32 type",
			typeStr:  "int32",
			expected: openapi.NewSchemaBuilder().AddType("integer").Format("int32").Build(),
		},
		{
			name:     "int64 type",
			typeStr:  "int64",
			expected: openapi.NewSchemaBuilder().AddType("integer").Format("int64").Build(),
		},
		{
			name:     "uint type",
			typeStr:  "uint",
			expected: openapi.NewSchemaBuilder().AddType("integer").Format("int32").Build(),
		},
		{
			name:     "uint32 type",
			typeStr:  "uint32",
			expected: openapi.NewSchemaBuilder().AddType("integer").Format("int32").Build(),
		},
		{
			name:     "uint64 type",
			typeStr:  "uint64",
			expected: openapi.NewSchemaBuilder().AddType("integer").Format("int64").Build(),
		},
		{
			name:     "float32 type",
			typeStr:  "float32",
			expected: openapi.NewSchemaBuilder().AddType("number").Format("float").Build(),
		},
		{
			name:     "float64 type",
			typeStr:  "float64",
			expected: openapi.NewSchemaBuilder().AddType("number").Format("double").Build(),
		},
		{
			name:     "bool type",
			typeStr:  "bool",
			expected: openapi.NewSchemaBuilder().AddType("boolean").Build(),
		},
		{
			name:     "string type",
			typeStr:  "string",
			expected: openapi.NewSchemaBuilder().AddType("string").Build(),
		},
		{
			name:     "file type",
			typeStr:  "file",
			expected: openapi.NewSchemaBuilder().AddType("string").Format("binary").Build(),
		},
		{
			name:     "any type",
			typeStr:  "any",
			expected: openapi.NewSchemaBuilder().Build(),
		},
		{
			name:     "unknown type",
			typeStr:  "someUnknownType",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseBasicType(tt.typeStr)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseType(%q) = %+v; want %+v", tt.typeStr, got, tt.expected)
			}
		})
	}
}
