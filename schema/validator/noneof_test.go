package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/options"
	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestNoneOfTag_Tag(t *testing.T) {
	tag := validator.NoneOfTag{}
	assert.Equal(t, "noneof", tag.Tag())
}

func TestNoneOfTag_Usage(t *testing.T) {
	tag := validator.NoneOfTag{}
	assert.Equal(t, "noneof=value1 value2", tag.Usage())
}

func TestNoneOfTag_Parse(t *testing.T) {
	tests := []struct {
		name      string
		as        string
		value     string
		expected  []options.SchemaOption
		expectErr bool
	}{
		{
			name:  "noneof strings",
			as:    "string",
			value: "abc def",
			expected: []options.SchemaOption{
				options.WithNot(openapi.NewSchemaBuilder().Enum("abc", "def").Build()),
			},
			expectErr: false,
		},
		{
			name:  "noneof integers",
			as:    "integer",
			value: "5 10 15",
			expected: []options.SchemaOption{
				options.WithNot(openapi.NewSchemaBuilder().Enum(5, 10, 15).Build()),
			},
			expectErr: false,
		},
		{
			name:  "noneof numbers",
			as:    "number",
			value: "1.1 2.2 3.3",
			expected: []options.SchemaOption{
				options.WithNot(openapi.NewSchemaBuilder().Enum(1.1, 2.2, 3.3).Build()),
			},
			expectErr: false,
		},
		{
			name:  "noneof booleans",
			as:    "boolean",
			value: "true false",
			expected: []options.SchemaOption{
				options.WithNot(openapi.NewSchemaBuilder().Enum(true, false).Build()),
			},
			expectErr: false,
		},
		{
			name:      "unsupported type",
			as:        "array",
			value:     "foo bar",
			expectErr: true,
		},
		{
			name:      "invalid integer",
			as:        "integer",
			value:     "1 foo 3",
			expectErr: true,
		},
		{
			name:      "invalid number",
			as:        "number",
			value:     "1.1 foo 3.3",
			expectErr: true,
		},
		{
			name:      "invalid boolean",
			as:        "boolean",
			value:     "true foo",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := validator.NoneOfTag{}
			got, err := tag.Parse(tt.value, tt.as)

			if tt.expectErr {
				require.Error(t, err)
				assert.Nil(t, got)
				return
			}

			require.NoError(t, err)
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

			assert.Equal(t, expectedSchema.Spec.Not.Spec.Enum, gotSchema.Spec.Not.Spec.Enum)
		})
	}
}
