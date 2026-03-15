package operation

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewResponseComment())
}

// ResponseComment adds a response to an operation.
type ResponseComment struct{}

func NewResponseComment() *ResponseComment {
	return &ResponseComment{}
}

func (c *ResponseComment) Tag() string {
	return "response"
}

func (c *ResponseComment) Usage() string {
	return "@response <status_code> <content_type> <type> [description]"
}

func (c *ResponseComment) ParseInto(content string, f *ast.File, op *Operation) error {
	fields := strings.Fields(content)
	if len(fields) < 3 {
		return fmt.Errorf("invalid @response format, expected: %s", c.Usage())
	}

	statusCode := fields[0]
	contentType := fields[1]
	typ := fields[2]
	var description string
	if len(fields) > 3 {
		// Extract the rest of the string as the description.
		// Find the index of the third field (after statusCode, contentType and typ).
		description = strings.TrimSpace(content[strings.Index(content, typ)+len(typ):])
	}

	s, err := op.Resolver.Resolve(typ, f)
	if err != nil {
		return fmt.Errorf("failed to resolve type %q: %w", typ, err)
	}

	mediaType := openapi.NewMediaTypeBuilder().Schema(s).Build()
	response := openapi.NewResponseBuilder().
		Description(description).
		AddContent(contentType, mediaType).
		Build()
	op.ResponsesBuilder.AddResponse(statusCode, response)

	return nil
}
