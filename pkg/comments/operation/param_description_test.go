package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/schema/resolver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestParamDescriptionComment(t *testing.T) {
	comment := operation.NewParamDescriptionComment()

	assert.Equal(t, "param.description", comment.Tag())
	assert.Equal(t, "@param.description <name> <description>", comment.Usage())

	tests := []struct {
		name        string
		content     string
		setupOp     func() *operation.Operation
		expectError bool
		validate    func(t *testing.T, op *openapi.Operation)
	}{
		{
			name:    "valid param description",
			content: "id the unique identifier of the item",
			setupOp: func() *operation.Operation {
				op := operation.NewOperation(resolver.NewResolver(nil, resolver.NewSchemaBuilder))
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
				return op
			},
			validate: func(t *testing.T, op *openapi.Operation) {
				require.Len(t, op.Parameters, 1)
				param := op.Parameters[0].Spec.Spec
				assert.Equal(t, "id", param.Name)
				assert.Equal(t, "the unique identifier of the item", param.Description)
			},
		},
		{
			name:    "parameter not found",
			content: "limit the maximum number of items",
			setupOp: func() *operation.Operation {
				op := operation.NewOperation(resolver.NewResolver(nil, resolver.NewSchemaBuilder))
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
				return op
			},
			expectError: true,
		},
		{
			name:    "invalid format",
			content: "id",
			setupOp: func() *operation.Operation {
				return operation.NewOperation(resolver.NewResolver(nil, resolver.NewSchemaBuilder))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := tt.setupOp()

			err := comment.ParseInto(tt.content, nil, op)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, op.Builder.Build().Spec)
				}
			}
		})
	}
}
