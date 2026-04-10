package resolver_test

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/resolver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestResolver_Resolve_BasicType_Success(t *testing.T) {
	r := resolver.NewResolver(nil, resolver.NewSchemaBuilder)

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Resolve(tt.typeStr, nil)
			require.NoError(t, err)
			assert.Equal(t, got, tt.expected)
		})
	}
}

func TestResolver_CollectAndResolve(t *testing.T) {
	r := resolver.NewResolver(nil, resolver.NewSchemaBuilder)
	fset := token.NewFileSet()
	src := `package test
type User struct {
	Name string
}`
	f, err := parser.ParseFile(fset, "test.go", src, 0)
	require.NoError(t, err)

	r.Collect("github.com/test", f)

	// Resolve the type
	got, err := r.Resolve("User", f)
	require.NoError(t, err)

	assert.NotNil(t, got.Ref)
	assert.Equal(t, "#/components/schemas/github_com_test:test.User", got.Ref.Ref)

	// Verify it's in definitions
	defs := r.Definitions()
	assert.Contains(t, defs, "github_com_test:test.User")
}

func TestResolver_ResolveComplexTypes(t *testing.T) {
	r := resolver.NewResolver(nil, resolver.NewSchemaBuilder)
	fset := token.NewFileSet()
	src := `package test
type User struct {
	Name string
}`
	f, err := parser.ParseFile(fset, "test.go", src, 0)
	require.NoError(t, err)

	r.Collect("github.com/test", f)

	t.Run("resolve map", func(t *testing.T) {
		got, err := r.Resolve("map[string]User", f)
		require.NoError(t, err)
		assert.NotNil(t, got.Spec)
		assert.Equal(t, "object", (*got.Spec.Type)[0])
		assert.NotNil(t, got.Spec.AdditionalProperties)
	})

	t.Run("resolve array", func(t *testing.T) {
		got, err := r.Resolve("[]User", f)
		require.NoError(t, err)
		assert.NotNil(t, got.Spec)
		assert.Equal(t, "array", (*got.Spec.Type)[0])
		assert.NotNil(t, got.Spec.Items)
	})

	t.Run("resolve fixed array", func(t *testing.T) {
		got, err := r.Resolve("[5]User", f)
		require.NoError(t, err)
		assert.NotNil(t, got.Spec)
		assert.Equal(t, "array", (*got.Spec.Type)[0])
		assert.Equal(t, 5, *got.Spec.MinItems)
		assert.Equal(t, 5, *got.Spec.MaxItems)
	})
}


