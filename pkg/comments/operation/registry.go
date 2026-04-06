package operation

import (
	"go/ast"
	"log/slog"
	"regexp"
)

var (
	commentPattern  = regexp.MustCompile(`^//\s*@([\w\.]+)\s+(.*)$`)
	DefaultRegistry = NewRegistry()
)

// Register registers a comment to the default registry.
func Register(c Comment) {
	DefaultRegistry.Register(c)
}

// Registry allows to register and parse comments.
type Registry interface {
	Register(c Comment)
	Parse(line string, f *ast.File, op *Operation) error
	Comments() []Comment
}

type standardRegistry struct {
	comments map[string]Comment
}

// NewRegistry creates a new Registry.
func NewRegistry() Registry {
	return &standardRegistry{
		comments: make(map[string]Comment),
	}
}

func (r *standardRegistry) Register(c Comment) {
	r.comments[c.Tag()] = c
}

func (r *standardRegistry) Comments() []Comment {
	var comments []Comment
	for _, c := range r.comments {
		comments = append(comments, c)
	}
	return comments
}

func (r *standardRegistry) Parse(line string, f *ast.File, op *Operation) error {
	matches := commentPattern.FindStringSubmatch(line)
	if len(matches) < 3 {
		return nil
	}
	tag := matches[1]
	content := matches[2]

	if handler, ok := r.comments[tag]; ok {
		return handler.ParseInto(content, f, op)
	}
	slog.Warn("unknown operation tag", "tag", tag)
	return nil
}
