package schema

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/sv-tools/openapi"
)

type schemaBuilder struct {
	resolver Resolver
	file     *ast.File
}

func (b *schemaBuilder) build(expr ast.Expr) (*openapi.RefOrSpec[openapi.Schema], error) {
	switch ty := expr.(type) {
	case *ast.Ident:
		return b.resolver.Resolve(ty.Name, b.file)
	case *ast.ArrayType:
		return b.buildArray(ty)
	case *ast.StructType:
		return b.buildStruct(ty)
	case *ast.StarExpr:
		return b.build(ty.X)
	default:
		return nil, fmt.Errorf("unsupported type: %T", expr)
	}
}

func (b *schemaBuilder) buildStruct(ty *ast.StructType) (*openapi.RefOrSpec[openapi.Schema], error) {
	builder := openapi.NewSchemaBuilder().AddType("object")
	for _, field := range ty.Fields.List {
		for _, fieldName := range field.Names {
			name := fieldName.Name
			if !ast.IsExported(name) {
				continue
			}
			if field.Tag != nil {
				tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
				jsonTag := tag.Get("json")
				if jsonTag == "-" {
					continue
				}
				if jsonTag != "" {
					name = jsonTag
				}
			}
			schema, err := b.build(field.Type)
			if err != nil {
				return nil, fmt.Errorf("unsupported property type %T: %w", field.Type, err)
			}
			builder.AddProperty(name, schema)
		}
	}
	return builder.Build(), nil
}

func (b *schemaBuilder) buildArray(ty *ast.ArrayType) (*openapi.RefOrSpec[openapi.Schema], error) {
	builder := openapi.NewSchemaBuilder().Type("array")
	elementSchema, err := b.build(ty.Elt)
	if err != nil {
		return nil, fmt.Errorf("unsupported element type %T: %w", ty.Elt, err)
	}
	itemSchema := openapi.NewBoolOrSchema(elementSchema)
	itemSchema.Allowed = true
	builder.Items(itemSchema)
	return builder.Build(), nil
}
