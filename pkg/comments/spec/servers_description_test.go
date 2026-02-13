package spec_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/spec"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestServerDescription_ParseInto_NoServer(t *testing.T) {
	serverDescription := spec.NewServersDescriptionComment()
	openapi := openapi.NewOpenAPIBuilder().Build()
	err := serverDescription.ParseInto("Test description", openapi)
	assert.Error(t, err)
}

func TestServerDescription_ParseInto_WithServer(t *testing.T) {
	serverDescription := spec.NewServersDescriptionComment()
	openapi := openapi.NewOpenAPIBuilder().AddServers(openapi.NewServerBuilder().Build()).Build()
	err := serverDescription.ParseInto("Test description", openapi)
	assert.NoError(t, err)
	assert.Equal(t, openapi.Spec.Servers[0].Spec.Description, "Test description")
}

func TestServerDescription_Tag(t *testing.T) {
	serverDescription := spec.NewServersDescriptionComment()
	assert.Equal(t, "servers.description", serverDescription.Tag())
}

func TestServerDescription_Usage(t *testing.T) {
	serverDescription := spec.NewServersDescriptionComment()
	assert.Equal(t, "// @servers.description <description>", serverDescription.Usage())
}
