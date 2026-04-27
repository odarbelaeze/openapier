package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNETag_Tag(t *testing.T) {
	tag := validator.NETag{}
	assert.Equal(t, "ne", tag.Tag())
}

func TestNETag_Usage(t *testing.T) {
	tag := validator.NETag{}
	assert.Equal(t, "ne=value", tag.Usage())
}

func TestNETag_Parse(t *testing.T) {
	tests := []struct {
		name      string
		as        string
		value     string
		expectErr bool
	}{
		{
			name:      "supported type string",
			as:        "string",
			value:     "abc",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := validator.NETag{}
			opts, err := tag.Parse(tt.value, tt.as)

			if tt.expectErr {
				require.Error(t, err)
				assert.Nil(t, opts)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, opts)
				assert.Len(t, opts, 1)
			}
		})
	}
}
