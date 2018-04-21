package qual

import (
	"testing"
)

type nop struct{}

func (t *nop) Helper()                              {}
func (t *nop) Error(args ...interface{})            {}
func (t *nop) Errorf(s string, args ...interface{}) {}

func TestCyclomaticComplexity(t *testing.T) {
	CyclomaticComplexity(5, false, t)
	CyclomaticComplexity(1, true, &nop{})
}

func TestSourceWidth(t *testing.T) {
	SourceWidth(80, false, t)
	SourceWidth(10, false, &nop{})
}
