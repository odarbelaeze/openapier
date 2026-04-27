package operation

import (
	"github.com/sv-tools/openapi"
)

func init() {
	Register(NewParamDeprecatedComment())
}

// ParamDeprecatedComment sets the deprecated flag for a parameter in an operation.
type ParamDeprecatedComment struct {
	paramBoolComment
}

func NewParamDeprecatedComment() *ParamDeprecatedComment {
	return &ParamDeprecatedComment{
		paramBoolComment{
			tag: "param.deprecated",
			setter: func(param *openapi.Parameter) {
				param.Deprecated = true
			},
		},
	}
}
