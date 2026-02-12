package comments_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestTagNameComment_ParseInto(t *testing.T) {
	c := comments.NewTagNameComment()
	s := openapi.NewOpenAPIBuilder().Build()

	err := c.ParseInto("MyTag", s)
	require.NoError(t, err)

	require.Len(t, s.Spec.Tags, 1)
	assert.Equal(t, "MyTag", s.Spec.Tags[0].Spec.Name)

	err = c.ParseInto("AnotherTag", s)
	require.NoError(t, err)

	require.Len(t, s.Spec.Tags, 2)
	assert.Equal(t, "AnotherTag", s.Spec.Tags[1].Spec.Name)
}
