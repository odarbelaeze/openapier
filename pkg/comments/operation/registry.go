package operation

import (
	"fmt"
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
	Parse(line string, op *Operation) error
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

func (r *standardRegistry) Parse(line string, op *Operation) error {
	matches := commentPattern.FindStringSubmatch(line)
	if len(matches) < 3 {
		return nil
	}
	tag := matches[1]
	content := matches[2]

	if handler, ok := r.comments[tag]; ok {
		return handler.ParseInto(content, op)
	}
	return fmt.Errorf("unknown comment tag: %s", tag)
}
