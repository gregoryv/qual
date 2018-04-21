package qual

import (
	"testing"
)

func TestCyclomaticComplexity(t *testing.T) {
	CyclomaticComplexity(5, false, t)
	CyclomaticComplexity(5, true, t)
}
