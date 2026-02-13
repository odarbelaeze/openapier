package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestHost_ParseInto(t *testing.T) {
	comment := spec.NewHostComment()
	o := openapi.NewOpenAPIBuilder().Build()
	err := comment.ParseInto("some host", o)
	assert.Error(t, err)
	assert.EqualError(t, err, "@host is not supported use @servers.url instead")
}

func TestHost_Tag(t *testing.T) {
	comment := spec.NewHostComment()
	assert.Equal(t, "host", comment.Tag())
}

func TestHost_Usage(t *testing.T) {
	comment := spec.NewHostComment()
	assert.Equal(t, "// @host <host>", comment.Usage())
}
