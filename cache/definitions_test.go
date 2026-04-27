package cache_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/cache"
	"github.com/odarbelaeze/openapier/schema/locator"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestDefinitionsCache(t *testing.T) {
	c := cache.NewDefinitionsCache()
	assert.NotNil(t, c)

	l := &locator.Locator{
		Path:    "github.com/test/pkg",
		Package: "pkg",
		Name:    "TestSchema",
	}
	spec := openapi.NewSchemaBuilder().AddType("object").Build()

	// Test Get on empty cache
	val, ok := c.Get(l)
	assert.False(t, ok)
	assert.Nil(t, val)

	// Test Put
	ref := c.Put(l, spec)
	assert.NotNil(t, ref)
	// The first alias should just be the TypeName
	expectedRef := "#/components/schemas/" + l.TypeName()
	assert.Equal(t, expectedRef, ref.Ref.Ref)

	// Test Get after Put
	val, ok = c.Get(l)
	assert.True(t, ok)
	assert.Equal(t, ref, val)

	// Test collision
	l2 := &locator.Locator{
		Path:    "github.com/other/pkg",
		Package: "other",
		Name:    "TestSchema",
	}
	ref2 := c.Put(l2, spec)
	// Since "TestSchema" is taken, it should be "other.TestSchema"
	expectedRef2 := "#/components/schemas/other.TestSchema"
	assert.Equal(t, expectedRef2, ref2.Ref.Ref)

	// Test Definitions
	defs := c.Definitions()
	assert.Len(t, defs, 2)
	assert.Equal(t, spec, defs[l.TypeName()])
	assert.Equal(t, spec, defs[l2.Namespace()+"."+l2.TypeName()])
}
