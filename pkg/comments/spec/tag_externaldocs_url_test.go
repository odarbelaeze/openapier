package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagExternalDocsURLComment_Tag(t *testing.T) {
	comment := &spec.TagExternalDocsURLComment{}
	assert.Equal(t, "tag.externalDocs.url", comment.Tag())
}

func TestTagExternalDocsURLComment_Usage(t *testing.T) {
	comment := &spec.TagExternalDocsURLComment{}
	assert.Equal(t, "@tag.externalDocs.url <url>", comment.Usage())
}

func TestTagExternalDocsURLComment_ParseInto(t *testing.T) {
	tests := []struct {
		name          string
		line          string
		initialTags   []*openapi.Extendable[openapi.Tag]
		expectedURL   string
		expectedError string
	}{
		{
			name: "valid external docs url with preceding tag",
			line: "https://example.com",
			initialTags: []*openapi.Extendable[openapi.Tag]{
				openapi.NewTagBuilder().Name("pet").Build(),
			},
			expectedURL: "https://example.com",
		},
		{
			name: "empty external docs url with preceding tag",
			line: "",
			initialTags: []*openapi.Extendable[openapi.Tag]{
				openapi.NewTagBuilder().Name("pet").Build(),
			},
			expectedURL: "",
		},
		{
			name:          "error when no tags exist",
			line:          "https://example.com",
			initialTags:   nil,
			expectedError: "a @tag.externalDocs.url comment requires a preceding @tag.name comment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.TagExternalDocsURLComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			o.Spec.Tags = tt.initialTags

			err := comment.ParseInto(tt.line, o)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.Len(t, o.Spec.Tags, 1)
				require.NotNil(t, o.Spec.Tags[0].Spec.ExternalDocs)
				assert.Equal(t, tt.expectedURL, o.Spec.Tags[0].Spec.ExternalDocs.Spec.URL)
			}
		})
	}
}

func TestTagExternalDocsURLComment_ParseInto_MultipleTags(t *testing.T) {
	comment := &spec.TagExternalDocsURLComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	o.Spec.Tags = []*openapi.Extendable[openapi.Tag]{
		openapi.NewTagBuilder().Name("tag1").Description("first").Build(),
		openapi.NewTagBuilder().Name("tag2").Build(),
	}

	err := comment.ParseInto("https://example.com/2", o)
	require.NoError(t, err)

	require.Len(t, o.Spec.Tags, 2)
	assert.Nil(t, o.Spec.Tags[0].Spec.ExternalDocs)
	require.NotNil(t, o.Spec.Tags[1].Spec.ExternalDocs)
	assert.Equal(t, "https://example.com/2", o.Spec.Tags[1].Spec.ExternalDocs.Spec.URL)
}
