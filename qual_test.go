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
		{7, 5, 2 * DefaultWeight},          // exponentially
		{21, 5, (1 << 14) * DefaultWeight}, // max, if no limit it would be
		// 1 << 15
	} {
		got := FixDuration(c.complexity, c.max)
		if got != c.exp {
			t.Errorf("Expected %v got %v for testcase %#v", c.exp, got, c)
		}
	}
}

var mock = &noopT{}

func TestCyclomaticComplexity(t *testing.T) {
	CyclomaticComplexity(5, false, t)
	CyclomaticComplexity(1, true, mock)
}

func TestLineLength_Test(t *testing.T) {
	okCases := []struct {
		LineLength
		T
	}{
		{StandardLineLength, t},
	}
	for _, c := range okCases {
		t.Run("", func(t *testing.T) {
			c.LineLength.Test(c.T)
		})
	}

	errCases := []struct {
		LineLength
		T
	}{
		{
			LineLength: LineLength{
				MaxChars: 10,
				TabSize:  4,
			},
			T: mock,
		},
	}
	for _, c := range errCases {
		t.Run("", func(t *testing.T) {
			c.LineLength.Test(c.T)
		})
	}
	os.Chmod("qual_test.go", 0200)
	StandardLineLength.Test(mock)
	os.Chmod("qual_test.go", 0644)
}

func TestStandard(t *testing.T) {
	Standard(t)
}

func TestHigh(t *testing.T) {
	High(mock)
}
