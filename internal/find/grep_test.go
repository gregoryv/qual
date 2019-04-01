package find

import (
	"container/list"
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
	}{
		{"a hello", file.Name(), "1:a hello"},
		{"a*", file.Name(), "1:a hello,2:a world"},
	}

	for _, d := range data {
		res, err := InFile(d.pattern, d.file)
		if err != nil {
			t.Fatal(err)
		}
		// Convert to one string for comparison
		result := toString(res)

		// Assert
		if res == nil || result != d.exp {
			t.Errorf("InFile(%q, %q) expected \n%v\n, got\n %v",
				d.pattern, d.file, d.exp, result)
		}
	}
}

func toString(res *list.List) string {
	if res == nil {
		return ""
	}
	lineChecker := make([]string, 0, res.Len())
	for e := res.Front(); e != nil; e = e.Next() {
		if ref, ok := e.Value.(*Ref); ok {
			lineChecker = append(lineChecker, ref.String())
		} else {
			lineChecker = append(lineChecker, "no")
		}
	}
	return strings.Join(lineChecker, ",")
}

func Test_missing_file(t *testing.T) {
	_, err := InFile("*", "nosuchfile")
	if err == nil {
		t.Error("Expected find.InFile to fail when no such file exists")
	}
}
