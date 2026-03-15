package schema

import (
	"fmt"
	"go/ast"
	"go/token"
	"log/slog"
	"path"
	"strings"

	"github.com/sv-tools/openapi"
)

type TypeDef struct {
	// TypeSpec is the type specification of the type definition.
	TypeSpec *ast.TypeSpec

	// File is the file where the type definition is located.
	File *ast.File

	// Locator is the locator of the type definition.
	Locator *Locator
}

// Resolves types into a schema.
type Resolver interface {
	// Collect collects the type definitions for the given path and file, and caches them.
	Collect(path string, file *ast.File)

	// Resolve resolves the given type into a schema.
	Resolve(typeName string, file *ast.File) (*openapi.RefOrSpec[openapi.Schema], error)

	// Definitions returns the definitions that have been resolved.
	Definitions() map[string]*openapi.RefOrSpec[openapi.Schema]
}

type resolver struct {
	// definitions is a map of the definitions that have been resolved.
	definitions map[string]*openapi.RefOrSpec[openapi.Schema]

	// cache is a map of locators to type definitions.
	cache map[string]*TypeDef

	// loaded is a set of package paths that have been loaded.
	loaded map[string]struct{}
}

// NewResolver creates a new resolver.
func NewResolver() Resolver {
	return &resolver{
		definitions: make(map[string]*openapi.RefOrSpec[openapi.Schema]),
		cache:       make(map[string]*TypeDef),
		loaded:      make(map[string]struct{}),
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
			locator := Locator{
				Path:    path,
				Package: file.Name.Name,
				Name:    typeSpec.Name.Name,
			}
			slog.Debug("caching type definition", "locator", locator)
			r.cache[locator.String()] = &TypeDef{
				TypeSpec: typeSpec,
				File:     file,
				Locator:  &locator,
			}
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
func (r *resolver) Resolve(typeName string, file *ast.File) (*openapi.RefOrSpec[openapi.Schema], error) {
	if strings.HasPrefix(typeName, "[]") {
		itemTypeName := typeName[2:]
		itemRef, err := r.Resolve(itemTypeName, file)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve item type: %w", err)
		}
		return openapi.
			NewSchemaBuilder().
			AddType("array").
			Items(openapi.NewBoolOrSchema(itemRef)).
			Build(), nil
	}

	basicSchema := ParseBasicType(typeName)
	if basicSchema != nil {
		return basicSchema, nil
	}

	candidates, err := r.candidates(typeName, file)
	if err != nil {
		return nil, fmt.Errorf("failed to find any candidates: %w", err)
	}
	for _, loc := range candidates {
		schemasPath := fmt.Sprintf("#/components/schemas/%s", loc)
		if _, ok := r.definitions[loc.Name]; !ok {
			ref := openapi.NewRefOrSpec[openapi.Schema](schemasPath)
			return ref, nil
		}
		if t, ok := r.cache[loc.String()]; ok {
			spec := r.spec(t)
			r.definitions[schemasPath] = spec
			ref := openapi.NewRefOrSpec[openapi.Schema](schemasPath)
			return ref, nil
		}
	}
	return nil, fmt.Errorf("failed to resolve type: %s", typeName)
}

func (r *resolver) candidates(typeName string, file *ast.File) ([]*Locator, error) {
	var candidates []*Locator

	parts := strings.Split(typeName, ".")
	var pkgName, name string
	loc := &Locator{}

	if len(parts) == 1 {
		pkgName = file.Name.Name
		name = parts[0]
		for _, def := range r.cache {
			if def.File == file {
				loc.Path = def.Locator.Path
				break
			}
			if def.Locator.Package == pkgName {
				loc.Path = def.Locator.Path
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

	candidates = append(candidates, loc)

	return candidates, nil
}

func (r *resolver) loadExternal(importPath string) {
	slog.Debug("loading external package", "importPath", importPath)
}

func (r *resolver) spec(t *TypeDef) *openapi.RefOrSpec[openapi.Schema] {
	slog.Debug("finding spec for", "typeName", t.TypeSpec.Name.Name)
	return openapi.NewSchemaBuilder().
		AddType("object").
		AddProperty("field", openapi.NewSchemaBuilder().AddType("string").Build()).
		Build()
}
