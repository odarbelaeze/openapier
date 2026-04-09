package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEqTag(t *testing.T) {
	tag := EqTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "eq", tag.Tag())
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
			{"boolean", "true", "boolean", 1, false},
			{"string", "abc", "string", 1, false},
			{"invalid integer", "abc", "integer", 0, true},
			{"unsupported type", "1", "array", 0, true},
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
