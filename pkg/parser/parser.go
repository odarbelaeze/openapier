package parser

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/sv-tools/openapi"
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
	err = p.parseOperations(root, spec)
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

func (p *parser) parseOperations(root string, spec *openapi.Extendable[openapi.OpenAPI]) error {
	fileSet := token.NewFileSet()
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		node, err := goparser.ParseFile(fileSet, path, nil, goparser.ParseComments)
		if err != nil {
			return fmt.Errorf("failed to parse file: %w", err)
		}
		var opErr error
		ast.Inspect(node, func(n ast.Node) bool {
			if function, ok := n.(*ast.FuncDecl); ok {
				operation := operation.NewOperation()
				if function.Doc == nil {
					return false
				}
				for _, comment := range function.Doc.List {
					err = p.operationRegistry.Parse(comment.Text, operation)
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
	return nil
}
