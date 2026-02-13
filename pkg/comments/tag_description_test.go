package comments_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagDescriptionComment_ParseInto(t *testing.T) {
	c := comments.NewTagDescriptionComment()
	s := openapi.NewOpenAPIBuilder().Build()

	// Should fail if no tags exist
	err := c.ParseInto("Some description", s)
	require.Error(t, err)

	// Add a tag
	s.Spec.Tags = append(s.Spec.Tags, openapi.NewTagBuilder().Name("MyTag").Build())

	err = c.ParseInto("Some description", s)
	require.NoError(t, err)

	assert.Equal(t, "Some description", s.Spec.Tags[0].Spec.Description)
}

func TestTagDescriptionComment_Tag(t *testing.T) {
	c := comments.NewTagDescriptionComment()
	assert.Equal(t, "tag.description", c.Tag())
}

func TestTagDescriptionComment_Usage(t *testing.T) {
	c := comments.NewTagDescriptionComment()
	assert.Equal(t, "// @tag.description <description>", c.Usage())
}
