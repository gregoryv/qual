package qual

import (
	"os"
	"testing"
	"time"
)

func TestFixDuration(t *testing.T) {
	If := Wrap(t)
	for _, c := range []struct {
		complexity, max int
		exp             time.Duration
	}{
		{4, 5, 0},
		{6, 5, DefaultWeight},
		{7, 5, 2 * DefaultWeight},          // exponentially
		{21, 5, (1 << 14) * DefaultWeight}, // max, if no limit it would be
		// 1 << 15
	} {
		got := FixDuration(c.complexity, c.max)
		If(c.exp != got).Errorf(
			"Expected %v got %v for testcase %#v", c.exp, got, c,
		)
	}
}

var mock = &noopT{}

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
