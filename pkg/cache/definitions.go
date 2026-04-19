package cache

import (
	"github.com/odarbelaeze/openapier/pkg/schema/locator"
	"github.com/sv-tools/openapi"
)

type DefinitionsCache interface {
	Get(*locator.Locator) (*openapi.RefOrSpec[openapi.Schema], bool)
	Put(*locator.Locator, *openapi.RefOrSpec[openapi.Schema])
	Definitions() map[string]*openapi.RefOrSpec[openapi.Schema]
}

type definitionsCache struct {
	definitions map[string]*openapi.RefOrSpec[openapi.Schema]
}

func NewDefinitionsCache() DefinitionsCache {
	return &definitionsCache{
		definitions: make(map[string]*openapi.RefOrSpec[openapi.Schema]),
	}
}

// Get implements [DefinitionsCache].
func (d *definitionsCache) Get(l *locator.Locator) (*openapi.RefOrSpec[openapi.Schema], bool) {
	val, ok := d.definitions[l.String()]
	return val, ok
}

// Put implements [DefinitionsCache].
func (d *definitionsCache) Put(l *locator.Locator, value *openapi.RefOrSpec[openapi.Schema]) {
	d.definitions[l.String()] = value
}

// Definitions implements [DefinitionsCache].
func (d *definitionsCache) Definitions() map[string]*openapi.RefOrSpec[openapi.Schema] {
	return d.definitions
}
