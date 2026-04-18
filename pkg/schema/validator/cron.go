package validator

import (
	"fmt"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
)

func init() {
	Default().Register(CronTag{})
}

type CronTag struct{}

func (t CronTag) Tag() string {
	return "cron"
}

func (t CronTag) Parse(value string, as string) ([]options.SchemaOption, error) {
	if as != "string" {
		return nil, fmt.Errorf("cron is not supported for %s", as)
	}

	return []options.SchemaOption{
		options.WithFormat("cron"),
	}, nil
}

func (t CronTag) Usage() string {
	return "cron"
}
