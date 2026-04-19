package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLenTag(t *testing.T) {
	tag := validator.LenTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "len", tag.Tag())
	})

	t.Run("Usage", func(t *testing.T) {
		assert.Equal(t, "len=x", tag.Usage())
	})

	t.Run("Parse", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			opts, err := tag.Parse("3", "object")
			require.NoError(t, err)
			assert.Len(t, opts, 2)
		})

		t.Run("failure", func(t *testing.T) {
			_, err := tag.Parse("abc", "object")
			require.Error(t, err)
		})

		t.Run("integers are not supported", func(t *testing.T) {
			_, err := tag.Parse("3", "integer")
			require.Error(t, err)
		})

		t.Run("numbers are not supported", func(t *testing.T) {
			_, err := tag.Parse("3", "number")
			require.Error(t, err)
		})
	})
}
