package cache

import "github.com/sv-tools/openapi"

type DefinitionsCache interface {
	Get(string) (*openapi.RefOrSpec[openapi.Schema], bool)
	Put(string, *openapi.RefOrSpec[openapi.Schema])
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
func (d *definitionsCache) Get(key string) (*openapi.RefOrSpec[openapi.Schema], bool) {
	val, ok := d.definitions[key]
	return val, ok
}

// Put implements [DefinitionsCache].
func (d *definitionsCache) Put(key string, value *openapi.RefOrSpec[openapi.Schema]) {
	d.definitions[key] = value
}

// Definitions implements [DefinitionsCache].
func (d *definitionsCache) Definitions() map[string]*openapi.RefOrSpec[openapi.Schema] {
	return d.definitions
}
