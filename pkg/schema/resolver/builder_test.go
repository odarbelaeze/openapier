package resolver_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
	"github.com/odarbelaeze/openapier/pkg/schema/resolver"
	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

// mockResolver implements the resolver.Resolver interface
type mockResolver struct {
	mock.Mock
}

func (m *mockResolver) Collect(path string, file *ast.File) {
	m.Called(path, file)
}

func (m *mockResolver) Resolve(typeName string, file *ast.File, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	args := m.Called(typeName, file, opts)
	return args.Get(0).(*openapi.RefOrSpec[openapi.Schema]), args.Error(1)
}

func (m *mockResolver) Definitions() map[string]*openapi.RefOrSpec[openapi.Schema] {
	args := m.Called()
	return args.Get(0).(map[string]*openapi.RefOrSpec[openapi.Schema])
}

// mockValidatorRegistry implements validator.Registry
type mockValidatorRegistry struct {
	mock.Mock
}

func (m *mockValidatorRegistry) Register(v ...validator.ValidatorTag) {
	m.Called(v)
}

func (m *mockValidatorRegistry) Parse(tag string, schemaType string) ([]options.SchemaOption, error) {
	args := m.Called(tag, schemaType)
	return args.Get(0).([]options.SchemaOption), args.Error(1)
}

// Helper to parse a type expression from a string
func parseExpr(t *testing.T, expr string) ast.Expr {
	f, err := parser.ParseFile(token.NewFileSet(), "", "package p; type T "+expr, 0)
	require.NoError(t, err)
	return f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type
}

func TestSchemaBuilder_Build(t *testing.T) {
	t.Run("basic identifier", func(t *testing.T) {
		mr := new(mockResolver)
		expr := parseExpr(t, "string")
		expected := openapi.NewSchemaBuilder().AddType("string").Build()

		mr.On("Resolve", "string", mock.Anything, mock.Anything).Return(expected, nil)

		builder := resolver.NewSchemaBuilder(nil, mr, nil, nil)
		got, err := builder.Build(expr)

		require.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("aliased identifier", func(t *testing.T) {
		mr := new(mockResolver)
		expr := parseExpr(t, "T")
		expected := openapi.NewSchemaBuilder().AddType("string").Build()
		aliases := map[string]string{"T": "string"}

		mr.On("Resolve", "string", mock.Anything, mock.Anything).Return(expected, nil)

		builder := resolver.NewSchemaBuilder(nil, mr, nil, aliases)
		got, err := builder.Build(expr)

		require.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("array type", func(t *testing.T) {
		mr := new(mockResolver)
		expr := parseExpr(t, "[]string")
		stringSchema := openapi.NewSchemaBuilder().AddType("string").Build()

		mr.On("Resolve", "string", mock.Anything, mock.Anything).Return(stringSchema, nil)

		builder := resolver.NewSchemaBuilder(nil, mr, nil, nil)
		got, err := builder.Build(expr)

		require.NoError(t, err)
		assert.Equal(t, "array", (*got.Spec.Type)[0])
		assert.NotNil(t, got.Spec.Items)
	})

	t.Run("map type", func(t *testing.T) {
		mr := new(mockResolver)
		expr := parseExpr(t, "map[string]int")
		intSchema := openapi.NewSchemaBuilder().AddType("integer").Build()

		mr.On("Resolve", "int", mock.Anything, mock.Anything).Return(intSchema, nil)

		builder := resolver.NewSchemaBuilder(nil, mr, nil, nil)
		got, err := builder.Build(expr)

		require.NoError(t, err)
		assert.Equal(t, "object", (*got.Spec.Type)[0])
		assert.NotNil(t, got.Spec.AdditionalProperties)
	})

	t.Run("struct type with tags", func(t *testing.T) {
		mr := new(mockResolver)
		expr := parseExpr(t, "struct { Name string `json:\"user_name\" example:\"bob\"` }")
		stringSchema := openapi.NewSchemaBuilder().AddType("string").Build()

		// The builder will call Resolve for the field type
		mr.On("Resolve", "string", mock.Anything, mock.Anything).Return(stringSchema, nil)

		builder := resolver.NewSchemaBuilder(nil, mr, nil, nil)
		got, err := builder.Build(expr)

		require.NoError(t, err)
		assert.Equal(t, "object", (*got.Spec.Type)[0])
		_, ok := got.Spec.Properties["user_name"]
		assert.True(t, ok)
		assert.GreaterOrEqual(t, len(mr.Calls[0].Arguments[2].([]options.SchemaOption)), 1)
		// We expect the example option to be passed through to the field schema
		// assert.Equal(t, "bob", prop.Spec.Examples[0])
	})

	t.Run("star expression", func(t *testing.T) {
		mr := new(mockResolver)
		expr := parseExpr(t, "*string")
		stringSchema := openapi.NewSchemaBuilder().AddType("string").Build()

		mr.On("Resolve", "string", mock.Anything, mock.Anything).Return(stringSchema, nil)

		builder := resolver.NewSchemaBuilder(nil, mr, nil, nil)
		got, err := builder.Build(expr)

		require.NoError(t, err)
		assert.Equal(t, stringSchema, got)
	})

	t.Run("selector expression", func(t *testing.T) {
		mr := new(mockResolver)
		expr := parseExpr(t, "pkg.Type")
		expected := openapi.NewSchemaBuilder().AddType("object").Build()

		mr.On("Resolve", "pkg.Type", mock.Anything, mock.Anything).Return(expected, nil)

		builder := resolver.NewSchemaBuilder(nil, mr, nil, nil)
		got, err := builder.Build(expr)

		require.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}
