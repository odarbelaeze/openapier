package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoContactNameComment_Tag(t *testing.T) {
	comment := &spec.InfoContactNameComment{}
	assert.Equal(t, "info.contact.name", comment.Tag())
}

func TestInfoContactNameComment_Usage(t *testing.T) {
	comment := &spec.InfoContactNameComment{}
	assert.Equal(t, "@info.contact.name <name>", comment.Usage())
}

func TestInfoContactNameComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		line                string
		expectedContactName string
	}{
		{
			name:                "valid name",
			line:                "API Support",
			expectedContactName: "API Support",
		},
		{
			name:                "empty name",
			line:                "",
			expectedContactName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoContactNameComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			require.NotNil(t, o.Spec.Info.Spec.Contact)
			assert.Equal(t, tt.expectedContactName, o.Spec.Info.Spec.Contact.Spec.Name)
		})
	}
}
