package spec

import (
	"regexp"
)

var (
	serversVariablesPattern = regexp.MustCompile(`^(\w+)\s+(.+)$`)
)
