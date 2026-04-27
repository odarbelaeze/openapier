package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/options"
	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestLTTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    string
		as       string
		expected []options.SchemaOption
		wantErr  bool
	}{
		{
			name:  "lt integer",
			value: "5",
			as:    "integer",
			expected: []options.SchemaOption{
				options.WithExclusiveMaximum(5),
			},
		},
		{
			name:  "lt number",
			value: "5.5",
			as:    "number",
			expected: []options.SchemaOption{
				options.WithExclusiveMaximum(5),
			},
		},
		{
			name:  "lt string",
			value: "5",
			as:    "string",
			expected: []options.SchemaOption{
				options.WithMaxLength(4),
			},
		},
		{
			name:  "lt array",
			value: "3",
			as:    "array",
			expected: []options.SchemaOption{
				options.WithMaxItems(2),
			},
		},
		{
			name:  "lt object",
			value: "2",
			as:    "object",
			expected: []options.SchemaOption{
				options.WithMaxProperties(1),
			},
		},
		{
			name:    "invalid value",
			value:   "foo",
			as:      "integer",
			wantErr: true,
		},
		{
			name:    "unsupported type",
			value:   "5",
			as:      "boolean",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := validator.LTTag{}
			got, err := tag.Parse(tt.value, tt.as)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

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

			assert.Equal(t, expectedSchema.Spec.ExclusiveMaximum, gotSchema.Spec.ExclusiveMaximum)
			assert.Equal(t, expectedSchema.Spec.MaxLength, gotSchema.Spec.MaxLength)
			assert.Equal(t, expectedSchema.Spec.MaxItems, gotSchema.Spec.MaxItems)
			assert.Equal(t, expectedSchema.Spec.MaxProperties, gotSchema.Spec.MaxProperties)
		})
	}
}
