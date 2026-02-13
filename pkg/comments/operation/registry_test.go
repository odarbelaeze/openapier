package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

type mockComment struct {
	tag string
}

func (m *mockComment) Tag() string {
	return m.tag
}

func (m *mockComment) Usage() string {
	return "usage"
}

func (m *mockComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI], op *openapi.Extendable[openapi.Operation]) error {
	if s.Spec.Paths == nil {
		s.Spec.Paths = openapi.NewPaths()
	}
	s.Spec.Paths.Spec.Add("/test", openapi.NewPathItemBuilder().Build())
	return nil
}

func TestRegistry_Parse_WithSpec(t *testing.T) {
	r := operation.NewRegistry()
	comment := &mockComment{tag: "test.tag"}
	r.Register(comment)

	spec := openapi.NewOpenAPIBuilder().Build()
	op := openapi.NewOperationBuilder().Build()

	err := r.Parse("// @test.tag content", spec, op)
	require.NoError(t, err)

	assert.Contains(t, spec.Spec.Paths.Spec.Paths, "/test")
}
