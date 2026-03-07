package schema

import (
	"fmt"

	"github.com/sv-tools/openapi"
)

// Resolves types into a schema.
type Resolver interface {
	// Resolve resolves the given type into a schema.
	Resolve(l Locator) (*openapi.Ref, error)

	// Definitions returns the definitions that have been resolved.
	Definitions() map[string]*openapi.RefOrSpec[openapi.Schema]
}

type resolver struct {
	// definitions is a map of the definitions that have been resolved.
	definitions map[string]*openapi.RefOrSpec[openapi.Schema]
}

// NewResolver creates a new resolver.
func NewResolver() Resolver {
	return &resolver{
		definitions: make(map[string]*openapi.RefOrSpec[openapi.Schema]),
	}
}

// Definitions implements [Resolver].
func (r *resolver) Definitions() map[string]*openapi.RefOrSpec[openapi.Schema] {
	return r.definitions
}

// Resolve implements [Resolver].
func (r *resolver) Resolve(l Locator) (*openapi.Ref, error) {
	path := fmt.Sprintf("#/components/schemas/%s", l)
	if _, ok := r.definitions[l.Name]; !ok {
		ref := openapi.NewRefOrSpec[openapi.Schema](path)
		return ref.Ref, nil
	}
	return r.resolve(l)
}

// resolve resolves the given type into a schema, and adds it to the definitions.
func (r *resolver) resolve(Locator) (*openapi.Ref, error) {
	return nil, fmt.Errorf("not implemented")
}
