package operation

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sv-tools/openapi"
)

// Operation holds an operation being built
type Operation struct {
	Builder          *openapi.OperationBuilder
	ResponsesBuilder *openapi.ResponsesBuilder
	Routes           Routes
}

// NewOperation builds a new operation
func NewOperation() *Operation {
	return &Operation{
		Builder:          openapi.NewOperationBuilder(),
		ResponsesBuilder: openapi.NewResponsesBuilder(),
		Routes:           make(Routes, 0),
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
		pathItemBuilder := openapi.NewPathItemBuilder()
		for _, method := range methods {
			switch strings.ToUpper(method) {
			case http.MethodGet:
				pathItemBuilder.Get(operation)
			case http.MethodHead:
				pathItemBuilder.Head(operation)
			case http.MethodPost:
				pathItemBuilder.Post(operation)
			case http.MethodPut:
				pathItemBuilder.Put(operation)
			case http.MethodPatch:
				pathItemBuilder.Patch(operation)
			case http.MethodDelete:
				pathItemBuilder.Delete(operation)
			case http.MethodOptions:
				pathItemBuilder.Options(operation)
			case http.MethodTrace:
				pathItemBuilder.Trace(operation)
			default:
				return fmt.Errorf("unsupported method: %s", method)
			}
		}
		if spec.Spec.Paths == nil {
			spec.Spec.Paths = openapi.NewPaths()
		}
		spec.Spec.Paths.Spec.Add(path, pathItemBuilder.Build())
	}
	return nil
}
