package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartsWithTag(t *testing.T) {
	tag := StartsWithTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "startswith", tag.Tag())
	})

	t.Run("Parse", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			opts, err := tag.Parse("prefix", "string")
			require.NoError(t, err)
			assert.Len(t, opts, 1)
		})

		t.Run("unsupported type", func(t *testing.T) {
			_, err := tag.Parse("prefix", "integer")
			require.Error(t, err)
		})
	})
}
