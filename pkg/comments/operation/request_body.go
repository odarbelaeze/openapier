package operation

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewRequestBodyComment())
}

// RequestBodyComment defines the request body for an operation.
type RequestBodyComment struct{}

func NewRequestBodyComment() *RequestBodyComment {
	return &RequestBodyComment{}
}

func (c *RequestBodyComment) Tag() string {
	return "requestBody"
}

func (c *RequestBodyComment) Usage() string {
	return "@requestBody <content_type> <type> [description]"
}

func (c *RequestBodyComment) ParseInto(content string, f *ast.File, op *Operation) error {
	fields := strings.Fields(content)
	if len(fields) < 2 {
		return fmt.Errorf("invalid @requestBody format, expected: %s", c.Usage())
	}

	contentType := fields[0]
	typ := fields[1]
	var description string
	if len(fields) > 2 {
		// Extract the rest of the string as the description.
		// Find the index of the second space (after contentType and typ).
		// We'll use strings.Join for safety or just find the third field start.
		description = strings.TrimSpace(content[strings.Index(content, typ)+len(typ):])
	}

	s, err := op.Resolver.Resolve(typ)
	if err != nil {
		return fmt.Errorf("failed to resolve type %q: %w", typ, err)
	}

	mediaType := openapi.NewMediaTypeBuilder().Schema(s).Build()
	requestBody := openapi.NewRequestBodyBuilder().
		Required(true).
		Description(description).
		AddContent(contentType, mediaType).
		Build()
	op.Builder.RequestBody(requestBody)

	return nil
}
