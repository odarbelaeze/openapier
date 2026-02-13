package spec

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewTagDescriptionMarkdownComment("."))
}

type tagDescriptionMarkdownComment struct {
	markdownFileDir string
}

// NewTagDescriptionMarkdownComment creates a new tagDescriptionMarkdownComment.
func NewTagDescriptionMarkdownComment(markdownFileDir string) *tagDescriptionMarkdownComment {
	return &tagDescriptionMarkdownComment{
		markdownFileDir: markdownFileDir,
	}
}

// ParseInto implements Comment.
func (t *tagDescriptionMarkdownComment) ParseInto(c string, s *openapi.Extendable[openapi.OpenAPI]) error {
	if len(s.Spec.Tags) == 0 {
		return fmt.Errorf("@tag.description.markdown must follow @tag.name")
	}
	tag := s.Spec.Tags[len(s.Spec.Tags)-1]

	commentInfo, err := getMarkdownForTag(c, t.markdownFileDir)
	if err != nil {
		return err
	}

	tag.Spec.Description = string(commentInfo)
	return nil
}

// Tag implements Comment.
func (t *tagDescriptionMarkdownComment) Tag() string {
	return "tag.description.markdown"
}

// Usage implements Comment.
func (t *tagDescriptionMarkdownComment) Usage() string {
	return `// @tag.description.markdown <filename>`
}

func getMarkdownForTag(filename, dir string) ([]byte, error) {
	path := filepath.Join(dir, filename)
	return os.ReadFile(path)
}
