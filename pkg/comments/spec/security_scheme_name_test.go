package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestSecuritySchemeNameComment_Tag(t *testing.T) {
	comment := &spec.SecuritySchemeNameComment{}
	assert.Equal(t, "securityScheme.name", comment.Tag())
}

func TestSecuritySchemeNameComment_Usage(t *testing.T) {
	comment := &spec.SecuritySchemeNameComment{}
	assert.Equal(t, "@securityScheme.name <securitySchemeName> <name>", comment.Usage())
}

func TestSecuritySchemeNameComment_ParseInto(t *testing.T) {
	o := openapi.NewOpenAPIBuilder().Build()
	schemeComment := &spec.SecuritySchemeComment{}
	nameComment := &spec.SecuritySchemeNameComment{}

	err := schemeComment.ParseInto("myAuth apiKey", o)
	require.NoError(t, err)

	err = nameComment.ParseInto("myAuth X-API-Key", o)
	require.NoError(t, err)

	assert.Equal(t, "X-API-Key", o.Spec.Components.Spec.SecuritySchemes["myAuth"].Spec.Spec.Name)
}

func TestSecuritySchemeNameComment_ParseInto_NotFound(t *testing.T) {
	o := openapi.NewOpenAPIBuilder().Build()
	nameComment := &spec.SecuritySchemeNameComment{}

	err := nameComment.ParseInto("nonExistent param", o)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot add security scheme name without a preceding @securityScheme")
}
