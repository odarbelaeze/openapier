package parser

import (
	"fmt"
	"go/ast"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/odarbelaeze/openapier/cache"
	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/odarbelaeze/openapier/schema/resolver"
	"github.com/odarbelaeze/openapier/schema/validator"
	ignore "github.com/sabhiram/go-gitignore"
	"github.com/sv-tools/openapi"
	"golang.org/x/mod/modfile"
)

type Parser interface {
	Parse() (*openapi.Extendable[openapi.OpenAPI], error)
}

type parser struct {
	root              string
	main              string
	operationRegistry operation.Registry
	specRegistry      spec.Registry
	validatorRegistry validator.Registry
	parserCache       cache.ParserCache
	typeDefCache      cache.TypeDefCache
	definitionsCache  cache.DefinitionsCache
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

func WithValidatorRegistry(validatorRegistry validator.Registry) ParserOption {
	return func(p *parser) {
		p.validatorRegistry = validatorRegistry
	}
}

func NewParser(root string, main string, opts ...ParserOption) Parser {
	parserCache := cache.NewParserCache()
	p := &parser{
		root:              root,
		main:              main,
		operationRegistry: operation.Default(),
		specRegistry:      spec.Default(),
		validatorRegistry: validator.Default(),
		parserCache:       parserCache,
		typeDefCache:      cache.NewTypeDefCache(root, parserCache),
		definitionsCache:  cache.NewDefinitionsCache(),
	}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Parse implements [Parser].
func (p *parser) Parse() (*openapi.Extendable[openapi.OpenAPI], error) {
	spec := openapi.NewOpenAPIBuilder().Build()
	err := p.parseSpec(spec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse spec: %w", err)
	}

	err = p.parseOperations(spec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operations: %w", err)
	}
	return spec, nil
}

func (p *parser) parseSpec(spec *openapi.Extendable[openapi.OpenAPI]) error {
	node, err := p.parserCache.Parse(path.Join(p.root, p.main))
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

func (p *parser) parseOperations(spec *openapi.Extendable[openapi.OpenAPI]) error {
	modulePath, err := p.findModulePath()
	if err != nil {
		return fmt.Errorf("failed to find module path: %w", err)
	}
	err = p.walkGoFiles(p.root, func(path string) error {
		node, err := p.parserCache.Parse(path)
		if err != nil {
			return fmt.Errorf("failed to parse file: %w", err)
		}
		folder := filepath.Dir(path)
		relativeFolder, err := filepath.Rel(p.root, folder)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}
		from := filepath.Join(*modulePath, relativeFolder)
		var opErr error
		ast.Inspect(node, func(n ast.Node) bool {
			if function, ok := n.(*ast.FuncDecl); ok {
				if function.Name.Name == "main" {
					return false
				}
				resolver := resolver.NewResolver(
					p.validatorRegistry,
					p.typeDefCache,
					p.definitionsCache,
					resolver.NewSchemaBuilder,
					node,
					from,
				)
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
	if len(p.definitionsCache.Definitions()) > 0 {
		if spec.Spec.Components == nil {
			spec.Spec.Components = openapi.NewComponents()
		}
		spec.Spec.Components.Spec.Schemas = p.definitionsCache.Definitions()
	}
	return nil
}

func (p *parser) findModulePath() (*string, error) {
	gomodPath := path.Join(p.root, "go.mod")
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
	modulePath := f.Module.Mod.Path
	return &modulePath, nil
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
