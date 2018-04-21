package qual

import (
	"bufio"
	"container/list"
	"github.com/gregoryv/find"
	"github.com/gregoryv/gocyclo"
	"os"
	"strings"
)

type T interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
}

// High is the same as Standard, only it includes all vendor
// source as well.
func High(t T) {
	standard(true, t)
}

// Standard tests a set of metrics which might be considered necessary
// for production code. This is very opinionated, but the values are
// based on community insights from various sources.
func Standard(t T) {
	standard(false, t)
}

func standard(includeVendor bool, t T) {
	CyclomaticComplexity(5, includeVendor, t)
	LineLength(80, includeVendor, t)
}

// SourceWidth fails if any go file contains lines exceeding maxChars.
// All lines are considered, source and comments.
func LineLength(maxChars int, includeVendor bool, t T) {
	t.Helper()
	files := findGoFiles(includeVendor)
	for _, file := range files {
		fh, err := os.Open(file)
		if err != nil {
			t.Error(err)
		}
		scanner := bufio.NewScanner(fh)
		no := 0
		for scanner.Scan() {
			no++
			line := scanner.Text()
			if len(line) > maxChars {
				t.Errorf("%s:%v %s...", file, no, line[:maxChars])
			}
		}

	}
}

// CyclomaticComplexity fails if max is exceeded in any go files of
// this project.
func CyclomaticComplexity(max int, includeVendor bool, t T) {
	t.Helper()
	files := findGoFiles(includeVendor)
	result, ok := gocyclo.Assert(files, max)
	if !ok {
		t.Errorf("Exceeded maximum complexity %v", max)
		for _, l := range result {
			t.Error(l)
		}
	}
}

func findGoFiles(includeVendor bool) (result []string) {
	found, _ := find.ByName("*.go", ".")
	if includeVendor {
		return toSlice(found)
	}
	return exclude("vendor/", found)
}

func exclude(pattern string, files *list.List) (result []string) {
	for e := files.Front(); e != nil; e = e.Next() {
		s, _ := e.Value.(string)
		if !strings.Contains(s, pattern) {
			result = append(result, s)
		}
	}
	return
}

func toSlice(files *list.List) (result []string) {
	for e := files.Front(); e != nil; e = e.Next() {
		s, _ := e.Value.(string)
		result = append(result, s)
	}
	return
}
