package comments

import (
	"errors"
	"fmt"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewServersVariablesDescriptionMarkdownComment("."))
}

type serversVariablesDescriptionMarkdownComment struct {
	markdownFileDir string
}

// NewServersVariablesDescriptionMarkdownComment creates a new serversVariablesDescriptionMarkdownComment.
func NewServersVariablesDescriptionMarkdownComment(markdownFileDir string) *serversVariablesDescriptionMarkdownComment {
	return &serversVariablesDescriptionMarkdownComment{
		markdownFileDir: markdownFileDir,
	}
}

// ParseInto implements Comment.
func (s *serversVariablesDescriptionMarkdownComment) ParseInto(c string, o *openapi.Extendable[openapi.OpenAPI]) error {
	if len(o.Spec.Servers) == 0 {
		return errors.New("use @servers.url before you use @servers.variables.description.markdown")
	}
	server := o.Spec.Servers[len(o.Spec.Servers)-1]
	matches := serversVariablesPattern.FindStringSubmatch(c)
	if len(matches) > 0 {
		if server.Spec.Variables == nil {
			return errors.New("variables are not detected")
		}
		variable, ok := server.Spec.Variables[matches[1]]
		if !ok {
			return fmt.Errorf("variable %q is not detected", matches[1])
		}

		commentInfo, err := getMarkdownForTag(matches[2], s.markdownFileDir)
		if err != nil {
			return err
		}
		variable.Spec.Description = string(commentInfo)
	}
	return nil
}

// Tag implements Comment.
func (s *serversVariablesDescriptionMarkdownComment) Tag() string {
	return "servers.variables.description.markdown"
}

// Usage implements Comment.
func (s *serversVariablesDescriptionMarkdownComment) Usage() string {
	return `// @servers.variables.description.markdown <variable> <filename>`
}
