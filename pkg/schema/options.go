package schema

import "github.com/sv-tools/openapi"

type SchemaOption func(*openapi.SchemaBuilder)

func WithDescription(description string) SchemaOption {
	return func(b *openapi.SchemaBuilder) {
		b.Description(description)
	}
}
