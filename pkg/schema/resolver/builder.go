package resolver

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

type SchemaBuilder interface {
	Build(expr ast.Expr, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error)
}

type SchemaBuilderFactory func(
	validatorRegistry validator.Registry,
	resolver Resolver,
	aliases map[string]string,
) SchemaBuilder

type schemaBuilder struct {
	resolver          Resolver
	validatorRegistry validator.Registry
	aliases           map[string]string
}

func NewSchemaBuilder(
	validatorRegistry validator.Registry,
	resolver Resolver,
	aliases map[string]string,
) SchemaBuilder {
	return &schemaBuilder{
		validatorRegistry: validatorRegistry,
		resolver:          resolver,
		aliases:           aliases,
	}
}

func (b *schemaBuilder) aliased(typeName string) string {
	if alias, ok := b.aliases[typeName]; ok {
		return alias
	}
	return typeName
}

func (b *schemaBuilder) Build(expr ast.Expr, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	switch ty := expr.(type) {
	case *ast.Ident:
		return b.resolver.Resolve(b.aliased(ty.Name), opts...)
	case *ast.ArrayType:
		return b.buildArray(ty, opts...)
	case *ast.StructType:
		return b.buildStruct(ty, opts...)
	case *ast.StarExpr:
		return b.Build(ty.X, opts...)
	case *ast.SelectorExpr:
		{
			typeName, err := fullSelectorName(ty)
			if err != nil {
				return nil, fmt.Errorf("failed to parse selector expression: %w", err)
			}
			return b.resolver.Resolve(typeName, opts...)
		}
	case *ast.MapType:
		return b.buildMap(ty, opts...)
	default:
		return nil, fmt.Errorf("unsupported type: %T", expr)
	}
}

func fullSelectorName(expr *ast.SelectorExpr) (string, error) {
	var parts []string
	var current ast.Expr = expr
	for {
		if sel, ok := current.(*ast.SelectorExpr); ok {
			parts = append([]string{sel.Sel.Name}, parts...)
			current = sel.X
		} else if ident, ok := current.(*ast.Ident); ok {
			parts = append([]string{ident.Name}, parts...)
			break
		} else {
			return "", fmt.Errorf("unsupported expression type in selector: %T", current)
		}
	}
	return strings.Join(parts, "."), nil
}

func (b *schemaBuilder) buildMap(ty *ast.MapType, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	builder := openapi.NewSchemaBuilder().AddType("object")
	valueSchema, err := b.Build(ty.Value)
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
	propertiesBuilder := openapi.NewSchemaBuilder().AddType("object")
	required := []string{}
	var embeds []*openapi.RefOrSpec[openapi.Schema]

	for _, field := range ty.Fields.List {
		if field.Names == nil {
			embedSchema, err := b.Build(field.Type)
			if err != nil {
				return nil, fmt.Errorf("error resolving embed schema: %w", err)
			}
			embeds = append(embeds, embedSchema)
			continue
		}
		for _, fieldName := range field.Names {
			name := fieldName.Name
			if !ast.IsExported(name) {
				continue
			}
			fieldOptions := []options.SchemaOption{}
			omitEmpty := false
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
					for _, part := range parts[1:] {
						if part == "omitempty" {
							omitEmpty = true
						}
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
			isPointer := false
			if _, ok := field.Type.(*ast.StarExpr); ok {
				isPointer = true
			}
			if !isPointer && !omitEmpty {
				required = append(required, name)
			}
			schema, err := b.Build(field.Type, fieldOptions...)
			if err != nil {
				return nil, fmt.Errorf("unsupported property type %T: %w", field.Type, err)
			}
			propertiesBuilder.AddProperty(name, schema)
		}
	}
	propertiesBuilder.Required(required...)

	var finalBuilder *openapi.SchemaBuilder
	if len(embeds) > 0 {
		finalBuilder = openapi.NewSchemaBuilder()
		for _, embed := range embeds {
			finalBuilder.AddAllOf(embed)
		}
		// Only add the properties as an object if there are properties or required fields
		if len(required) > 0 || len(ty.Fields.List) > len(embeds) {
			finalBuilder.AddAllOf(propertiesBuilder.Build())
		}
	} else {
		finalBuilder = propertiesBuilder
	}

	for _, option := range opts {
		option(finalBuilder)
	}
	return finalBuilder.Build(), nil
}

func (b *schemaBuilder) buildArray(ty *ast.ArrayType, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	builder := openapi.NewSchemaBuilder().Type("array")
	elementSchema, err := b.Build(ty.Elt)
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
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune":
			return new("integer"), nil
		case "float32", "float64":
			return new("number"), nil
		case "bool":
			return new("boolean"), nil
		case "string":
			return new("string"), nil
		case "complex64", "complex128":
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
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune":
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
