package operation_test

import (
	"net/http"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestNewOperation(t *testing.T) {
	op := operation.NewOperation()
	assert.NotNil(t, op)
	assert.NotNil(t, op.Builder)
	assert.NotNil(t, op.Routes)
	assert.Empty(t, op.Routes)
}

func TestOperation_Attach(t *testing.T) {
	tests := []struct {
		name          string
		routes        operation.Routes
		setupBuilder  func(*openapi.OperationBuilder)
		expectedPaths []string
		expectedCheck func(*testing.T, *openapi.OpenAPI)
		expectedError string
	}{
		{
			name: "attach single route get",
			routes: operation.Routes{
				{Path: "/test", Method: http.MethodGet},
			},
			setupBuilder: func(b *openapi.OperationBuilder) {
				b.Summary("Test Summary")
			},
			expectedPaths: []string{"/test"},
			expectedCheck: func(t *testing.T, spec *openapi.OpenAPI) {
				pathItemRef := spec.Paths.Spec.Paths["/test"]
				require.NotNil(t, pathItemRef)
				require.NotNil(t, pathItemRef.Spec)
				pathItem := pathItemRef.Spec.Spec
				require.NotNil(t, pathItem.Get)
				assert.Equal(t, "Test Summary", pathItem.Get.Spec.Summary)
			},
		},
		{
			name: "attach multiple routes different methods",
			routes: operation.Routes{
				{Path: "/test", Method: http.MethodGet},
				{Path: "/test", Method: http.MethodPost},
			},
			expectedPaths: []string{"/test"},
			expectedCheck: func(t *testing.T, spec *openapi.OpenAPI) {
				pathItemRef := spec.Paths.Spec.Paths["/test"]
				require.NotNil(t, pathItemRef)
				require.NotNil(t, pathItemRef.Spec)
				pathItem := pathItemRef.Spec.Spec
				assert.NotNil(t, pathItem.Get)
				assert.NotNil(t, pathItem.Post)
			},
		},
		{
			name: "attach multiple routes different paths",
			routes: operation.Routes{
				{Path: "/test1", Method: http.MethodGet},
				{Path: "/test2", Method: http.MethodPost},
			},
			expectedPaths: []string{"/test1", "/test2"},
			expectedCheck: func(t *testing.T, spec *openapi.OpenAPI) {
				pathItemRef1 := spec.Paths.Spec.Paths["/test1"]
				require.NotNil(t, pathItemRef1)
				assert.NotNil(t, pathItemRef1.Spec.Spec.Get)

				pathItemRef2 := spec.Paths.Spec.Paths["/test2"]
				require.NotNil(t, pathItemRef2)
				assert.NotNil(t, pathItemRef2.Spec.Spec.Post)
			},
		},
		{
			name: "unsupported method",
			routes: operation.Routes{
				{Path: "/test", Method: "INVALID"},
			},
			expectedError: "unsupported method: INVALID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := operation.NewOperation()
			op.Routes = tt.routes
			if tt.setupBuilder != nil {
				tt.setupBuilder(op.Builder)
			}

			specExt := openapi.NewOpenAPIBuilder().Build()
			// Ensure Paths is initialized as operation.go might not handle nil Paths
			if specExt.Spec.Paths == nil {
				specExt.Spec.Paths = openapi.NewPaths()
			}

			err := op.Attach(specExt)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				for _, path := range tt.expectedPaths {
					assert.Contains(t, specExt.Spec.Paths.Spec.Paths, path)
				}
				if tt.expectedCheck != nil {
					tt.expectedCheck(t, specExt.Spec)
				}
			}
		})
	}
}
