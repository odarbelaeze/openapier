package schema

import "github.com/sv-tools/openapi"

type SchemaOption func(*openapi.SchemaBuilder)

func WithDescription(description string) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Description(description)
	}
}

func WithExample(example any) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Example(example)
	}
}
