package parser

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/odarbelaeze/openapier/pkg/schema/resolver"
	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	ignore "github.com/sabhiram/go-gitignore"
	"github.com/sv-tools/openapi"
	"golang.org/x/mod/modfile"
)

type Parser interface {
	Parse(root string, main string) (*openapi.Extendable[openapi.OpenAPI], error)
}

type parser struct {
	operationRegistry operation.Registry
	specRegistry      spec.Registry
}

type ParserOption func(*parser)

func WithOperationRegistry(operationRegistry operation.Registry) ParserOption {
	return func(p *parser) {
		p.operationRegistry = operationRegistry
	}
}

func WithSpecRegistry(specRegistry spec.Registry) ParserOption {
	return func(p *parser) {
		p.specRegistry = specRegistry
	}
}

func NewParser(opts ...ParserOption) Parser {
	p := &parser{
		operationRegistry: operation.DefaultRegistry,
		specRegistry:      spec.DefaultRegistry,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// Parse implements [Parser].
func (p *parser) Parse(root string, main string) (*openapi.Extendable[openapi.OpenAPI], error) {
	spec := openapi.NewOpenAPIBuilder().Build()
	err := p.parseSpec(path.Join(root, main), spec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse spec: %w", err)
	}

	resolver, err := p.parseTypes(root)
	if err != nil {
		return nil, fmt.Errorf("failed to parse types: %w", err)
	}

	err = p.parseOperations(root, resolver, spec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operations: %w", err)
	}
	return spec, nil
}

func (p *parser) parseSpec(main string, spec *openapi.Extendable[openapi.OpenAPI]) error {
	fileSet := token.NewFileSet()
	node, err := goparser.ParseFile(fileSet, main, nil, goparser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse main file: %w", err)
	}
	for _, comment := range node.Comments {
		for _, line := range comment.List {
			err := p.specRegistry.Parse(line.Text, spec)
			if err != nil {
				return fmt.Errorf("failed to parse spec comment: %w", err)
			}
		}
	}
	return nil
}

func (p *parser) parseTypes(root string) (resolver.Resolver, error) {
	gomodPath := path.Join(root, "go.mod")
	gomodData, err := os.ReadFile(gomodPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("go.mod file not found in root directory: %s", gomodPath)
		}
		return nil, fmt.Errorf("failed to open go.mod file: %w", err)
	}
	f, err := modfile.ParseLax(gomodPath, gomodData, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go.mod file: %w", err)
	}
	if f.Module == nil {
		return nil, fmt.Errorf("module declaration not found in go.mod file")
	}
	cache := resolver.NewResolver(validator.Default())
	fileSet := token.NewFileSet()
	err = p.walkGoFiles(root, func(p string) error {
		node, err := goparser.ParseFile(fileSet, p, nil, goparser.ParseComments)
		if err != nil {
			return fmt.Errorf("failed to parse file: %w", err)
		}
		realPath, err := filepath.Rel(root, p)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}
		realPath = path.Join(f.Module.Mod.Path, filepath.Dir(realPath))
		cache.Collect(realPath, node)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}
	return cache, nil
}

func (p *parser) parseOperations(
	root string,
	resolver resolver.Resolver,
	spec *openapi.Extendable[openapi.OpenAPI],
) error {
	fileSet := token.NewFileSet()
	err := p.walkGoFiles(root, func(path string) error {
		node, err := goparser.ParseFile(fileSet, path, nil, goparser.ParseComments)
		if err != nil {
			return fmt.Errorf("failed to parse file: %w", err)
		}
		var opErr error
		ast.Inspect(node, func(n ast.Node) bool {
			if function, ok := n.(*ast.FuncDecl); ok {
				if function.Name.Name == "main" {
					return false
				}
				operation := operation.NewOperation(resolver)
				if function.Doc == nil {
					return false
				}
				for _, comment := range function.Doc.List {
					err = p.operationRegistry.Parse(comment.Text, node, operation)
					if err != nil {
						opErr = err
						return false
					}
				}
				err := operation.Attach(spec)
				if err != nil {
					opErr = err
					return false
				}
				return false
			}
			return true
		})
		if opErr != nil {
			return fmt.Errorf("failed to parse operations: %w", opErr)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}
	if len(resolver.Definitions()) > 0 {
		spec.Spec.Components = openapi.NewComponents()
		spec.Spec.Components.Spec.Schemas = resolver.Definitions()
	}
	return nil
}

func (p *parser) walkGoFiles(root string, fn func(path string) error) error {
	var ign *ignore.GitIgnore
	ign, err := ignore.CompileIgnoreFile(filepath.Join(root, ".gitignore"))
	if err != nil {
		slog.Debug("failed to compile gitignore", "err", err)
	}
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ign != nil && ign.MatchesPath(path) {
			return nil
		}
		if info.IsDir() {
			if info.Name() == ".git" || info.Name() == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		return fn(path)
	})
}
