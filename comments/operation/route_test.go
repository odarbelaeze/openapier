package operation_test

import (
	"net/http"
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
)

func TestRoute_Summarize(t *testing.T) {
	tests := []struct {
		name     string
		routes   operation.Routes
		expected map[string][]string
	}{
		{
			name:     "empty routes",
			routes:   operation.Routes{},
			expected: map[string][]string{},
		},
		{
			name: "single route",
			routes: operation.Routes{
				{Path: "/test", Method: http.MethodGet},
			},
			expected: map[string][]string{
				"/test": {http.MethodGet},
			},
		},
		{
			name: "multiple routes same path",
			routes: operation.Routes{
				{Path: "/test", Method: http.MethodGet},
				{Path: "/test", Method: http.MethodPost},
			},
			expected: map[string][]string{
				"/test": {http.MethodGet, http.MethodPost},
			},
		},
		{
			name: "multiple routes different paths",
			routes: operation.Routes{
				{Path: "/test", Method: http.MethodGet},
				{Path: "/other", Method: http.MethodPost},
			},
			expected: map[string][]string{
				"/test":  {http.MethodGet},
				"/other": {http.MethodPost},
			},
		},
		{
			name: "duplicate routes",
			routes: operation.Routes{
				{Path: "/test", Method: http.MethodGet},
				{Path: "/test", Method: http.MethodGet},
			},
			expected: map[string][]string{
				"/test": {http.MethodGet, http.MethodGet},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summary := tt.routes.Summarize()
			assert.Equal(t, tt.expected, summary)
		})
	}
}
