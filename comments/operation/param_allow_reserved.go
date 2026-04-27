package operation

import (
	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewParamAllowReservedComment())
}

// ParamAllowReservedComment sets the allowReserved flag for a parameter in an operation.
type ParamAllowReservedComment struct {
	paramBoolComment
}

func NewParamAllowReservedComment() *ParamAllowReservedComment {
	return &ParamAllowReservedComment{
		paramBoolComment{
			tag: "param.allowReserved",
			setter: func(param *openapi.Parameter) {
				param.AllowReserved = true
			},
		},
	}
}
