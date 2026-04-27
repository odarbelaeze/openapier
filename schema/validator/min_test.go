package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMinTag(t *testing.T) {
	tag := validator.MinTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "min", tag.Tag())
	})

	t.Run("Parse", func(t *testing.T) {
		tests := []struct {
			name     string
			value    string
			as       string
			expected int
			wantErr  bool
		}{
			{"integer", "10", "integer", 1, false},
			{"number", "5.5", "number", 1, false},
			{"string", "3", "string", 1, false},
			{"array", "2", "array", 1, false},
			{"object", "1", "object", 1, false},
			{"invalid value", "abc", "integer", 0, true},
			{"unsupported type", "1", "boolean", 0, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				opts, err := tag.Parse(tt.value, tt.as)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Len(t, opts, tt.expected)
				}
			})
		}
	})
}
