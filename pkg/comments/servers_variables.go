package comments

import (
	"regexp"
)

var (
	serversVariablesPattern = regexp.MustCompile(`^(\w+)\s+(.+)$`)
)
