package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestTitle_ParseInto(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected string
	}{
		{
			name:     "simple title",
			title:    "My API",
			expected: "My API",
		},
		{
			name:     "title with special chars",
			title:    "My API - v1",
			expected: "My API - v1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := spec.NewTitleComment()
			o := openapi.NewOpenAPIBuilder().Build()

			err := comment.ParseInto(tt.title, o)
			assert.NoError(t, err)

			assert.NotNil(t, o.Spec.Info)
			assert.Equal(t, tt.expected, o.Spec.Info.Spec.Title)
		})
	}
}

func TestTitle_Tag(t *testing.T) {
	comment := spec.NewTitleComment()
	assert.Equal(t, "title", comment.Tag())
}

func TestTitle_Usage(t *testing.T) {
	comment := spec.NewTitleComment()
	assert.NotEmpty(t, comment.Usage())
}
