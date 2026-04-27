package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/schema/options"
)

func init() {
	Default().Register(PortTag{})
}

type PortTag struct{}

func (t PortTag) Tag() string {
	return "port"
}

func (t PortTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	switch as {
	case "integer", "number":
		return []options.SchemaOption{
			options.WithMinimum(1),
			options.WithMaximum(65535),
		}, nil
	case "string":
		return []options.SchemaOption{
			options.WithPattern("^[0-9]+$"),
		}, nil
	default:
		return nil, fmt.Errorf("port is not supported for %s", as)
	}
}

func (t PortTag) Usage() string {
	return "port"
}
