package resolver

import (
	"context"
	"fmt"
	"go/ast"

	"github.com/odarbelaeze/openapier/pkg/cache"
	"github.com/odarbelaeze/openapier/pkg/schema/locator"
	"golang.org/x/tools/go/packages"
)

type TypeDef struct {
	// TypeSpec is the type specification of the type definition.
	TypeSpec *ast.TypeSpec

	// File is the file where the type definition is located.
	File *ast.File

	// Locator is the locator of the type definition.
	Locator *locator.Locator
}

type TypeDefCache interface {
	Load(ctx context.Context, pkgName string) error
	Get(pkgName string, typeName string) (*TypeDef, bool)
}

type typeDefCache struct {
	root        string
	parserCache cache.ParserCache
	cache       map[string]map[string]*TypeDef
	loaded      map[string]struct{}
}

func NewTypeDefCache(root string, parserCache cache.ParserCache) TypeDefCache {
	return &typeDefCache{
		root:        root,
		parserCache: parserCache,
		cache:       make(map[string]map[string]*TypeDef),
		loaded:      make(map[string]struct{}),
	}
}

// Get implements [TypeDefCache].
func (t *typeDefCache) Get(pkgName string, typeName string) (*TypeDef, bool) {
	if _, ok := t.cache[pkgName]; !ok {
		return nil, false
	}
	def, ok := t.cache[pkgName][typeName]
	return def, ok
}

// Load implements [TypeDefCache].
func (t *typeDefCache) Load(ctx context.Context, pkgName string) error {
	if _, ok := t.loaded[pkgName]; ok {
		return nil
	}
	packages, err := packages.Load(&packages.Config{
		Dir: t.root,
	}, pkgName)
	if err != nil {
		return fmt.Errorf("failed to load package %s: %w", pkgName, err)
	}
	for _, pkg := range packages {
		for _, filename := range pkg.GoFiles {
			file, err := t.parserCache.Parse(filename)
			if err != nil {
				return fmt.Errorf("failed to parse file %s: %w", filename, err)
			}
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if _, ok := t.cache[pkgName]; !ok {
								t.cache[pkgName] = make(map[string]*TypeDef)
							}
							t.cache[pkgName][typeSpec.Name.Name] = &TypeDef{
								TypeSpec: typeSpec,
								File:     file,
								Locator: &locator.Locator{
									Path:    pkg.PkgPath,
									Package: pkg.Name,
									Name:    typeSpec.Name.Name,
								},
							}
						}
					}
				}
			}
		}
	}
	t.loaded[pkgName] = struct{}{}
	return nil
}
