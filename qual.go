package qual

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/gregoryv/qual/internal/find"

	"github.com/gregoryv/gocyclo"
)

// High is the same as Standard, only it includes all vendor
// source as well.
func High(t T) {
	t.Helper()
	standard(true, t)
}

// Standard tests a set of metrics which might be considered necessary
// for production code. This is very opinionated, but the values are
// based on community insights from various sources.
func Standard(t T) {
	t.Helper()
	standard(false, t)
}

func standard(includeVendor bool, t T) {
	t.Helper()
	CyclomaticComplexity(5, includeVendor, t)
	LineLength(80, 4, includeVendor, t)
}

// SourceWidth fails if any go file contains lines exceeding maxChars.
// All lines are considered, source and comments.
func LineLength(maxChars, tabSize int, includeVendor bool, t T) {
	t.Helper()
	files := findGoFiles(includeVendor)
	long := &lineChecker{
		lines:    make([]string, 0),
		maxChars: maxChars,
		tab:      strings.Repeat(" ", tabSize),
	}
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
			long.check(file, line, no)
		}
	}
	long.failIfFound(t)
}

type lineChecker struct {
	lines    []string
	maxChars int
	tab      string
}

func (long *lineChecker) check(file, line string, no int) {
	line = strings.Replace(line, "\t", long.tab, -1) // tabs are 4 chars wide
	if len(line) <= long.maxChars {
		return
	}
	format := "%s:%v trim %v chars"
	long.lines = append(long.lines, fmt.Sprintf(format, file, no,
		len(line)-long.maxChars))
}

func (long *lineChecker) failIfFound(t T) {
	t.Helper()
	if len(long.lines) > 0 {
		format := "Following lines exceed the specified length %v\n%s"
		t.Errorf(format, long.maxChars, strings.Join(long.lines, "\n"))
	}
}

// CyclomaticComplexity fails if max is exceeded in any go files of
// this project.
func CyclomaticComplexity(max int, includeVendor bool, t T) {
	t.Helper()
	files := findGoFiles(includeVendor)
	result, ok := gocyclo.Assert(files, max)
	total := 0
	var totalFixDur time.Duration
	if !ok {
		t.Errorf("Exceeded maximum complexity %v", max)
		for _, l := range result {
			dur := FixDuration(l.Complexity, max)
			t.Errorf("%s (%v to fix)", l, dur)
			total += l.Complexity
			totalFixDur += dur
		}
		total -= len(result) * max
		t.Errorf("Total complexity overload %v expected to be done %v",
			total, totalFixDur)
	}
}

/*
DefaultWeight is the duration it takes to fix overloaded complexity level.
E.g. if complexity is 6 and you've set max to 5 this is the duration it
takes to fix the code from 6 to 5.

*/
var DefaultWeight = 20 * 60 * time.Second

// FixDuration calculates the duration to fix all overloaded complexity.
// Everything more complex than 14+max is timed as if 14.
func FixDuration(complexity, max int) (exp time.Duration) {
	top := complexity - max - 1
	if top > 14 {
		top = 14
	}
	return DefaultWeight * time.Duration(math.Exp2(float64(top)))
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
