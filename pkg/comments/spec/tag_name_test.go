package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagNameComment_Tag(t *testing.T) {
	comment := &spec.TagNameComment{}
	assert.Equal(t, "tag.name", comment.Tag())
}

func TestTagNameComment_Usage(t *testing.T) {
	comment := &spec.TagNameComment{}
	assert.Equal(t, "@tag.name <name>", comment.Usage())
}

func TestTagNameComment_ParseInto(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		expectedName string
	}{
		{
			name:         "valid name",
			line:         "pet",
			expectedName: "pet",
		},
		{
			name:         "empty name",
			line:         "",
			expectedName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.TagNameComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			err := comment.ParseInto(tt.line, o)

			require.NoError(t, err)
			require.Len(t, o.Spec.Tags, 1)
			assert.Equal(t, tt.expectedName, o.Spec.Tags[0].Spec.Name)
		})
	}
}

func TestTagNameComment_ParseInto_MultipleTags(t *testing.T) {
	comment := &spec.TagNameComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	err := comment.ParseInto("tag1", o)
	require.NoError(t, err)

	err = comment.ParseInto("tag2", o)
	require.NoError(t, err)

	require.Len(t, o.Spec.Tags, 2)
	assert.Equal(t, "tag1", o.Spec.Tags[0].Spec.Name)
	assert.Equal(t, "tag2", o.Spec.Tags[1].Spec.Name)
}
