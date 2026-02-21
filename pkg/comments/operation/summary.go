package operation

import (
	"strings"
)

func init() {
	Register(NewSummaryComment())
}

// SummaryComment sets the summary of an operation.
type SummaryComment struct{}

func NewSummaryComment() *SummaryComment {
	return &SummaryComment{}
}

func (c *SummaryComment) Tag() string {
	return "summary"
}

func (c *SummaryComment) Usage() string {
	return "@summary <summary>"
}

func (c *SummaryComment) ParseInto(content string, op *Operation) error {
	summary := strings.TrimSpace(content)
	if summary != "" {
		op.Builder.Summary(summary)
	}
	return nil
}
