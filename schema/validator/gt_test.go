package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/options"
	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestGTTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    string
		as       string
		expected []options.SchemaOption
		wantErr  bool
	}{
		{
			name:  "gt integer",
			value: "5",
			as:    "integer",
			expected: []options.SchemaOption{
				options.WithExclusiveMinimum(5),
			},
		},
		{
			name:  "gt number",
			value: "5.5",
			as:    "number",
			expected: []options.SchemaOption{
				options.WithExclusiveMinimum(5),
			},
		},
		{
			name:  "gt string",
			value: "5",
			as:    "string",
			expected: []options.SchemaOption{
				options.WithMinLength(6),
			},
		},
		{
			name:  "gt array",
			value: "3",
			as:    "array",
			expected: []options.SchemaOption{
				options.WithMinItems(4),
			},
		},
		{
			name:  "gt object",
			value: "2",
			as:    "object",
			expected: []options.SchemaOption{
				options.WithMinProperties(3),
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
			tag := validator.GTTag{}
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

			assert.Equal(t, expectedSchema.Spec.ExclusiveMinimum, gotSchema.Spec.ExclusiveMinimum)
			assert.Equal(t, expectedSchema.Spec.MinLength, gotSchema.Spec.MinLength)
			assert.Equal(t, expectedSchema.Spec.MinItems, gotSchema.Spec.MinItems)
			assert.Equal(t, expectedSchema.Spec.MinProperties, gotSchema.Spec.MinProperties)
		})
	}
}
