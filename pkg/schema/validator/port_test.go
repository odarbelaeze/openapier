package validator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPortTag_Tag(t *testing.T) {
	tag := PortTag{}
	assert.Equal(t, "port", tag.Tag())
}

func TestPortTag_Usage(t *testing.T) {
	tag := PortTag{}
	assert.Equal(t, "port", tag.Usage())
}

func TestPortTag_Parse(t *testing.T) {
	tests := []struct {
		name      string
		as        string
		value     string
		expectErr bool
	}{
		{
			name:      "supported type integer",
			as:        "integer",
			expectErr: false,
		},
		{
			name:      "supported type string",
			as:        "string",
			expectErr: false,
		},
		{
			name:      "unsupported type array",
			as:        "array",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := PortTag{}
			opts, err := tag.Parse(tt.value, tt.as)

			if tt.expectErr {
				require.Error(t, err)
				assert.Nil(t, opts)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, opts)
				if tt.as == "integer" || tt.as == "number" {
					assert.Len(t, opts, 2)
				} else {
					assert.Len(t, opts, 1)
				}
			}
		})
	}
}
