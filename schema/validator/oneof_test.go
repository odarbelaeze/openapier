package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/options"
	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestOneOfTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    string
		as       string
		expected []options.SchemaOption
		wantErr  bool
	}{
		{
			name:  "oneof strings",
			value: "red green blue",
			as:    "string",
			expected: []options.SchemaOption{
				options.WithEnum("red", "green", "blue"),
			},
		},
		{
			name:  "oneof integers",
			value: "5 10 15",
			as:    "integer",
			expected: []options.SchemaOption{
				options.WithEnum(5, 10, 15),
			},
		},
		{
			name:  "oneof numbers",
			value: "1.1 2.2 3.3",
			as:    "number",
			expected: []options.SchemaOption{
				options.WithEnum(1.1, 2.2, 3.3),
			},
		},
		{
			name:  "oneof booleans",
			value: "true false",
			as:    "boolean",
			expected: []options.SchemaOption{
				options.WithEnum(true, false),
			},
		},
		{
			name:    "unsupported type",
			value:   "foo bar",
			as:      "array",
			wantErr: true,
		},
		{
			name:    "invalid integer",
			value:   "1 foo 3",
			as:      "integer",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := validator.OneOfTag{}
			got, err := tag.Parse(tt.value, tt.as)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, got, len(tt.expected))

			// Validate by applying to a builder
			gotBuilder := openapi.NewSchemaBuilder()
			for _, opt := range got {
				opt(gotBuilder)
			}
			gotSchema := gotBuilder.Build()

			expectedBuilder := openapi.NewSchemaBuilder()
			for _, opt := range tt.expected {
				opt(expectedBuilder)
			}
			expectedSchema := expectedBuilder.Build()

			assert.Equal(t, expectedSchema.Spec.Enum, gotSchema.Spec.Enum)
		})
	}
}
