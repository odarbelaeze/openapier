package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoContactEmailComment_Tag(t *testing.T) {
	comment := &spec.InfoContactEmailComment{}
	assert.Equal(t, "info.contact.email", comment.Tag())
}

func TestInfoContactEmailComment_Usage(t *testing.T) {
	comment := &spec.InfoContactEmailComment{}
	assert.Equal(t, "@info.contact.email <email>", comment.Usage())
}

func TestInfoContactEmailComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                 string
		line                 string
		expectedContactEmail string
	}{
		{
			name:                 "valid email",
			line:                 "support@example.com",
			expectedContactEmail: "support@example.com",
		},
		{
			name:                 "empty email",
			line:                 "",
			expectedContactEmail: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoContactEmailComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			require.NotNil(t, o.Spec.Info.Spec.Contact)
			assert.Equal(t, tt.expectedContactEmail, o.Spec.Info.Spec.Contact.Spec.Email)
		})
	}
}
