package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestSecuritySchemeInComment_Tag(t *testing.T) {
	comment := &spec.SecuritySchemeInComment{}
	assert.Equal(t, "securityScheme.in", comment.Tag())
}

func TestSecuritySchemeInComment_Usage(t *testing.T) {
	comment := &spec.SecuritySchemeInComment{}
	assert.Equal(t, "@securityScheme.in <securitySchemeName> <in>", comment.Usage())
}

func TestSecuritySchemeInComment_ParseInto(t *testing.T) {
	o := openapi.NewOpenAPIBuilder().Build()
	schemeComment := &spec.SecuritySchemeComment{}
	inComment := &spec.SecuritySchemeInComment{}

	err := schemeComment.ParseInto("myAuth apiKey", o)
	require.NoError(t, err)

	err = inComment.ParseInto("myAuth header", o)
	require.NoError(t, err)

	assert.Equal(t, "header", o.Spec.Components.Spec.SecuritySchemes["myAuth"].Spec.Spec.In)
}

func TestSecuritySchemeInComment_ParseInto_NotFound(t *testing.T) {
	o := openapi.NewOpenAPIBuilder().Build()
	inComment := &spec.SecuritySchemeInComment{}

	err := inComment.ParseInto("nonExistent param", o)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot add security scheme location without a preceding @securityScheme")
}
