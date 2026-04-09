package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/comments/operation"
	"github.com/odarbelaeze/openapier/pkg/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestParamBoolFlags(t *testing.T) {
	tests := []struct {
		name        string
		comment     operation.Comment
		tag         string
		usage       string
		content     string
		setupOp     func() *operation.Operation
		expectError bool
		validate    func(t *testing.T, op *openapi.Operation)
	}{
		{
			name:    "param.allowReserved valid",
			comment: operation.NewParamAllowReservedComment(),
			tag:     "param.allowReserved",
			usage:   "@param.allowReserved <param> [param]...",
			content: "id other",
			setupOp: func() *operation.Operation {
				op := operation.NewOperation(schema.NewResolver(nil))
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("other").Build())
				return op
			},
			validate: func(t *testing.T, op *openapi.Operation) {
				require.Len(t, op.Parameters, 2)
				assert.True(t, op.Parameters[0].Spec.Spec.AllowReserved)
				assert.True(t, op.Parameters[1].Spec.Spec.AllowReserved)
			},
		},
		{
			name:    "param.allowEmptyValue valid",
			comment: operation.NewParamAllowEmptyValueComment(),
			tag:     "param.allowEmptyValue",
			usage:   "@param.allowEmptyValue <param> [param]...",
			content: "id",
			setupOp: func() *operation.Operation {
				op := operation.NewOperation(schema.NewResolver(nil))
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
				return op
			},
			validate: func(t *testing.T, op *openapi.Operation) {
				require.Len(t, op.Parameters, 1)
				assert.True(t, op.Parameters[0].Spec.Spec.AllowEmptyValue)
			},
		},
		{
			name:    "param.deprecated valid",
			comment: operation.NewParamDeprecatedComment(),
			tag:     "param.deprecated",
			usage:   "@param.deprecated <param> [param]...",
			content: "id",
			setupOp: func() *operation.Operation {
				op := operation.NewOperation(schema.NewResolver(nil))
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
				return op
			},
			validate: func(t *testing.T, op *openapi.Operation) {
				require.Len(t, op.Parameters, 1)
				assert.True(t, op.Parameters[0].Spec.Spec.Deprecated)
			},
		},
		{
			name:    "param.required valid",
			comment: operation.NewParamRequiredComment(),
			tag:     "param.required",
			usage:   "@param.required <param> [param]...",
			content: "id",
			setupOp: func() *operation.Operation {
				op := operation.NewOperation(schema.NewResolver(nil))
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
				return op
			},
			validate: func(t *testing.T, op *openapi.Operation) {
				require.Len(t, op.Parameters, 1)
				assert.True(t, op.Parameters[0].Spec.Spec.Required)
			},
		},
		{
			name:    "parameter not found",
			comment: operation.NewParamAllowReservedComment(),
			tag:     "param.allowReserved",
			usage:   "@param.allowReserved <param> [param]...",
			content: "missing_param",
			setupOp: func() *operation.Operation {
				op := operation.NewOperation(schema.NewResolver(nil))
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
				return op
			},
			expectError: true,
		},
		{
			name:    "invalid format empty",
			comment: operation.NewParamAllowReservedComment(),
			tag:     "param.allowReserved",
			usage:   "@param.allowReserved <param> [param]...",
			content: "",
			setupOp: func() *operation.Operation {
				return operation.NewOperation(schema.NewResolver(nil))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.tag, tt.comment.Tag())
			assert.Equal(t, tt.usage, tt.comment.Usage())

			op := tt.setupOp()

			err := tt.comment.ParseInto(tt.content, nil, op)
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
