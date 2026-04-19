package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
)

func TestCommentInterface(t *testing.T) {
	var _ spec.Comment
	assert.True(t, true)
}
