package operation

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/odarbelaeze/openapier/schema/resolver"
	"github.com/sv-tools/openapi"
)

// Operation holds an operation being built
type Operation struct {
	Builder          *openapi.OperationBuilder
	ResponsesBuilder *openapi.ResponsesBuilder
	Routes           Routes
	Resolver         resolver.Resolver
}

// NewOperation builds a new operation
func NewOperation(resolver resolver.Resolver) *Operation {
	return &Operation{
		Builder:          openapi.NewOperationBuilder(),
		ResponsesBuilder: openapi.NewResponsesBuilder(),
		Routes:           make(Routes, 0),
		Resolver:         resolver,
	}
}

// Attach attaches the operation to the given openapi spec.
func (o *Operation) Attach(spec *openapi.Extendable[openapi.OpenAPI]) error {
	operation := o.Builder.Build()
	responses := o.ResponsesBuilder.Build()
	if responses.Spec.Spec.Response != nil {
		operation.Spec.Responses = responses.Spec
	}
	summary := o.Routes.Summarize()
	for path, methods := range summary {
		if spec.Spec.Paths == nil {
			spec.Spec.Paths = openapi.NewPaths()
		}
		if _, ok := spec.Spec.Paths.Spec.Paths[path]; !ok {
			spec.Spec.Paths.Spec.Add(path, openapi.NewPathItemBuilder().Build())
		}
		pathItem := spec.Spec.Paths.Spec.Paths[path]
		for _, method := range methods {
			switch strings.ToUpper(method) {
			case http.MethodGet:
				pathItem.Spec.Spec.Get = operation
			case http.MethodHead:
				pathItem.Spec.Spec.Head = operation
			case http.MethodPost:
				pathItem.Spec.Spec.Post = operation
			case http.MethodPut:
				pathItem.Spec.Spec.Put = operation
			case http.MethodPatch:
				pathItem.Spec.Spec.Patch = operation
			case http.MethodDelete:
				pathItem.Spec.Spec.Delete = operation
			case http.MethodOptions:
				pathItem.Spec.Spec.Options = operation
			case http.MethodTrace:
				pathItem.Spec.Spec.Trace = operation
			default:
				return fmt.Errorf("unsupported method: %s", method)
			}
		}
	}
	return nil
}
