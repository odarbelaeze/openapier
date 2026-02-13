package operation_test

import (
	"net/http"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
)

func TestRoute_Summarize(t *testing.T) {
	routes := operation.Routes{
		operation.Route{
			Path:   "/test",
			Method: http.MethodGet,
		},
	}
	summary := routes.Summarize()
	assert.Equal(t, map[string][]string{"/test": {http.MethodGet}}, summary)
}
