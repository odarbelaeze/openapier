package comments_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestServerDescription_ParseInto_NoServer(t *testing.T) {
	serverDescription := comments.NewServersDescriptionComment()
	openapi := openapi.NewOpenAPIBuilder().Build()
	err := serverDescription.ParseInto("Test description", *openapi.Spec)
	assert.Error(t, err)
}

func TestServerDescription_ParseInto_WithServer(t *testing.T) {
	serverDescription := comments.NewServersDescriptionComment()
	openapi := openapi.NewOpenAPIBuilder().AddServers(openapi.NewServerBuilder().Build()).Build()
	err := serverDescription.ParseInto("Test description", *openapi.Spec)
	assert.NoError(t, err)
	assert.Equal(t, openapi.Spec.Servers[0].Spec.Description, "Test description")
}
