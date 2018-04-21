package qual

import (
	"testing"
)

type sssh struct{}

func (t *sssh) Helper()                              {}
func (t *sssh) Error(args ...interface{})            {}
func (t *sssh) Errorf(s string, args ...interface{}) {}

func TestCyclomaticComplexity(t *testing.T) {
	CyclomaticComplexity(5, false, t)
	CyclomaticComplexity(1, true, &sssh{})
}
