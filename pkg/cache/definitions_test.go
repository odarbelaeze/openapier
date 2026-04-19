package cache_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/cache"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestDefinitionsCache(t *testing.T) {
	c := cache.NewDefinitionsCache()
	assert.NotNil(t, c)

	key := "TestSchema"
	spec := openapi.NewSchemaBuilder().AddType("object").Build()

	// Test Get on empty cache
	val, ok := c.Get(key)
	assert.False(t, ok)
	assert.Nil(t, val)

	// Test Put
	c.Put(key, spec)

	// Test Get after Put
	val, ok = c.Get(key)
	assert.True(t, ok)
	assert.Equal(t, spec, val)

	// Test Definitions
	defs := c.Definitions()
	assert.Len(t, defs, 1)
	assert.Equal(t, spec, defs[key])
}
