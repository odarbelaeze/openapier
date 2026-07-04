package operation

import (
	"errors"
	"go/ast"
	"strings"

	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewExternalDocsDescriptionComment())
}

// ExternalDocsDescriptionComment sets the description for an operation's external documentation.
type ExternalDocsDescriptionComment struct{}

func NewExternalDocsDescriptionComment() *ExternalDocsDescriptionComment {
	return &ExternalDocsDescriptionComment{}
}

func (c *ExternalDocsDescriptionComment) Tag() string {
	return "externaldocs.description"
}

func (c *ExternalDocsDescriptionComment) Usage() string {
	return "@externalDocs.description <description>"
}

func (c *ExternalDocsDescriptionComment) ParseInto(content string, f *ast.File, op *Operation) error {
	desc := strings.TrimSpace(content)
	if desc == "" {
		return errors.New("externalDocs.description comment must have a description")
	}
	if op.Builder.Build().Spec.ExternalDocs == nil {
		op.Builder.ExternalDocs(openapi.NewExternalDocsBuilder().Description(desc).Build())
	} else {
		op.Builder.Build().Spec.ExternalDocs.Spec.Description = desc
	}
	return nil
}
