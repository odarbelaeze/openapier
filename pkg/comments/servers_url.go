package comments

import (
	"regexp"

	"github.com/sv-tools/openapi"
)

var (
	serversURLPattern       = regexp.MustCompile(`\{([^}]+)\}`)
	serversVariablesPattern = regexp.MustCompile(`^(\w+)\s+(.+)$`)

	_ Comment = &serversURLComment{}
)

type serversURLComment struct{}

// NewServersURLComment creates a new serversURLComment.
func NewServersURLComment() *serversURLComment {
	return &serversURLComment{}
}

// ParseInto implements Comment.
func (*serversURLComment) ParseInto(c string, s openapi.OpenAPI) error {
	builder := openapi.NewServerBuilder().URL(c)
	matches := serversURLPattern.FindAllStringSubmatch(c, -1)
	for _, match := range matches {
		name := match[1]
		builder.AddVariable(name, openapi.NewServerVariableBuilder().Build())
	}
	s.Servers = append(s.Servers, builder.Build())
	return nil
}

// Tag implements Comment.
func (s *serversURLComment) Tag() string {
	return "servers.url"
}

// Usage implements Comment.
func (s *serversURLComment) Usage() string {
	return `// @servers.url <url>`
}
