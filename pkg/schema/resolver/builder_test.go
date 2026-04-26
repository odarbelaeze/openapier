package resolver_test

import (
	"go/ast"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/resolver"
	resolvermocks "github.com/odarbelaeze/openapier/pkg/schema/resolver/generated_mocks"
	validatormocks "github.com/odarbelaeze/openapier/pkg/schema/validator/generated_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestBuildStruct_Simple(t *testing.T) {
	registry := validatormocks.NewMockRegistry(t)
	res := resolvermocks.NewMockResolver(t)
	builder := resolver.NewSchemaBuilder(registry, res, nil)

	expr := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "Name"}},
					Type:  &ast.Ident{Name: "string"},
				},
			},
		},
	}

	res.On("Resolve", "string").Return(openapi.NewSchemaBuilder().AddType("string").Build(), nil)

	schema, err := builder.Build(expr)
	assert.NoError(t, err)
	assert.NotNil(t, schema)
	assert.Equal(t, "object", (*schema.Spec.Type)[0])
	assert.Contains(t, schema.Spec.Properties, "Name")
	assert.Contains(t, schema.Spec.Required, "Name")
}

func TestBuildStruct_Omitempty(t *testing.T) {
	registry := validatormocks.NewMockRegistry(t)
	res := resolvermocks.NewMockResolver(t)
	builder := resolver.NewSchemaBuilder(registry, res, nil)

	expr := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "Name"}},
					Type:  &ast.Ident{Name: "string"},
					Tag:   &ast.BasicLit{Value: "`json:\"name,omitempty\"`"},
				},
			},
		},
	}

	res.On("Resolve", "string").Return(openapi.NewSchemaBuilder().AddType("string").Build(), nil)

	schema, err := builder.Build(expr)
	assert.NoError(t, err)
	assert.NotNil(t, schema)
	assert.Contains(t, schema.Spec.Properties, "name")
	// Currently this will FAIL because omitempty is ignored
	assert.NotContains(t, schema.Spec.Required, "name", "Field with omitempty should not be required")
}

func TestBuildStruct_Embedded(t *testing.T) {
	registry := validatormocks.NewMockRegistry(t)
	res := resolvermocks.NewMockResolver(t)
	builder := resolver.NewSchemaBuilder(registry, res, nil)

	// type Base struct { ID string }
	// type Derived struct { Base }

	baseRef := openapi.NewRefOrSpec[openapi.Schema]("#/components/schemas/Base")

	expr := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Type: &ast.Ident{Name: "Base"},
				},
			},
		},
	}

	res.On("Resolve", "Base").Return(baseRef, nil)

	schema, err := builder.Build(expr)
	assert.NoError(t, err)
	assert.NotNil(t, schema)
	assert.Len(t, schema.Spec.AllOf, 1)
	assert.Equal(t, "#/components/schemas/Base", schema.Spec.AllOf[0].Ref.Ref)
}
