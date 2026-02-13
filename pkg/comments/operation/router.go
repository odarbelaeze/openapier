package operation

import (
	"fmt"
	"regexp"
)

var (
	_             Comment = routerComment{}
	routerPattern         = regexp.MustCompile(`^(/[\w./\-{}\(\)+:$]*)[[:blank:]]+\[(\w+)]`)
)

func init() {
	Register(NewRouterComment())
}

type routerComment struct{}

func NewRouterComment() *routerComment {
	return &routerComment{}
}

// ParseInto implements [Comment].
func (r routerComment) ParseInto(c string, op *Operation) error {
	matches := routerPattern.FindStringSubmatch(c)
	if len(matches) != 3 {
		return fmt.Errorf("invalid router comment format: %s", c)
	}
	path := matches[1]
	method := matches[2]
	op.Routes = append(op.Routes, Route{Path: path, Method: method})
	return nil
}

// Tag implements [Comment].
func (r routerComment) Tag() string {
	return "router"
}

// Usage implements [Comment].
func (r routerComment) Usage() string {
	return "// @router <path> [method]"
}
