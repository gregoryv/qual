package qual

import (
	"os/exec"
	"strings"
)

// Gofumpt uses gofumpt -l . command to check if files have been
// formatted.
func Gofumpt(t T) {
	t.Helper()
	// List files that would change if gofumpt were applied
	cmd := exec.Command(gofumptCmd, "-l", ".")
	out, err := cmd.Output()
	if err != nil {
		t.Error("running gofumpt failed: ", err,
			"\nRun: go install mvdan.cc/gofumpt@latest",
		)
	}

	// gofumpt prints filenames that need formatting
	files := strings.TrimSpace(string(out))
	if files != "" {
		t.Errorf(`unformatted files:
%s

gofumpt -w .`, files)
	}
}

var gofumptCmd = "gofumpt"
