package comments_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagDocsURLComment_ParseInto(t *testing.T) {
	c := comments.NewTagDocsURLComment()
	s := openapi.NewOpenAPIBuilder().Build()

	// Should fail if no tags exist
	err := c.ParseInto("http://example.com", s)
	require.Error(t, err)

	// Add a tag
	s.Spec.Tags = append(s.Spec.Tags, openapi.NewTagBuilder().Name("MyTag").Build())

	err = c.ParseInto("http://example.com", s)
	require.NoError(t, err)

	assert.NotNil(t, s.Spec.Tags[0].Spec.ExternalDocs)
	assert.Equal(t, "http://example.com", s.Spec.Tags[0].Spec.ExternalDocs.Spec.URL)
}

func TestTagDocsURLComment_Tag(t *testing.T) {
	c := comments.NewTagDocsURLComment()
	assert.Equal(t, "tag.docs.url", c.Tag())
}

func TestTagDocsURLComment_Usage(t *testing.T) {
	c := comments.NewTagDocsURLComment()
	assert.Equal(t, "// @tag.docs.url <url>", c.Usage())
}
