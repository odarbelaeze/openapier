package options

import "github.com/sv-tools/openapi"

type SchemaOption func(*openapi.SchemaBuilder)

func WithRequired() SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Required()
	}
}

func WithDescription(description string) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Description(description)
	}
}

func WithExample(example any) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Examples(example)
	}
}

func WithMinLength(minLength int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.MinLength(minLength)
	}
}

func WithMaxLength(maxLength int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.MaxLength(maxLength)
	}
}

func WithMinItems(minItems int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.MinItems(minItems)
	}
}

func WithMaxItems(maxItems int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.MaxItems(maxItems)
	}
}

func WithMinProperties(minProperties int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.MinProperties(minProperties)
	}
}

func WithMaxProperties(maxProperties int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.MaxProperties(maxProperties)
	}
}

func WithMinimum(minimum int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Minimum(minimum)
	}
}

func WithMaximum(maximum int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Maximum(maximum)
	}
}

func WithExclusiveMinimum(minimum int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.ExclusiveMinimum(minimum)
	}
}

func WithExclusiveMaximum(maximum int) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.ExclusiveMaximum(maximum)
	}
}

func WithFormat(format string) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Format(format)
	}
}

func WithEnum(values ...any) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Enum(values...)
	}
}

func WithPattern(pattern string) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Pattern(pattern)
	}
}

func WithUniqueItems() SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.UniqueItems(true)
	}
}

func WithNot(ref *openapi.RefOrSpec[openapi.Schema]) SchemaOption {
	return func(sb *openapi.SchemaBuilder) {
		sb.Not(ref)
	}
}
