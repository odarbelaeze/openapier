package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequiredTag(t *testing.T) {
	tag := RequiredTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "required", tag.Tag())
	})

	t.Run("Usage", func(t *testing.T) {
		assert.Equal(t, "required", tag.Usage())
	})

	t.Run("Parse", func(t *testing.T) {
		opts, err := tag.Parse("", "")
		require.NoError(t, err)
		assert.Len(t, opts, 1)
		// We can't easily inspect the option function itself, but we've verified
		// it returns one option as expected.
	})
}
