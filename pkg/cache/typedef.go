package cache

import (
	"context"
	"fmt"
	"go/ast"
	"go/token"
	"log/slog"
	"strconv"
	"strings"

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
	var prevType *string
	var iotaCounter int
	var useIota bool
	for _, pkg := range packages {
		for _, filename := range pkg.GoFiles {
			file, err := t.parserCache.Parse(filename)
			if err != nil {
				return fmt.Errorf("failed to parse file %s: %w", filename, err)
			}
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if valueSpec, ok := spec.(*ast.ValueSpec); ok {
							if valueSpec.Type == nil && prevType == nil {
								continue
							}
							var typeName string
							if valueSpec.Type != nil {
								if typeIdent, ok := valueSpec.Type.(*ast.Ident); ok {
									typeName = typeIdent.Name
									prevType = &typeName
									iotaCounter = 0
								} else {
									continue
								}
							} else {
								if prevType == nil {
									// This should not happen, but if it does, skip it
									continue
								}
								iotaCounter++
								typeName = *prevType
							}
							var value any
							if len(valueSpec.Values) == 0 {
								if useIota {
									value = iotaCounter
								} else {
									continue
								}
							} else if len(valueSpec.Values) > 1 {
								continue
							} else {
								switch concreteValue := valueSpec.Values[0].(type) {
								case *ast.Ident:
									{
										if concreteValue.Name == "iota" {
											useIota = true
											value = iotaCounter
										} else {
											useIota = false
										}
									}
								case *ast.BasicLit:
									{
										useIota = false
										switch concreteValue.Kind {
										case token.STRING:
											value = strings.Trim(concreteValue.Value, "\"")
										case token.INT:
											{
												var err error
												value, err = strconv.Atoi(concreteValue.Value)
												if err != nil {
													continue
												}
											}
										default:
											continue
										}
									}
								default:
									continue
								}
							}
							slog.Debug(
								"found a value spec",
								"type", typeName,
								"value", value,
							)
							// TODO: Filter out any other non-custom types
							if !ast.IsExported(typeName) {
								slog.Debug("skipping unexported type", "type", typeName)
								continue
							}
							if _, ok := t.cache[pkgName]; !ok {
								slog.Warn("please define your types before any const exprs")
								continue
							}
							if _, ok := t.cache[pkgName][typeName]; !ok {
								slog.Warn("please define your types before any const exprs")
								continue
							}
							typeDef := t.cache[pkgName][typeName]
							typeDef.EnumValues = append(typeDef.EnumValues, value)
						}
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
