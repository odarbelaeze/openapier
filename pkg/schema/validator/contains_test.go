package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContainsTag_Tag(t *testing.T) {
	tag := validator.ContainsTag{}
	assert.Equal(t, "contains", tag.Tag())
}

func TestContainsTag_Usage(t *testing.T) {
	tag := validator.ContainsTag{}
	assert.Equal(t, "contains=value", tag.Usage())
}

func TestContainsTag_Parse(t *testing.T) {
	tests := []struct {
		name      string
		as        string
		value     string
		expectErr bool
	}{
		{
			name:      "supported type string",
			as:        "string",
			value:     "foo",
			expectErr: false,
		},
		{
			name:      "unsupported type integer",
			as:        "integer",
			value:     "foo",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := validator.ContainsTag{}
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
