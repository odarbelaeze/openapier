package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestSecuritySchemeSchemeComment_Tag(t *testing.T) {
	comment := &spec.SecuritySchemeSchemeComment{}
	assert.Equal(t, "securityScheme.scheme", comment.Tag())
}

func TestSecuritySchemeSchemeComment_Usage(t *testing.T) {
	comment := &spec.SecuritySchemeSchemeComment{}
	assert.Equal(t, "@securityScheme.scheme <securitySchemeName> <scheme>", comment.Usage())
}

func TestSecuritySchemeSchemeComment_ParseInto(t *testing.T) {
	o := openapi.NewOpenAPIBuilder().Build()
	schemeComment := &spec.SecuritySchemeComment{}
	httpSchemeComment := &spec.SecuritySchemeSchemeComment{}

	err := schemeComment.ParseInto("myAuth http", o)
	require.NoError(t, err)

	err = httpSchemeComment.ParseInto("myAuth bearer", o)
	require.NoError(t, err)

	assert.Equal(t, "bearer", o.Spec.Components.Spec.SecuritySchemes["myAuth"].Spec.Spec.Scheme)
}

func TestSecuritySchemeSchemeComment_ParseInto_NotFound(t *testing.T) {
	o := openapi.NewOpenAPIBuilder().Build()
	httpSchemeComment := &spec.SecuritySchemeSchemeComment{}

	err := httpSchemeComment.ParseInto("nonExistent bearer", o)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot add security scheme HTTP scheme without a preceding @securityScheme")
}
