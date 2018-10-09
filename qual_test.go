package qual

import (
	"os"
	"testing"
	"time"
)

func TestFixDuration(t *testing.T) {
	for _, c := range []struct {
		complexity, max int
		exp             time.Duration
	}{
		{4, 5, 0},
		{6, 5, DefaultWeight},
		{7, 5, 2 * DefaultWeight}, // exponentially
	} {
		res := FixDuration(c.complexity, c.max)
		Assert(t, Vars{res, c.exp},
			res == c.exp,
		)
	}
}

type nop struct{}

func (t *nop) Helper()                              {}
func (t *nop) Error(args ...interface{})            {}
func (t *nop) Errorf(s string, args ...interface{}) {}
func (t *nop) Log(args ...interface{})              {}

var mock = &nop{}

func TestCyclomaticComplexity(t *testing.T) {
	CyclomaticComplexity(5, false, t)
	CyclomaticComplexity(1, true, mock)
}

func TestLineLength(t *testing.T) {
	LineLength(80, 4, false, t)
	LineLength(10, 4, false, mock)
	// And the error
	os.Chmod("qual_test.go", 0200)
	LineLength(10, 4, false, mock)
	os.Chmod("qual_test.go", 0644)
}

func TestStandard(t *testing.T) {
	Standard(t)
}

func TestHigh(t *testing.T) {
	High(mock)
}
