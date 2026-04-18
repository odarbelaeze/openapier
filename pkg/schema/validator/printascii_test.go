package validator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrintASCIITag_Tag(t *testing.T) {
	tag := PrintASCIITag{}
	assert.Equal(t, "printascii", tag.Tag())
}

func TestPrintASCIITag_Usage(t *testing.T) {
	tag := PrintASCIITag{}
	assert.Equal(t, "printascii", tag.Usage())
}

func TestPrintASCIITag_Parse(t *testing.T) {
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
			tag := PrintASCIITag{}
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
