package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoContactURLComment_Tag(t *testing.T) {
	comment := &spec.InfoContactURLComment{}
	assert.Equal(t, "info.contact.url", comment.Tag())
}

func TestInfoContactURLComment_Usage(t *testing.T) {
	comment := &spec.InfoContactURLComment{}
	assert.Equal(t, "@info.contact.url <url>", comment.Usage())
}

func TestInfoContactURLComment_ParseInto(t *testing.T) {
	tests := []struct {
		name               string
		line               string
		expectedContactURL string
	}{
		{
			name:               "valid url",
			line:               "https://www.example.com/support",
			expectedContactURL: "https://www.example.com/support",
		},
		{
			name:               "empty url",
			line:               "",
			expectedContactURL: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoContactURLComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			require.NotNil(t, o.Spec.Info.Spec.Contact)
			assert.Equal(t, tt.expectedContactURL, o.Spec.Info.Spec.Contact.Spec.URL)
		})
	}
}
