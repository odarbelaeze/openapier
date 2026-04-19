package operation_test

import (
	"errors"
	"go/ast"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
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

func (m *mockComment) ParseInto(c string, f *ast.File, op *operation.Operation) error {
	return m.err
}

func TestRegistry_Parse_KnownTag(t *testing.T) {
	registry := operation.NewRegistry()
	mock := &mockComment{tag: "mock", usage: "// @mock", err: nil}
	registry.Register(mock)

	op := &operation.Operation{}
	err := registry.Parse("// @mock foo", nil, op)
	assert.NoError(t, err)
}

func TestRegistry_Parse_ErrorFromHandler(t *testing.T) {
	registry := operation.NewRegistry()
	mock := &mockComment{tag: "mock", usage: "// @mock", err: errors.New("handler error")}
	registry.Register(mock)

	op := &operation.Operation{}
	err := registry.Parse("// @mock foo", nil, op)
	assert.Error(t, err)
	assert.EqualError(t, err, "handler error")
}

func TestRegistry_Parse_IgnoreNonComments(t *testing.T) {
	registry := operation.NewRegistry()
	op := &operation.Operation{}
	err := registry.Parse("not a comment", nil, op)
	assert.NoError(t, err)

	err = registry.Parse("// regular comment", nil, op)
	assert.NoError(t, err)
}

func TestRegistry_Comments(t *testing.T) {
	registry := operation.NewRegistry()
	mock1 := &mockComment{tag: "t1"}
	mock2 := &mockComment{tag: "t2"}
	registry.Register(mock1)
	registry.Register(mock2)

	comments := registry.Comments()
	assert.Len(t, comments, 2)
}

func TestDefault(t *testing.T) {
	assert.NotNil(t, operation.Default())
}
