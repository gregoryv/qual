package qual

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"

	"github.com/gregoryv/gocyclo"
	"github.com/gregoryv/qual/internal/find"
)

// High is the same as Standard, only it includes all vendor
// source as well.
func High(t T) {
	t.Helper()
	CyclomaticComplexity(5, true, t)
	standardLineLength.Test(t)
}

// Standard tests a set of metrics which might be considered necessary
// for production code. This is very opinionated, but the values are
// based on community insights from various sources.
func Standard(t T) {
	t.Helper()
	CyclomaticComplexity(5, false, t)
	standardLineLength.Test(t)
}

var standardLineLength = LineLength{
	MaxChars: 80,
	TabSize:  4,
}

type LineLength struct {
	MaxChars         int
	TabSize          int
	IncludeVendor    bool
	IncludeGenerated bool
	// Set to true if lines with urls should be considered
	IncludeURLs bool

	failedLines []string
}

// LineLength fails if any go file contains lines exceeding maxChars.
// All lines are considered, source and comments.
func (l *LineLength) Test(t T) {
	t.Helper()
	files := findGoFiles(l.IncludeVendor)

	for _, file := range files {
		fh, err := os.Open(file)
		if err != nil {
			t.Error(err)
		}
		l.checkFile(file, fh)
		fh.Close()
	}
	l.failIfFound(t)
}

func (l *LineLength) checkFile(file string, fh *os.File) {
	scanner := bufio.NewScanner(fh)
	no := 0
	tab := strings.Repeat(" ", l.TabSize)
	for scanner.Scan() {
		no++
		line := scanner.Text()
		if !l.IncludeGenerated && strings.Contains(line, "DO NOT EDIT") {
			return
		}
		l.check(file, line, tab, no)
	}
}

func (l *LineLength) check(file, line, tab string, no int) {
	line = strings.ReplaceAll(line, "\t", tab)
	if len(line) <= l.MaxChars {
		return
	}
	if !l.IncludeURLs && lineContainsURL(line) {
		return
	}
	format := "%s:%v trim %v chars"
	l.failedLines = append(l.failedLines, fmt.Sprintf(format, file, no,
		len(line)-l.MaxChars))
}

func lineContainsURL(line string) bool {
	// e.g. https://example.com/with/a/very/long/pathname/we/really/want/to/keep/on/one/line
	switch {
	case strings.Contains(line, "https://"):
	case strings.Contains(line, "http://"):
	default:
		return false
	}
	return true
}

func (l *LineLength) failIfFound(t T) {
	t.Helper()
	if len(l.failedLines) > 0 {
		format := "Following lines exceed the specified length %v\n%s"
		t.Errorf(format, l.MaxChars, strings.Join(l.failedLines, "\n"))
	}
}

// CyclomaticComplexity fails if max is exceeded in any go files of
// this project.
func CyclomaticComplexity(max int, includeVendor bool, t T) {
	t.Helper()
	files := findGoFiles(includeVendor)
	result, ok := gocyclo.Assert(files, max)
	total := 0
	if !ok {
		t.Errorf("Exceeded maximum complexity %v", max)
		for _, l := range result {
			total += l.Complexity
		}
		total -= len(result) * max
	}
}

func findGoFiles(includeVendor bool) (result []string) {
	found, _ := find.ByName("*.go", ".")
	if includeVendor {
		return toSlice(found)
	}
	return exclude("vendor"+string(os.PathSeparator), found)
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
