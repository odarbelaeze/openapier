package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagDocsDescriptionComment_ParseInto(t *testing.T) {
	c := spec.NewTagDocsDescriptionComment()
	s := openapi.NewOpenAPIBuilder().Build()

	// Should fail if no tags exist
	err := c.ParseInto("Docs description", s)
	require.Error(t, err)

	// Add a tag without ExternalDocs
	s.Spec.Tags = append(s.Spec.Tags, openapi.NewTagBuilder().Name("MyTag").Build())

	// Should fail if ExternalDocs doesn't exist
	err = c.ParseInto("Docs description", s)
	require.Error(t, err)

	// Add ExternalDocs
	s.Spec.Tags[0].Spec.ExternalDocs = openapi.NewExternalDocsBuilder().URL("http://example.com").Build()

	err = c.ParseInto("Docs description", s)
	require.NoError(t, err)

	assert.Equal(t, "Docs description", s.Spec.Tags[0].Spec.ExternalDocs.Spec.Description)
}

func TestTagDocsDescriptionComment_Tag(t *testing.T) {
	c := spec.NewTagDocsDescriptionComment()
	assert.Equal(t, "tag.docs.description", c.Tag())
}

func TestTagDocsDescriptionComment_Usage(t *testing.T) {
	c := spec.NewTagDocsDescriptionComment()
	assert.Equal(t, "// @tag.docs.description <description>", c.Usage())
}
