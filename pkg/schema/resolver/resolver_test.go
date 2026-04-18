package resolver

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/cache"
	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	"github.com/stretchr/testify/assert"
)

func TestResolve_BasicTypes(t *testing.T) {
	registry := validator.NewMockRegistry(t)
	typeDefCache := cache.NewMockTypeDefCache(t)
	definitionsCache := cache.NewMockDefinitionsCache(t)
	factory := func(validatorRegistry validator.Registry, resolver Resolver, aliases map[string]string) SchemaBuilder {
		return nil
	}

	r := NewResolver(registry, typeDefCache, definitionsCache, factory, nil, "")

	tests := []struct {
		typeName string
		expected string
		format   string
	}{
		{"int", "integer", "int32"},
		{"int64", "integer", "int64"},
		{"float64", "number", "double"},
		{"bool", "boolean", ""},
		{"string", "string", ""},
		{"time.Time", "string", "date-time"},
		{"uuid.UUID", "string", "uuid"},
	}

	for _, tt := range tests {
		t.Run(tt.typeName, func(t *testing.T) {
			schema, err := r.Resolve(tt.typeName)
			assert.NoError(t, err)
			assert.NotNil(t, schema)
			assert.Equal(t, tt.expected, (*schema.Spec.Type)[0])
			if tt.format != "" {
				assert.Equal(t, tt.format, schema.Spec.Format)
			}
		})
	}
}
