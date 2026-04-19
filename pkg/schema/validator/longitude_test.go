package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLongitudeTag(t *testing.T) {
	tag := validator.LongitudeTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "longitude", tag.Tag())
	})

	t.Run("Parse", func(t *testing.T) {
		t.Run("success number", func(t *testing.T) {
			opts, err := tag.Parse("", "number")
			require.NoError(t, err)
			assert.Len(t, opts, 2)
		})

		t.Run("success integer", func(t *testing.T) {
			opts, err := tag.Parse("", "integer")
			require.NoError(t, err)
			assert.Len(t, opts, 2)
		})

		t.Run("unsupported type", func(t *testing.T) {
			_, err := tag.Parse("", "string")
			require.Error(t, err)
		})
	})
}
