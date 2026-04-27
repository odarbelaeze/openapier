package operation

import (
	"go/ast"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewExternalDocsURLComment())
}

// ExternalDocsURLComment sets the URL for an operation's external documentation.
type ExternalDocsURLComment struct{}

func NewExternalDocsURLComment() *ExternalDocsURLComment {
	return &ExternalDocsURLComment{}
}

func (c *ExternalDocsURLComment) Tag() string {
	return "externaldocs.url"
}

func (c *ExternalDocsURLComment) Usage() string {
	return "@externalDocs.url <url>"
}

func (c *ExternalDocsURLComment) ParseInto(content string, f *ast.File, op *Operation) error {
	url := strings.TrimSpace(content)
	if url != "" {
		if op.Builder.Build().Spec.ExternalDocs == nil {
			op.Builder.ExternalDocs(openapi.NewExternalDocsBuilder().URL(url).Build())
		} else {
			op.Builder.Build().Spec.ExternalDocs.Spec.URL = url
		}
	}
	return nil
}
