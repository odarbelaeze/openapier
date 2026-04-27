package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaxTag(t *testing.T) {
	tag := validator.MaxTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "max", tag.Tag())
	})

	t.Run("Parse", func(t *testing.T) {
		tests := []struct {
			name     string
			value    string
			as       string
			expected int
			wantErr  bool
		}{
			{"integer", "100", "integer", 1, false},
			{"number", "50.5", "number", 1, false},
			{"string", "10", "string", 1, false},
			{"array", "5", "array", 1, false},
			{"object", "3", "object", 1, false},
			{"invalid value", "xyz", "integer", 0, true},
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
