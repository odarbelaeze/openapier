package cache

import (
	"context"
	"fmt"
	"go/ast"
	"go/token"
	"iter"
	"strconv"

	"github.com/odarbelaeze/openapier/schema/locator"
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
	Load(ctx context.Context, pkgPath string) error
	Get(pkgPath string, typeName string) (*TypeDef, bool)
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
func (t *typeDefCache) Get(pkgPath string, typeName string) (*TypeDef, bool) {
	if _, ok := t.cache[pkgPath]; !ok {
		return nil, false
	}
	def, ok := t.cache[pkgPath][typeName]
	return def, ok
}

// Load implements [TypeDefCache].
func (t *typeDefCache) Load(ctx context.Context, pkgPath string) error {
	if _, ok := t.loaded[pkgPath]; ok {
		return nil
	}
	packages, err := packages.Load(&packages.Config{
		Dir: t.root,
	}, pkgPath)
	if err != nil {
		return fmt.Errorf("failed to load package %s: %w", pkgPath, err)
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
			for decl := range genericDeclarations(file, token.TYPE) {
				for typeSpec := range typeSpecs(decl) {
					if _, ok := t.cache[pkgPath]; !ok {
						t.cache[pkgPath] = make(map[string]*TypeDef)
					}
					t.cache[pkgPath][typeSpec.Name.Name] = &TypeDef{
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

		// Second pass: collect all constants
		for _, file := range files {
			for decl := range genericDeclarations(file, token.CONST) {
				var prevType string
				var prevValues []ast.Expr
				iotaCounter := 0

				for valueSpec := range valueSpecs(decl) {
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

					// Only process exported types
					if prevType != "" && ast.IsExported(prevType) {
						if typeDef, ok := t.Get(pkgPath, prevType); ok {
							values := enumValues(valueSpec, prevValues, iotaCounter)
							typeDef.EnumValues = append(typeDef.EnumValues, values...)
						}
					}

					// Increment iota counter for each const declaration
					iotaCounter++
				}
			}
		}
	}

	// Tag the package as loaded
	t.loaded[pkgPath] = struct{}{}
	return nil
}

func enumValues(valueSpec *ast.ValueSpec, prevValues []ast.Expr, iotaCounter int) []any {
	result := make([]any, 0, len(valueSpec.Names))
	for i, name := range valueSpec.Names {
		if name != nil && name.Name == "_" {
			// Skip iota placeholder values
			continue
		}
		var valExpr ast.Expr
		if i < len(prevValues) {
			valExpr = prevValues[i]
		}
		if valExpr != nil {
			if val, ok := evaluate(valExpr, iotaCounter); ok {
				result = append(result, val)
			}
		}
	}
	return result
}

func evaluate(expr ast.Expr, iotaValue int) (any, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return evaluateBasicLit(e)
	case *ast.Ident:
		if e.Name == "iota" {
			return iotaValue, true
		}
	case *ast.BinaryExpr:
		return evaluateBinaryExpr(e, iotaValue)
	case *ast.ParenExpr:
		return evaluate(e.X, iotaValue)
	case *ast.UnaryExpr:
		return evaluateUnaryExpr(e, iotaValue)
	}
	return nil, false
}

func evaluateBasicLit(e *ast.BasicLit) (any, bool) {
	switch e.Kind {
	case token.INT:
		v, err := strconv.ParseInt(e.Value, 0, 0)
		if err != nil {
			return nil, false
		}
		// TODO: handle overflow
		return int(v), true
	case token.STRING:
		v, err := strconv.Unquote(e.Value)
		if err != nil {
			return nil, false
		}
		return v, true
	default:
		return nil, false
	}
}

func evaluateBinaryExpr(e *ast.BinaryExpr, iotaValue int) (any, bool) {
	left, okL := evaluate(e.X, iotaValue)
	right, okR := evaluate(e.Y, iotaValue)
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
	case token.SHR:
		return l >> uint(r), true
	case token.AND:
		return l & r, true
	case token.OR:
		return l | r, true
	case token.XOR:
		return l ^ r, true
	case token.AND_NOT:
		return l &^ r, true
	case token.ADD:
		return l + r, true
	case token.SUB:
		return l - r, true
	case token.MUL:
		return l * r, true
	case token.QUO:
		if r == 0 {
			return nil, false
		}
		return l / r, true
	case token.REM:
		if r == 0 {
			return nil, false
		}
		return l % r, true
	}
	return nil, false
}

func evaluateUnaryExpr(e *ast.UnaryExpr, iotaValue int) (any, bool) {
	val, ok := evaluate(e.X, iotaValue)
	if !ok {
		return nil, false
	}
	if v, ok := val.(int); ok {
		switch e.Op {
		case token.SUB:
			return -v, true
		case token.ADD:
			return v, true
		default:
			return nil, false
		}
	}
	return nil, false
}

func genericDeclarations(file *ast.File, ofType token.Token) iter.Seq[*ast.GenDecl] {
	return func(yield func(*ast.GenDecl) bool) {
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == ofType {
				if !yield(genDecl) {
					return
				}
			}
		}
	}
}

func typeSpecs(decl *ast.GenDecl) iter.Seq[*ast.TypeSpec] {
	return func(yield func(*ast.TypeSpec) bool) {
		for _, spec := range decl.Specs {
			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
				if !yield(typeSpec) {
					return
				}
			}
		}
	}
}

func valueSpecs(decl *ast.GenDecl) iter.Seq[*ast.ValueSpec] {
	return func(yield func(*ast.ValueSpec) bool) {
		for _, spec := range decl.Specs {
			if valueSpec, ok := spec.(*ast.ValueSpec); ok {
				if !yield(valueSpec) {
					return
				}
			}
		}
	}
}
