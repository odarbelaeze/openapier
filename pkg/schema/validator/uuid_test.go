package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUUIDTag(t *testing.T) {
	tag := UUIDTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "uuid", tag.Tag())
	})

	t.Run("Parse", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			opts, err := tag.Parse("", "string")
			require.NoError(t, err)
			assert.Len(t, opts, 1)
		})

		t.Run("unsupported type", func(t *testing.T) {
			_, err := tag.Parse("", "integer")
			require.Error(t, err)
			assert.Contains(t, err.Error(), "uuid is not supported for integer")
		})
	})
}
