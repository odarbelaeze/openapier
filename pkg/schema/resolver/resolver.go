package resolver

import (
	"context"
	"fmt"
	"go/ast"
	"log/slog"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/odarbelaeze/openapier/pkg/cache"
	"github.com/odarbelaeze/openapier/pkg/schema/locator"
	"github.com/odarbelaeze/openapier/pkg/schema/options"
	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	"github.com/sv-tools/openapi"
)

var (
	genericTypeRegex = regexp.MustCompile(`^(((\w+)\.)?(\w+))\[(.+)\]$`)
	arrayTypeRegex   = regexp.MustCompile(`^\[(\d*)\](.*)$`)
	mapTypeRegex     = regexp.MustCompile(`^map\[(.+)\](.*)$`)
)

// Resolves types into a schema.
type Resolver interface {
	// Resolve resolves the given type into a schema.
	Resolve(typeName string, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error)

	// From returns a new resolver re-positioned to the given file.
	From(file *ast.File, from string) Resolver
}

type resolver struct {
	// validatorRegistry is the registry of validators used to validate schemas.
	validatorRegistry validator.Registry

	// builderFactory is the factory used to create schema builders.
	builderFactory SchemaBuilderFactory

	// typeDefCache is the cache of type definitions.
	typeDefCache TypeDefCache

	// definitionsCache is the cache of resolved definitions.
	definitionsCache cache.DefinitionsCache

	// file is the AST file being resolved.
	file *ast.File

	// from is the package path of the file being resolved.
	from string
}

// NewResolver creates a new resolver.
func NewResolver(
	validatorRegistry validator.Registry,
	typeDefCache TypeDefCache,
	definitionsCache cache.DefinitionsCache,
	builderFactory SchemaBuilderFactory,
	file *ast.File,
	from string,
) Resolver {
	return &resolver{
		validatorRegistry: validatorRegistry,
		builderFactory:    builderFactory,
		typeDefCache:      typeDefCache,
		definitionsCache:  definitionsCache,
		file:              file,
		from:              from,
	}
}

func (r *resolver) From(file *ast.File, from string) Resolver {
	return &resolver{
		validatorRegistry: r.validatorRegistry,
		builderFactory:    r.builderFactory,
		typeDefCache:      r.typeDefCache,
		definitionsCache:  r.definitionsCache,
		file:              file,
		from:              from,
	}
}

// Resolve implements [Resolver].
func (r *resolver) Resolve(typeName string, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	if strings.HasPrefix(typeName, "[") {
		return r.resolveArray(typeName, opts...)
	}

	if strings.HasPrefix(typeName, "map[") {
		return r.resolveMap(typeName, opts...)
	}

	basicSchema := r.resolveBasicType(typeName, opts...)
	if basicSchema != nil {
		return basicSchema, nil
	}

	matches := genericTypeRegex.FindStringSubmatch(typeName)
	var typeParams []string
	if len(matches) == 6 {
		typeName = matches[1]
		typeParams = strings.Split(matches[5], ",")
		slog.Debug("resolving generic type", "typeName", typeName, "typeArgs", typeParams)
	}

	candidates, err := r.candidates(typeName, typeParams)
	if err != nil {
		return nil, fmt.Errorf("failed to find any candidates: %w", err)
	}
	for _, loc := range candidates {
		schemaPath := fmt.Sprintf("#/components/schemas/%s", loc)
		if _, ok := r.definitionsCache.Get(loc.String()); ok {
			ref := openapi.NewRefOrSpec[openapi.Schema](schemaPath)
			return ref, nil
		}
		if t, ok := r.typeDefCache.Get(loc.Path, loc.Name); ok {
			if len(loc.TypeParams) != t.TypeSpec.TypeParams.NumFields() {
				return nil, fmt.Errorf(
					"type parameter count mismatch for %s: expected %d, got %d",
					loc.TypeName(),
					t.TypeSpec.TypeParams.NumFields(),
					len(loc.TypeParams),
				)
			}
			aliases := make(map[string]string)
			if t.TypeSpec.TypeParams != nil {
				for i, field := range t.TypeSpec.TypeParams.List {
					paramName := field.Names[0].Name
					if i < len(loc.TypeParams) {
						aliases[paramName] = loc.TypeParams[i]
					}
				}
			}
			slog.Debug("finding spec for", "typeName", t.TypeSpec.Name.Name)
			b := r.builderFactory(r.validatorRegistry, r.From(t.File, t.Locator.Path), aliases)
			spec, err := b.Build(t.TypeSpec.Type, opts...)
			if err != nil {
				return nil, fmt.Errorf("failed to build spec: %w", err)
			}
			r.definitionsCache.Put(loc.String(), spec)
			ref := openapi.NewRefOrSpec[openapi.Schema](schemaPath)
			return ref, nil
		}
	}
	return nil, fmt.Errorf("failed to resolve type: %s", typeName)
}

func (r *resolver) resolveBasicType(typeName string, opts ...options.SchemaOption) *openapi.RefOrSpec[openapi.Schema] {
	b := openapi.NewSchemaBuilder()
	switch typeName {
	case "int", "int32", "uint", "uint32":
		b.AddType("integer").Format("int32")
	case "int64", "uint64":
		b.AddType("integer").Format("int64")
	case "float32":
		b.AddType("number").Format("float")
	case "float64":
		b.AddType("number").Format("double")
	case "bool":
		b.AddType("boolean")
	case "string":
		b.AddType("string")
	case "file":
		b.AddType("string").Format("binary")
	case "time.Time":
		b.AddType("string").Format("date-time")
	case "uuid.UUID":
		b.AddType("string").Format("uuid")
	case "any":
		// empty schema for any
	default:
		// this is not a basic type, let the caller figure it out
		return nil
	}
	for _, option := range opts {
		option(b)
	}
	return b.Build()
}

func (r *resolver) resolveMap(typeName string, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	matches := mapTypeRegex.FindStringSubmatch(typeName)
	if len(matches) != 3 {
		return nil, fmt.Errorf("invalid map type: %s", typeName)
	}
	keyTypeName := matches[1]
	if keyTypeName != "string" {
		return nil, fmt.Errorf("map key type must be string, got %s", keyTypeName)
	}
	valueTypeName := matches[2]
	valueRef, err := r.Resolve(valueTypeName)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve value type: %w", err)
	}
	valueBoolOrSchema := openapi.NewBoolOrSchema(valueRef)
	valueBoolOrSchema.Allowed = true
	builder := openapi.NewSchemaBuilder().
		AddType("object").
		AdditionalProperties(valueBoolOrSchema)
	for _, option := range opts {
		option(builder)
	}
	return builder.Build(), nil
}

func (r *resolver) resolveArray(typeName string, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	matches := arrayTypeRegex.FindStringSubmatch(typeName)
	if len(matches) != 3 {
		return nil, fmt.Errorf("invalid array type: %s", typeName)
	}
	itemTypeName := matches[2]
	itemRef, err := r.Resolve(itemTypeName)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve item type: %w", err)
	}
	items := openapi.NewBoolOrSchema(itemRef)
	items.Allowed = true
	builder := openapi.NewSchemaBuilder().
		AddType("array").
		Items(items)
	if matches[1] != "" {
		length, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("invalid array length: %s", matches[1])
		}
		builder.MinItems(length)
		builder.MaxItems(length)
	}
	for _, option := range opts {
		option(builder)
	}
	return builder.Build(), nil
}

func (r *resolver) candidates(typeName string, typeParams []string) ([]*locator.Locator, error) {
	var candidates []*locator.Locator

	parts := strings.Split(typeName, ".")
	var pkgName, name string
	loc := &locator.Locator{}

	if len(parts) == 1 {
		pkgName = r.file.Name.Name
		name = parts[0]
		err := r.typeDefCache.Load(context.Background(), r.from)
		if err != nil {
			return nil, fmt.Errorf("failed to load type def cache: %w", err)
		}
		loc.Path = r.from
	} else if len(parts) == 2 {
		pkgName = parts[0]
		name = parts[1]
		for _, imp := range r.file.Imports {
			if imp.Path == nil {
				continue
			}
			importPath := strings.Trim(imp.Path.Value, `"`)
			importParts := strings.Split(importPath, "/")
			importName := importParts[len(importParts)-1]
			if imp.Name != nil {
				importName = imp.Name.Name
			}
			if importName != pkgName {
				continue
			}
			err := r.typeDefCache.Load(context.Background(), importPath)
			if err != nil {
				return nil, fmt.Errorf("failed to load type def cache: %w", err)
			}
			if imp.Name != nil && imp.Name.Name == pkgName {
				loc.Path = importPath
				break
			} else if imp.Name == nil && path.Base(importPath) == pkgName {
				loc.Path = importPath
				break
			}
		}
	} else {
		return nil, fmt.Errorf("invalid type name: %s", typeName)
	}

	loc.Package = pkgName
	loc.Name = name
	loc.TypeParams = typeParams

	candidates = append(candidates, loc)

	return candidates, nil
}
