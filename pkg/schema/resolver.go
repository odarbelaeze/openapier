package schema

import (
	"fmt"
	"go/ast"
	"go/token"
	"log/slog"

	"github.com/sv-tools/openapi"
)

type TypeDef struct {
	// TypeSpec is the type specification of the type definition.
	TypeSpec *ast.TypeSpec

	// File is the file where the type definition is located.
	File *ast.File
}

// Resolves types into a schema.
type Resolver interface {
	// Collect collects the type definitions for the given path and file, and caches them.
	Collect(path string, file *ast.File)

	// Resolve resolves the given type into a schema.
	Resolve(l *Locator) (*openapi.Ref, error)

	// Definitions returns the definitions that have been resolved.
	Definitions() map[string]*openapi.RefOrSpec[openapi.Schema]
}

type resolver struct {
	// definitions is a map of the definitions that have been resolved.
	definitions map[string]*openapi.RefOrSpec[openapi.Schema]

	// cache is a map of locators to type definitions.
	cache map[string]*TypeDef
}

// NewResolver creates a new resolver.
func NewResolver() Resolver {
	return &resolver{
		definitions: make(map[string]*openapi.RefOrSpec[openapi.Schema]),
		cache:       make(map[string]*TypeDef),
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
			}
		}
	}
}

// Definitions implements [Resolver].
func (r *resolver) Definitions() map[string]*openapi.RefOrSpec[openapi.Schema] {
	return r.definitions
}

// Resolve implements [Resolver].
func (r *resolver) Resolve(l *Locator) (*openapi.Ref, error) {
	path := fmt.Sprintf("#/components/schemas/%s", l)
	if _, ok := r.definitions[l.Name]; !ok {
		ref := openapi.NewRefOrSpec[openapi.Schema](path)
		return ref.Ref, nil
	}
	return r.resolve(l)
}

// resolve resolves the given type into a schema, and adds it to the definitions.
func (r *resolver) resolve(l *Locator) (*openapi.Ref, error) {
	if def, ok := r.cache[l.String()]; ok {
		slog.Debug("cache hit for type definition", "locator", l)
		_ = def // Just to avoid unused var
	} else {
		slog.Debug("cache miss for type definition", "locator", l)
	}

	return nil, fmt.Errorf("not implemented")
}
