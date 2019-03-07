package find

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"testing"
)

var testRoot string // set by TestSuite

func init() {
	var err error
	// Setup directory structure for tests
	testRoot, err = ioutil.TempDir("", "search_test")
	if err != nil {
		panic(err)
	}
	content := []struct {
		path  string
		isDir bool
	}{
		{"cars/", true},
		{"a.txt", false},
		{"b.txt", false},
	}
	for _, c := range content {
		full := path.Join(testRoot, c.path)
		if c.isDir {
			os.MkdirAll(full, 0755)
		} else {
			ioutil.WriteFile(full, []byte{}, 0755)
		}
	}
}

func TestSuite(t *testing.T) {
	// Run tests
	os.Chdir(testRoot)
	t.Run("By", testBy)
	t.Run("ByName", testByName)
}

func ExampleByName() {
	os.Chdir(testRoot)
	result, _ := ByName("*.txt", ".")
	for e := result.Front(); e != nil; e = e.Next() {
		if s, ok := e.Value.(string); ok {
			fmt.Println(s)
		}
	}
	//output:
	//a.txt
	//b.txt
}

func testBy(t *testing.T) {
	data := []struct {
		m     Matcher
		root  string
		count int
	}{
		{NewRegexp(regexp.MustCompile(`.*\.txt`)), ".", 2},
	}
	for _, d := range data {
		result, _ := By(d.m, d.root)
		if result.Len() != d.count {
			t.Errorf("By(%q, %q) expected to find %v files, found %v",
				d.m, d.root, d.count, result.Len())
		}
	}
}

func testByName(t *testing.T) {
	data := []struct {
		pattern string
		root    string
		count   int
		ok      bool
	}{
		{"a.txt", ".", 1, true},
		{"*.txt", "", 2, true}, // no directory means "."
		{"x", ".", 0, false},
		{"whatever", "nosuchdir", 0, false},
		{"", ".", 0, false},
	}
	for _, d := range data {
		result, err := ByName(d.pattern, d.root)
		if d.ok && err != nil {
			t.Errorf("ByName(%q, %q): %s", d.pattern, d.root, err)
		}
		if result.Len() != d.count {
			t.Errorf("ByName(%q, %q) expected to find %v files, found %v",
				d.pattern, d.root, d.count, result.Len())
		}
	}
}
