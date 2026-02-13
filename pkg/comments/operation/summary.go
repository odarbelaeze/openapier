package operation

import "strings"

func init() {
	Register(NewSummaryComment())
}

// SummaryComment is a comment that updates the summary of an operation.
type SummaryComment struct{}

func NewSummaryComment() *SummaryComment {
	return &SummaryComment{}
}

func (s *SummaryComment) Tag() string {
	return "summary"
}

func (s *SummaryComment) Usage() string {
	return "@summary <summary>"
}

func (s *SummaryComment) ParseInto(content string, op *Operation) error {
	op.Builder.Summary(strings.TrimSpace(content))
	return nil
}
