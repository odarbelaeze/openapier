package operation_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/comments/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/openapi"
)

func TestParamBoolFlags_Errors(t *testing.T) {
	tests := []struct {
		name        string
		comment     operation.Comment
		content     string
		setupOp     func() *operation.Operation
		expectError string
	}{
		{
			name:    "parameter not found",
			comment: operation.NewParamAllowReservedComment(),
			content: "missing_param",
			setupOp: func() *operation.Operation {
				op := operation.NewOperation(nil)
				op.Builder.AddParameters(openapi.NewParameterBuilder().Name("id").Build())
				return op
			},
			expectError: "parameter \"missing_param\" not found, use @param missing_param ... to define it first",
		},
		{
			name:    "invalid format empty",
			comment: operation.NewParamAllowReservedComment(),
			content: "",
			setupOp: func() *operation.Operation {
				return operation.NewOperation(nil)
			},
			expectError: "invalid @param.allowReserved format, expected: @param.allowReserved <param> [param]...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := tt.setupOp()
			err := tt.comment.ParseInto(tt.content, nil, op)
			require.Error(t, err)
			assert.Equal(t, tt.expectError, err.Error())
		})
	}
}
