package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCIDRTag(t *testing.T) {
	tag := validator.CIDRTag{}

	t.Run("Tag", func(t *testing.T) {
		assert.Equal(t, "cidr", tag.Tag())
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
		})
	})
}
