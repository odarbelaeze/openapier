package schema

import (
	"fmt"
	"go/ast"
	"go/token"
	"log/slog"
	"path"
	"regexp"
	"strconv"
	"strings"

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

type typeDef struct {
	// TypeSpec is the type specification of the type definition.
	TypeSpec *ast.TypeSpec

	// File is the file where the type definition is located.
	File *ast.File

	// Locator is the locator of the type definition.
	Locator *locator.Locator
}

// Resolves types into a schema.
type Resolver interface {
	// Collect collects the type definitions for the given path and file, and caches them.
	Collect(path string, file *ast.File)

	// Resolve resolves the given type into a schema.
	Resolve(typeName string, file *ast.File, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error)

	// Definitions returns the definitions that have been resolved.
	Definitions() map[string]*openapi.RefOrSpec[openapi.Schema]
}

type resolver struct {
	// validatorRegistry is the registry of validators used to validate schemas.
	validatorRegistry validator.Registry

	// definitions is a map of the definitions that have been resolved.
	definitions map[string]*openapi.RefOrSpec[openapi.Schema]

	// cacheByPkg maps package names to type names to type definitions.
	cacheByPkg map[string]map[string]*typeDef

	// loaded is a set of package paths that have been loaded.
	loaded map[string]struct{}
}

// NewResolver creates a new resolver.
func NewResolver(validatorRegistry validator.Registry) Resolver {
	return &resolver{
		validatorRegistry: validatorRegistry,
		definitions:       make(map[string]*openapi.RefOrSpec[openapi.Schema]),
		cacheByPkg:        make(map[string]map[string]*typeDef),
		loaded:            make(map[string]struct{}),
	}
}

// Collect implements [Resolver].
func (r *resolver) Collect(path string, file *ast.File) {
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			loc := locator.Locator{
				Path:    path,
				Package: file.Name.Name,
				Name:    typeSpec.Name.Name,
			}
			slog.Debug("caching type definition", "locator", loc)
			def := &typeDef{
				TypeSpec: typeSpec,
				File:     file,
				Locator:  &loc,
			}

			if r.cacheByPkg[loc.Package] == nil {
				r.cacheByPkg[loc.Package] = make(map[string]*typeDef)
			}
			r.cacheByPkg[loc.Package][loc.Name] = def
		}
	}
	// Mark the path as loaded.
	r.loaded[path] = struct{}{}
}

// Definitions implements [Resolver].
func (r *resolver) Definitions() map[string]*openapi.RefOrSpec[openapi.Schema] {
	return r.definitions
}

// Resolve implements [Resolver].
func (r *resolver) Resolve(typeName string, file *ast.File, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	if strings.HasPrefix(typeName, "[") {
		return r.resolveArray(typeName, file, opts...)
	}

	if strings.HasPrefix(typeName, "map[") {
		return r.resolveMap(typeName, file, opts...)
	}

	basicSchema := parseBasicType(typeName, opts...)
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

	candidates, err := r.candidates(typeName, typeParams, file)
	if err != nil {
		return nil, fmt.Errorf("failed to find any candidates: %w", err)
	}
	for _, loc := range candidates {
		schemaPath := fmt.Sprintf("#/components/schemas/%s", loc)
		if _, ok := r.definitions[loc.String()]; ok {
			ref := openapi.NewRefOrSpec[openapi.Schema](schemaPath)
			return ref, nil
		}
		if pkgCache, pkgOk := r.cacheByPkg[loc.Package]; pkgOk {
			if t, ok := pkgCache[loc.Name]; ok {
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
				spec, err := r.spec(t, aliases, opts...)
				if err != nil {
					return nil, fmt.Errorf("failed to build spec: %w", err)
				}
				r.definitions[loc.String()] = spec
				ref := openapi.NewRefOrSpec[openapi.Schema](schemaPath)
				return ref, nil
			}
		}
	}
	return nil, fmt.Errorf("failed to resolve type: %s", typeName)
}

func (r *resolver) resolveMap(typeName string, file *ast.File, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	matches := mapTypeRegex.FindStringSubmatch(typeName)
	if len(matches) != 3 {
		return nil, fmt.Errorf("invalid map type: %s", typeName)
	}
	keyTypeName := matches[1]
	if keyTypeName != "string" {
		return nil, fmt.Errorf("map key type must be string, got %s", keyTypeName)
	}
	valueTypeName := matches[2]
	valueRef, err := r.Resolve(valueTypeName, file)
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

func (r *resolver) resolveArray(typeName string, file *ast.File, opts ...options.SchemaOption) (*openapi.RefOrSpec[openapi.Schema], error) {
	matches := arrayTypeRegex.FindStringSubmatch(typeName)
	if len(matches) != 3 {
		return nil, fmt.Errorf("invalid array type: %s", typeName)
	}
	itemTypeName := matches[2]
	itemRef, err := r.Resolve(itemTypeName, file)
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

func (r *resolver) candidates(typeName string, typeParams []string, file *ast.File) ([]*locator.Locator, error) {
	var candidates []*locator.Locator

	parts := strings.Split(typeName, ".")
	var pkgName, name string
	loc := &locator.Locator{}

	if len(parts) == 1 {
		pkgName = file.Name.Name
		name = parts[0]
		if pkgCache, ok := r.cacheByPkg[pkgName]; ok {
			if def, ok := pkgCache[name]; ok {
				loc.Path = def.Locator.Path
			} else {
				for _, def := range pkgCache {
					loc.Path = def.Locator.Path
					break
				}
			}
		}
	} else if len(parts) == 2 {
		pkgName = parts[0]
		name = parts[1]
		for _, imp := range file.Imports {
			if imp.Path == nil {
				continue
			}
			importPath := strings.Trim(imp.Path.Value, `"`)
			if _, ok := r.loaded[importPath]; !ok {
				r.loadExternal(importPath)
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

func (r *resolver) loadExternal(importPath string) {
	slog.Debug("loading external package", "importPath", importPath)
}

func (r *resolver) spec(
	t *typeDef,
	aliases map[string]string,
	opts ...options.SchemaOption,
) (*openapi.RefOrSpec[openapi.Schema], error) {
	slog.Debug("finding spec for", "typeName", t.TypeSpec.Name.Name)
	builder := schemaBuilder{
		validatorRegistry: r.validatorRegistry,
		resolver:          r,
		file:              t.File,
		aliases:           aliases,
	}
	return builder.build(t.TypeSpec.Type, opts...)
}
