package schema

import (
	"go/ast"
	"go/token"
	"log/slog"
)

type TypeDef struct {
	// TypeSpec is the type specification of the type definition.
	TypeSpec *ast.TypeSpec

	// File is the file where the type definition is located.
	File *ast.File
}

type TypeDefCache interface {
	// Get returns the cached type definition for the given locator, or nil if not found.
	Get(locator *Locator) *TypeDef

	// Collect collects the type definitions for the given path and file, and caches them.
	Collect(path string, file *ast.File)
}

type typeDefCache struct {
	// cache is a map of locators to type definitions.
	cache map[string]*TypeDef
}

// NewTypeSpecCache creates a new type specification cache.
func NewTypeSpecCache() TypeDefCache {
	return &typeDefCache{
		cache: make(map[string]*TypeDef),
	}
}

// Collect implements [TypeDefCache].
func (t *typeDefCache) Collect(path string, file *ast.File) {
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
			t.cache[locator.String()] = &TypeDef{
				TypeSpec: typeSpec,
				File:     file,
			}
		}
	}
}

// Get implements [TypeDefCache].
func (t *typeDefCache) Get(locator *Locator) *TypeDef {
	if def, ok := t.cache[locator.String()]; ok {
		slog.Debug("cache hit for type definition", "locator", locator)
		return def
	}
	slog.Debug("cache miss for type definition", "locator", locator)
	return nil
}
