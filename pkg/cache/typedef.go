package cache

import (
	"context"
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

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

	// EnumValues is the list of enum values, if the type is an enum.
	EnumValues []any
}

type TypeDefCache interface {
	Load(ctx context.Context, pkgName string) error
	Get(pkgName string, typeName string) (*TypeDef, bool)
}

type typeDefCache struct {
	root        string
	parserCache ParserCache
	cache       map[string]map[string]*TypeDef
	loaded      map[string]struct{}
}

func NewTypeDefCache(root string, parserCache ParserCache) TypeDefCache {
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

	// First pass: collect all types
	for _, pkg := range packages {
		files := make([]*ast.File, 0, len(pkg.GoFiles))
		for _, filename := range pkg.GoFiles {
			file, err := t.parserCache.Parse(filename)
			if err != nil {
				return fmt.Errorf("failed to parse file %s: %w", filename, err)
			}
			files = append(files, file)

			// First pass: collect all types
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
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

		// Second pass: collect all constants
		for _, file := range files {
			for _, decl := range file.Decls {
				// Only process const declarations; never var
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
					var prevType string
					var prevValues []ast.Expr
					iotaCounter := 0

					for _, spec := range genDecl.Specs {
						valueSpec, ok := spec.(*ast.ValueSpec)
						if !ok {
							continue
						}

						// Update type if present
						if valueSpec.Type != nil {
							if typeIdent, ok := valueSpec.Type.(*ast.Ident); ok {
								prevType = typeIdent.Name
							} else {
								prevType = ""
							}
						}

						// Update values if present
						if len(valueSpec.Values) > 0 {
							prevValues = valueSpec.Values
						}

						if prevType != "" && ast.IsExported(prevType) {
							if pkgCache, ok := t.cache[pkgName]; ok {
								if typeDef, ok := pkgCache[prevType]; ok {
									// We have a candidate for an enum value
									for i := range valueSpec.Names {
										var valExpr ast.Expr
										if i < len(prevValues) {
											valExpr = prevValues[i]
										}
										if valExpr != nil {
											if val, ok := t.evaluate(valExpr, iotaCounter); ok {
												typeDef.EnumValues = append(typeDef.EnumValues, val)
											}
										}
									}
								}
							}
						}

						// Increment iota counter for each const declaration
						iotaCounter++
					}
				}
			}
		}
	}

	// Tag the package as loaded
	t.loaded[pkgName] = struct{}{}
	return nil
}

func (t *typeDefCache) evaluate(expr ast.Expr, iotaValue int) (any, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		switch e.Kind {
		case token.INT:
			v, err := strconv.Atoi(e.Value)
			if err != nil {
				return nil, false
			}
			return v, true
		case token.STRING:
			v, err := strconv.Unquote(e.Value)
			if err != nil {
				return nil, false
			}
			return v, true
		}
	case *ast.Ident:
		if e.Name == "iota" {
			return iotaValue, true
		}
	case *ast.BinaryExpr:
		left, okL := t.evaluate(e.X, iotaValue)
		right, okR := t.evaluate(e.Y, iotaValue)
		if !okL || !okR {
			return nil, false
		}
		l, okLi := left.(int)
		r, okRi := right.(int)
		if !okLi || !okRi {
			return nil, false
		}
		switch e.Op {
		case token.SHL:
			return l << uint(r), true
		case token.ADD:
			return l + r, true
		case token.SUB:
			return l - r, true
		case token.MUL:
			return l * r, true
		}
	case *ast.ParenExpr:
		return t.evaluate(e.X, iotaValue)
	case *ast.UnaryExpr:
		val, ok := t.evaluate(e.X, iotaValue)
		if !ok {
			return nil, false
		}
		if v, ok := val.(int); ok {
			switch e.Op {
			case token.SUB:
				return -v, true
			case token.ADD:
				return v, true
			}
		}
	}
	return nil, false
}
