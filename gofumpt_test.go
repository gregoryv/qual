package qual

import (
	"os"
	"testing"
)

func TestGofumpt(t *testing.T) {
	Gofumpt(t)

	// unformatted files error case
	data := []byte("package x\n\n\nvar y =1\n\n")
	filename := "dummy.go"
	os.WriteFile(filename, data, 0644)
	defer os.RemoveAll(filename)
	Gofumpt(&noopT{})

	gofumptCmd = "nosuch-command"
	defer func() { gofumptCmd = "gofumpt" }()
	Gofumpt(&noopT{})
}
