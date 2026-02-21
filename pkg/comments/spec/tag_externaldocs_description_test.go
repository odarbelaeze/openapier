package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagExternalDocsDescriptionComment_Tag(t *testing.T) {
	comment := &spec.TagExternalDocsDescriptionComment{}
	assert.Equal(t, "tag.externaldocs.description", comment.Tag())
}

func TestTagExternalDocsDescriptionComment_Usage(t *testing.T) {
	comment := &spec.TagExternalDocsDescriptionComment{}
	assert.Equal(t, "@tag.externalDocs.description <description>", comment.Usage())
}

func TestTagExternalDocsDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		line                string
		initialTags         []*openapi.Extendable[openapi.Tag]
		expectedDescription string
		expectedError       string
	}{
		{
			name: "valid external docs description with preceding tag",
			line: "Find more info here",
			initialTags: []*openapi.Extendable[openapi.Tag]{
				openapi.NewTagBuilder().Name("pet").Build(),
			},
			expectedDescription: "Find more info here",
		},
		{
			name: "empty external docs description with preceding tag",
			line: "",
			initialTags: []*openapi.Extendable[openapi.Tag]{
				openapi.NewTagBuilder().Name("pet").Build(),
			},
			expectedDescription: "",
		},
		{
			name:          "error when no tags exist",
			line:          "Find more info here",
			initialTags:   nil,
			expectedError: "a @tag.externalDocs.description comment requires a preceding @tag.name comment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.TagExternalDocsDescriptionComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			o.Spec.Tags = tt.initialTags

			err := comment.ParseInto(tt.line, o)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.Len(t, o.Spec.Tags, 1)
				require.NotNil(t, o.Spec.Tags[0].Spec.ExternalDocs)
				assert.Equal(t, tt.expectedDescription, o.Spec.Tags[0].Spec.ExternalDocs.Spec.Description)
			}
		})
	}
}

func TestTagExternalDocsDescriptionComment_ParseInto_MultipleTags(t *testing.T) {
	comment := &spec.TagExternalDocsDescriptionComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	o.Spec.Tags = []*openapi.Extendable[openapi.Tag]{
		openapi.NewTagBuilder().Name("tag1").Description("first").Build(),
		openapi.NewTagBuilder().Name("tag2").Build(),
	}

	err := comment.ParseInto("Docs here", o)
	require.NoError(t, err)

	require.Len(t, o.Spec.Tags, 2)
	assert.Nil(t, o.Spec.Tags[0].Spec.ExternalDocs)
	require.NotNil(t, o.Spec.Tags[1].Spec.ExternalDocs)
	assert.Equal(t, "Docs here", o.Spec.Tags[1].Spec.ExternalDocs.Spec.Description)
}
