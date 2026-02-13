package comments_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagDescriptionMarkdownComment_ParseInto(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "desc.md"
	fileContent := "# Description\nThis is a markdown description."
	err := os.WriteFile(filepath.Join(tempDir, fileName), []byte(fileContent), 0644)
	require.NoError(t, err)

	c := comments.NewTagDescriptionMarkdownComment(tempDir)
	s := openapi.NewOpenAPIBuilder().Build()

	// Should fail if no tags exist
	err = c.ParseInto(fileName, s)
	require.Error(t, err)

	// Add a tag
	s.Spec.Tags = append(s.Spec.Tags, openapi.NewTagBuilder().Name("MyTag").Build())

	err = c.ParseInto(fileName, s)
	require.NoError(t, err)

	assert.Equal(t, fileContent, s.Spec.Tags[0].Spec.Description)

	// Should fail if file not found
	err = c.ParseInto("non-existent.md", s)
	require.Error(t, err)
}

func TestTagDescriptionMarkdownComment_Tag(t *testing.T) {
	c := comments.NewTagDescriptionMarkdownComment(".")
	assert.Equal(t, "tag.description.markdown", c.Tag())
}

func TestTagDescriptionMarkdownComment_Usage(t *testing.T) {
	c := comments.NewTagDescriptionMarkdownComment(".")
	assert.Equal(t, "// @tag.description.markdown <filename>", c.Usage())
}
