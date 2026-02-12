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
	s := &openapi.Extendable[openapi.OpenAPI]{
		Spec: &openapi.OpenAPI{},
	}

	// Should fail if no tags exist
	err = c.ParseInto(fileName, s)
	require.Error(t, err)

	// Add a tag
	s.Spec.Tags = append(s.Spec.Tags, &openapi.Extendable[openapi.Tag]{
		Spec: &openapi.Tag{Name: "MyTag"},
	})

	err = c.ParseInto(fileName, s)
	require.NoError(t, err)

	assert.Equal(t, fileContent, s.Spec.Tags[0].Spec.Description)
}
