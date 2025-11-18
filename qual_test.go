package qual

import (
	"os"
	"testing"
)

func TestCyclomaticComplexity(t *testing.T) {
	CyclomaticComplexity(5, false, t)
	CyclomaticComplexity(1, true, &noopT{})
}

func TestLineLength_Test(t *testing.T) {
	okCases := []struct {
		LineLength
		T
	}{
		{standardLineLength, t},
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
			T: &noopT{},
		},
		{
			LineLength: LineLength{
				MaxChars:         80,
				TabSize:          4,
				IncludeGenerated: true,
			},
			T: &noopT{},
		},
	}
	for _, c := range errCases {
		t.Run("", func(t *testing.T) {
			c.LineLength.Test(c.T)
		})
	}
	os.Chmod("qual_test.go", 0200)
	standardLineLength.Test(&noopT{})
	os.Chmod("qual_test.go", 0644)
}

func TestStandard(t *testing.T) {
	Standard(t)
}

func TestHigh(t *testing.T) {
	High(&noopT{})
}

func TestGofumpt(t *testing.T) {
	Gofumpt(t)
}

type noopT struct{}

func (t *noopT) Helper()                       {}
func (t *noopT) Error(...interface{})          {}
func (t *noopT) Errorf(string, ...interface{}) {}
