package operation

import (
	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewParamRequiredComment())
}

// ParamRequiredComment sets the required flag for a parameter in an operation.
type ParamRequiredComment struct {
	paramBoolComment
}

func NewParamRequiredComment() *ParamRequiredComment {
	return &ParamRequiredComment{
		paramBoolComment{
			tag: "param.required",
			setter: func(param *openapi.Parameter) {
				param.Required = true
			},
		},
	}
}
