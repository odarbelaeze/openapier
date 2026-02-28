package operation

import (
	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewParamAllowEmptyValueComment())
}

// ParamAllowEmptyValueComment sets the allowEmptyValue flag for a parameter in an operation.
type ParamAllowEmptyValueComment struct {
	paramBoolComment
}

func NewParamAllowEmptyValueComment() *ParamAllowEmptyValueComment {
	return &ParamAllowEmptyValueComment{
		paramBoolComment{
			tag: "param.allowEmptyValue",
			setter: func(param *openapi.Parameter) {
				param.AllowEmptyValue = true
			},
		},
	}
}
