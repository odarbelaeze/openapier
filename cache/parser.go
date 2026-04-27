package cache

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log/slog"
	"path/filepath"
	"sync"
)

type ParserCache interface {
	Parse(filename string) (*ast.File, error)
}

type parserCache struct {
	cache map[string]*ast.File
	fset  *token.FileSet
	mutex sync.RWMutex
}

func NewParserCache() ParserCache {
	return &parserCache{
		cache: make(map[string]*ast.File),
		fset:  token.NewFileSet(),
	}
}

// Parse implements [ParserCache].
func (p *parserCache) Parse(filename string) (*ast.File, error) {
	filename, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("error getting absolute path for file %s: %w", filename, err)
	}
	p.mutex.RLock()
	file, ok := p.cache[filename]
	p.mutex.RUnlock()
	if ok {
		slog.Debug("cache hit", "filename", filename)
		return file, nil
	}
	file, err = parser.ParseFile(p.fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing file %s: %w", filename, err)
	}
	p.mutex.Lock()
	p.cache[filename] = file
	p.mutex.Unlock()
	return file, nil
}
