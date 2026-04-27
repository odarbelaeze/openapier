package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestSecuritySchemeBearerFormatComment_Tag(t *testing.T) {
	comment := &spec.SecuritySchemeBearerFormatComment{}
	assert.Equal(t, "securityScheme.bearerFormat", comment.Tag())
}

func TestSecuritySchemeBearerFormatComment_Usage(t *testing.T) {
	comment := &spec.SecuritySchemeBearerFormatComment{}
	assert.Equal(t, "@securityScheme.bearerFormat <securitySchemeName> <format>", comment.Usage())
}

func TestSecuritySchemeBearerFormatComment_ParseInto(t *testing.T) {
	o := openapi.NewOpenAPIBuilder().Build()
	schemeComment := &spec.SecuritySchemeComment{}
	bearerFormatComment := &spec.SecuritySchemeBearerFormatComment{}

	err := schemeComment.ParseInto("myAuth http", o)
	require.NoError(t, err)

	err = bearerFormatComment.ParseInto("myAuth JWT", o)
	require.NoError(t, err)

	assert.Equal(t, "JWT", o.Spec.Components.Spec.SecuritySchemes["myAuth"].Spec.Spec.BearerFormat)
}

func TestSecuritySchemeBearerFormatComment_ParseInto_NotFound(t *testing.T) {
	o := openapi.NewOpenAPIBuilder().Build()
	bearerFormatComment := &spec.SecuritySchemeBearerFormatComment{}

	err := bearerFormatComment.ParseInto("nonExistent JWT", o)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot add security scheme bearer format without a preceding @securityScheme")
}
