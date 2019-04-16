package find

import (
	"testing"
)

// Enforce test coverage without outside tools.
// Make sure this is the last test in the package.
func TestMinimumCoverage(t *testing.T) {
	exp := 100.0
	got := testing.Coverage() * 100.0
	if got != exp {
		t.Errorf("Coverage is %.2f%%, expected %.2f%%", got, exp)
	}
}
