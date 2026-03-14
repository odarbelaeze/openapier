package schema_test

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema"
	"github.com/stretchr/testify/assert"
)

func TestResolver_Resolve(t *testing.T) {
	r := schema.NewResolver()
	file1 := &ast.File{
		Name: &ast.Ident{Name: "mypackage"},
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: &ast.Ident{Name: "MyType"},
					},
				},
			},
		},
	}
	file2 := &ast.File{
		Name: &ast.Ident{Name: "otherpack"},
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"github.com/example/mypackage"`}},
		},
	}

	r.Collect("github.com/example/mypackage", file1)

	// Test local resolve
	ref, err := r.Resolve("MyType", file1)
	assert.NoError(t, err)
	assert.NotNil(t, ref)

	// Test import resolve
	ref, err = r.Resolve("mypackage.MyType", file2)
	assert.NoError(t, err)
	assert.NotNil(t, ref)
}
