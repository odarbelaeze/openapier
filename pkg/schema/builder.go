package schema

import (
	"fmt"
	"go/ast"
	"reflect"
	"strconv"
	"strings"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	"github.com/sv-tools/openapi"
)

const (
	jsonStructTag     = "json"
	exampleStructTag  = "example"
	validateStructTag = "validate"
)

type schemaBuilder struct {
	resolver          Resolver
	validatorRegistry validator.Registry
	file              *ast.File
	aliases           map[string]string
}

func (b *schemaBuilder) aliased(typeName string) string {
	if alias, ok := b.aliases[typeName]; ok {
		return alias
	}
	return typeName
}

func (b *schemaBuilder) build(expr ast.Expr, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	switch ty := expr.(type) {
	case *ast.Ident:
		return b.resolver.Resolve(b.aliased(ty.Name), b.file, opts...)
	case *ast.ArrayType:
		return b.buildArray(ty, opts...)
	case *ast.StructType:
		return b.buildStruct(ty, opts...)
	case *ast.StarExpr:
		return b.build(ty.X, opts...)
	case *ast.SelectorExpr:
		{
			pkgIdent, ok := ty.X.(*ast.Ident)
			if !ok {
				return nil, fmt.Errorf("unsupported package identifier in selector expression: %T", ty.X)
			}
			pkgName := pkgIdent.Name
			typeName := ty.Sel.Name
			return b.resolver.Resolve(fmt.Sprintf("%s.%s", pkgName, typeName), b.file, opts...)
		}
	case *ast.MapType:
		return b.buildMap(ty, opts...)
	default:
		return nil, fmt.Errorf("unsupported type: %T", expr)
	}
}

func (b *schemaBuilder) buildMap(ty *ast.MapType, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	builder := openapi.NewSchemaBuilder().AddType("object")
	valueSchema, err := b.build(ty.Value)
	if err != nil {
		return nil, err
	}
	valueBoolOrSchema := openapi.NewBoolOrSchema(valueSchema)
	valueBoolOrSchema.Allowed = true
	builder.AdditionalProperties(valueBoolOrSchema)
	for _, opt := range opts {
		opt(builder)
	}
	return builder.Build(), nil
}

func (b *schemaBuilder) buildStruct(ty *ast.StructType, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	builder := openapi.NewSchemaBuilder().AddType("object")
	required := []string{}
	for _, field := range ty.Fields.List {
		for _, fieldName := range field.Names {
			name := fieldName.Name
			if !ast.IsExported(name) {
				continue
			}
			fieldOptions := []options.SchemaOption{}
			if field.Tag != nil {
				tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
				jsonTag := tag.Get(jsonStructTag)
				if jsonTag == "-" {
					continue
				}
				if jsonTag != "" {
					parts := strings.Split(jsonTag, ",")
					if parts[0] != "" {
						name = parts[0]
					}
				}
				example := tag.Get(exampleStructTag)
				if example != "" {
					fieldOptions = append(fieldOptions, options.WithExample(parseExampleValue(example, field.Type)))
				}

				validate := tag.Get(validateStructTag)
				if validate != "" {
					schemaType, err := parseSchemaType(field.Type)
					if err != nil {
						return nil, fmt.Errorf("failed to parse schema type for field %s: %w", name, err)
					}
					validateOpts, err := b.validatorRegistry.Parse(validate, *schemaType)
					if err != nil {
						return nil, fmt.Errorf("failed to parse validate tag for field %s: %w", name, err)
					}
					fieldOptions = append(fieldOptions, validateOpts...)
				}
			}
			if field.Doc != nil {
				fieldOptions = append(fieldOptions, options.WithDescription(field.Doc.Text()))
			}
			if _, ok := field.Type.(*ast.StarExpr); !ok {
				required = append(required, name)
			}
			schema, err := b.build(field.Type, fieldOptions...)
			if err != nil {
				return nil, fmt.Errorf("unsupported property type %T: %w", field.Type, err)
			}
			builder.AddProperty(name, schema)
		}
	}
	builder.Required(required...)
	for _, option := range opts {
		option(builder)
	}
	return builder.Build(), nil
}

func (b *schemaBuilder) buildArray(ty *ast.ArrayType, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	builder := openapi.NewSchemaBuilder().Type("array")
	elementSchema, err := b.build(ty.Elt)
	if err != nil {
		return nil, fmt.Errorf("unsupported element type %T: %w", ty.Elt, err)
	}
	itemSchema := openapi.NewBoolOrSchema(elementSchema)
	itemSchema.Allowed = true
	builder.Items(itemSchema)
	if ty.Len != nil {
		if basicLit, ok := ty.Len.(*ast.BasicLit); ok {
			if val, err := strconv.Atoi(basicLit.Value); err == nil {
				builder.MinItems(val)
				builder.MaxItems(val)
			}
		}
	}
	for _, option := range opts {
		option(builder)
	}
	return builder.Build(), nil
}

func parseSchemaType(ty ast.Expr) (*string, error) {
	switch expr := ty.(type) {
	case *ast.Ident:
		switch expr.Name {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return new("integer"), nil
		case "float32", "float64":
			return new("number"), nil
		case "bool":
			return new("boolean"), nil
		case "string":
			return new("string"), nil
		default:
			return nil, fmt.Errorf("unsupported type %s", expr.Name)
		}
	case *ast.StarExpr:
		return parseSchemaType(expr.X)
	case *ast.ArrayType:
		return new("array"), nil
	case *ast.MapType:
		return new("object"), nil
	}
	return nil, fmt.Errorf("unsupported type %T", ty)
}

func parseExampleValue(example string, ty ast.Expr) any {
	switch expr := ty.(type) {
	case *ast.Ident:
		switch expr.Name {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			if val, err := strconv.ParseInt(example, 10, 64); err == nil {
				return int(val)
			}
		case "float32", "float64":
			if val, err := strconv.ParseFloat(example, 64); err == nil {
				return val
			}
		case "bool":
			if val, err := strconv.ParseBool(example); err == nil {
				return val
			}
		}
	case *ast.StarExpr:
		return parseExampleValue(example, expr.X)
	}
	return example
}
