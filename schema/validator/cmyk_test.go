package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCMYKTag_Tag(t *testing.T) {
	tag := validator.CMYKTag{}
	assert.Equal(t, "cmyk", tag.Tag())
}

func TestCMYKTag_Usage(t *testing.T) {
	tag := validator.CMYKTag{}
	assert.Equal(t, "cmyk", tag.Usage())
}

func TestCMYKTag_Parse(t *testing.T) {
	tests := []struct {
		name      string
		as        string
		value     string
		expectErr bool
	}{
		{
			name:      "supported type string",
			as:        "string",
			expectErr: false,
		},
		{
			name:      "unsupported type integer",
			as:        "integer",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := validator.CMYKTag{}
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
