package find_test

import (
	"github.com/gregoryv/find"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestInFile(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "grep_test")
	ioutil.WriteFile(file.Name(), []byte(`
a hello
a world
`), 0644)
	data := []struct {
		pattern, file string
		exp           string
		ok            bool
	}{
		{"a hello", file.Name(), "1:a hello", true},
		{"a*", file.Name(), "1:a hello,2:a world", true},
		{"*", "nosuchfile", "", false},
	}

	for _, d := range data {
		res, err := find.InFile(d.pattern, d.file)
		// Convert to one string for comparison
		result := ""
		if res != nil {
			lines := make([]string, 0, res.Len())
			for e := res.Front(); e != nil; e = e.Next() {
				if ref, ok := e.Value.(*find.Ref); ok {
					lines = append(lines, ref.String())
				} else {
					lines = append(lines, "no")
				}
			}
			result = strings.Join(lines, ",")
		}
		// Assert
		if d.ok && (res == nil || result != d.exp) {
			t.Errorf("Grep(%q, %q) expected \n%v\n, got\n %v", d.pattern, d.file, d.exp, result)
		}
		if !d.ok && err == nil {
			t.Errorf("Grep(%q, %q) expected to fail", d.pattern, d.file)
		}
	}
}
