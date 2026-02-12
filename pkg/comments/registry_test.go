package comments_test

import (
	"errors"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

type mockComment struct {
	tag   string
	usage string
	err   error
}

func (m *mockComment) Tag() string {
	return m.tag
}

func (m *mockComment) Usage() string {
	return m.usage
}

func (m *mockComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	return m.err
}

func TestRegistry_Parse_UnknownTag(t *testing.T) {
	registry := comments.NewRegistry()
	o := openapi.NewOpenAPIBuilder().Build()
	err := registry.Parse("// @unknown foo", o)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown comment tag: unknown")
}

func TestRegistry_Parse_KnownTag(t *testing.T) {
	registry := comments.NewRegistry()
	mock := &mockComment{tag: "mock", usage: "// @mock", err: nil}
	registry.Register(mock)

	o := openapi.NewOpenAPIBuilder().Build()
	err := registry.Parse("// @mock foo", o)
	assert.NoError(t, err)
}

func TestRegistry_Parse_ErrorFromHandler(t *testing.T) {
	registry := comments.NewRegistry()
	mock := &mockComment{tag: "mock", usage: "// @mock", err: errors.New("handler error")}
	registry.Register(mock)

	o := openapi.NewOpenAPIBuilder().Build()
	err := registry.Parse("// @mock foo", o)
	assert.Error(t, err)
	assert.EqualError(t, err, "handler error")
}

func TestRegistry_Parse_IgnoreNonComments(t *testing.T) {
	registry := comments.NewRegistry()
	o := openapi.NewOpenAPIBuilder().Build()
	err := registry.Parse("not a comment", o)
	assert.NoError(t, err)

	err = registry.Parse("// regular comment", o)
	assert.NoError(t, err)
}

func TestDefaultRegistry(t *testing.T) {
	// Verify that standard comments are registered
	o := openapi.NewOpenAPIBuilder().Build()

	// Test servers.url
	err := comments.DefaultRegistry.Parse("// @servers.url https://example.com", o)
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com", o.Spec.Servers[0].Spec.URL)

	// Test servers.description
	err = comments.DefaultRegistry.Parse("// @servers.description My Server", o)
	assert.NoError(t, err)
	assert.Equal(t, "My Server", o.Spec.Servers[0].Spec.Description)

	// Test host (should error)
	err = comments.DefaultRegistry.Parse("// @host example.com", o)
	assert.Error(t, err)
	assert.EqualError(t, err, "@host is not supported use @servers.url instead")
}
