package cache

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/locator"
	"github.com/sv-tools/openapi"
)

type DefinitionsCache interface {
	Get(*locator.Locator) (*openapi.RefOrSpec[openapi.Schema], bool)
	Put(*locator.Locator, *openapi.RefOrSpec[openapi.Schema]) *openapi.RefOrSpec[openapi.Schema]
	Definitions() map[string]*openapi.RefOrSpec[openapi.Schema]
}

type definitionsCache struct {
	definitions map[string]*openapi.RefOrSpec[openapi.Schema]
	aliases     map[string]string
	taken       map[string]struct{}
}

func NewDefinitionsCache() DefinitionsCache {
	return &definitionsCache{
		definitions: make(map[string]*openapi.RefOrSpec[openapi.Schema]),
		aliases:     make(map[string]string),
		taken:       make(map[string]struct{}),
	}
}

// Get implements [DefinitionsCache].
func (d *definitionsCache) Get(l *locator.Locator) (*openapi.RefOrSpec[openapi.Schema], bool) {
	alias := d.alias(l)
	if _, ok := d.definitions[alias]; ok {
		schemaPath := fmt.Sprintf("#/components/schemas/%s", alias)
		ref := openapi.NewRefOrSpec[openapi.Schema](schemaPath)
		return ref, true
	}
	return nil, false
}

// Put implements [DefinitionsCache].
func (d *definitionsCache) Put(
	l *locator.Locator,
	value *openapi.RefOrSpec[openapi.Schema],
) *openapi.RefOrSpec[openapi.Schema] {
	alias := d.alias(l)
	d.definitions[alias] = value
	schemaPath := fmt.Sprintf("#/components/schemas/%s", alias)
	ref := openapi.NewRefOrSpec[openapi.Schema](schemaPath)
	return ref
}

// Definitions implements [DefinitionsCache].
func (d *definitionsCache) Definitions() map[string]*openapi.RefOrSpec[openapi.Schema] {
	return d.definitions
}

func (d *definitionsCache) alias(l *locator.Locator) string {
	if _, ok := d.aliases[l.String()]; ok {
		return d.aliases[l.String()]
	}
	prefixSha := sha1.Sum([]byte(l.Prefix()))
	sha := hex.EncodeToString(prefixSha[:])
	candidates := []string{
		l.TypeName(),
		fmt.Sprintf("%s.%s", l.Namespace(), l.TypeName()),
		fmt.Sprintf("%s:%s.%s", sha[:4], l.Namespace(), l.TypeName()),
		fmt.Sprintf("%s:%s.%s", sha[:8], l.Namespace(), l.TypeName()),
		fmt.Sprintf("%s:%s.%s", sha, l.Namespace(), l.TypeName()),
	}
	for _, candidate := range candidates {
		if _, ok := d.taken[candidate]; !ok {
			d.aliases[l.String()] = candidate
			d.taken[candidate] = struct{}{}
			return candidate
		}
	}
	// TODO: handle collision
	return ""
}
