package validator

import (
	"fmt"
	"strings"

	"github.com/odarbelaeze/openapier/schema/options"
)

var defaultRegistry = NewRegistry()

type Registry interface {
	Register(tag Tag)
	Parse(string, string) ([]options.SchemaOption, error)
	Validators() []Tag
}

type registry struct {
	tags map[string]Tag
}

func Default() Registry {
	return defaultRegistry
}

func NewRegistry() Registry {
	return &registry{
		tags: make(map[string]Tag),
	}
}

func (r *registry) Register(tag Tag) {
	r.tags[tag.Tag()] = tag
}

func (r *registry) Validators() []Tag {
	validators := make([]Tag, 0, len(r.tags))
	for _, v := range r.tags {
		validators = append(validators, v)
	}
	return validators
}

func (r *registry) Parse(tagValue string, as string) ([]options.SchemaOption, error) {
	opts := make([]options.SchemaOption, 0)
	for tag := range strings.SplitSeq(tagValue, ",") {
		parts := strings.Split(tag, "=")
		if len(parts) > 2 {
			return nil, fmt.Errorf("invalid validator tag: %s", tag)
		}
		tagName := parts[0]
		var value string
		if len(parts) == 2 {
			value = parts[1]
		}
		if t, ok := r.tags[tagName]; ok {
			tagOpts, err := t.Parse(value, as)
			if err != nil {
				return nil, fmt.Errorf("failed to parse tag %s: %w", tagName, err)
			}
			opts = append(opts, tagOpts...)
		}
	}
	return opts, nil
}
