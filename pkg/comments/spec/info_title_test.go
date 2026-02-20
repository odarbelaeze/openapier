package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestInfoTitleComment_Tag(t *testing.T) {
	comment := &spec.InfoTitleComment{}
	assert.Equal(t, "info.title", comment.Tag())
}

func TestInfoTitleComment_Usage(t *testing.T) {
	comment := &spec.InfoTitleComment{}
	assert.Equal(t, "@info.title <title>", comment.Usage())
}

func TestInfoTitleComment_ParseInto(t *testing.T) {
	tests := []struct {
		name          string
		line          string
		expectedTitle string
	}{
		{
			name:          "valid title",
			line:          "My API Title",
			expectedTitle: "My API Title",
		},
		{
			name:          "empty title",
			line:          "",
			expectedTitle: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.InfoTitleComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expectedTitle, o.Spec.Info.Spec.Title)
		})
	}
}

func TestInfoTitleComment_ParseInto_ExistingInfo(t *testing.T) {
	comment := &spec.InfoTitleComment{}
	o := openapi.NewOpenAPIBuilder().Build()
	o.Spec.Info = openapi.NewInfoBuilder().Build()
	o.Spec.Info.Spec.Version = "1.0.0" // pre-existing value

	err := comment.ParseInto("New Title", o)
	require.NoError(t, err)
	assert.Equal(t, "New Title", o.Spec.Info.Spec.Title)
	assert.Equal(t, "1.0.0", o.Spec.Info.Spec.Version)
}
