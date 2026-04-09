package schema

import (
	"github.com/odarbelaeze/openapier/pkg/schema/options"
	"github.com/sv-tools/openapi"
)

// parseBasicType converts a string type representation to an OpenAPI schema object
func parseBasicType(t string, opts ...options.SchemaOption) *openapi.RefOrSpec[openapi.Schema] {
	b := openapi.NewSchemaBuilder()
	switch t {
	case "int", "int32", "uint", "uint32":
		b.AddType("integer").Format("int32")
	case "int64", "uint64":
		b.AddType("integer").Format("int64")
	case "float32":
		b.AddType("number").Format("float")
	case "float64":
		b.AddType("number").Format("double")
	case "bool":
		b.AddType("boolean")
	case "string":
		b.AddType("string")
	case "file":
		b.AddType("string").Format("binary")
	case "time.Time":
		b.AddType("string").Format("date-time")
	case "uuid.UUID":
		b.AddType("string").Format("uuid")
	case "any":
		// empty schema for any
	default:
		// this is not a basic type, let the caller figure it out
		return nil
	}
	for _, option := range opts {
		option(b)
	}
	return b.Build()
}
