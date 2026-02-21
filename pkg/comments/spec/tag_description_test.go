package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagDescriptionComment_Tag(t *testing.T) {
	comment := &spec.TagDescriptionComment{}
	assert.Equal(t, "tag.description", comment.Tag())
}

func TestTagDescriptionComment_Usage(t *testing.T) {
	comment := &spec.TagDescriptionComment{}
	assert.Equal(t, "@tag.description <description>", comment.Usage())
}

func TestTagDescriptionComment_ParseInto(t *testing.T) {
	tests := []struct {
		name                string
		line                string
		initialTags         []*openapi.Extendable[openapi.Tag]
		expectedDescription string
		expectedError       string
	}{
		{
			name: "valid description with preceding tag",
			line: "Pets operations",
			initialTags: []*openapi.Extendable[openapi.Tag]{
				openapi.NewTagBuilder().Name("pet").Build(),
			},
			expectedDescription: "Pets operations",
		},
		{
			name: "empty description with preceding tag",
			line: "",
			initialTags: []*openapi.Extendable[openapi.Tag]{
				openapi.NewTagBuilder().Name("pet").Build(),
			},
			expectedDescription: "",
		},
		{
			name:          "error when no tags exist",
			line:          "Pets operations",
			initialTags:   nil,
			expectedError: "a @tag.description comment requires a preceding @tag.name comment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := &spec.TagDescriptionComment{}
			o := openapi.NewOpenAPIBuilder().Build()
			o.Spec.Tags = tt.initialTags

			err := comment.ParseInto(tt.line, o)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.Len(t, o.Spec.Tags, 1)
				assert.Equal(t, tt.expectedDescription, o.Spec.Tags[0].Spec.Description)
			}
		})
	}
}

func TestTagDescriptionComment_ParseInto_MultipleTags(t *testing.T) {
	comment := &spec.TagDescriptionComment{}
	o := openapi.NewOpenAPIBuilder().Build()

	o.Spec.Tags = []*openapi.Extendable[openapi.Tag]{
		openapi.NewTagBuilder().Name("tag1").Description("first").Build(),
		openapi.NewTagBuilder().Name("tag2").Build(),
	}

	err := comment.ParseInto("second", o)
	require.NoError(t, err)

	require.Len(t, o.Spec.Tags, 2)
	assert.Equal(t, "first", o.Spec.Tags[0].Spec.Description)
	assert.Equal(t, "second", o.Spec.Tags[1].Spec.Description)
}
