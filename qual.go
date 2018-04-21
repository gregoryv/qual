package qual

import (
	"container/list"
	"github.com/gregoryv/find"
	"github.com/gregoryv/gocyclo"
	"strings"
)

type T interface {
	Helper()
	Error(...interface{})
	Errorf(string, ...interface{})
}

// CyclomaticComplexity fails if max is exceeded in any go files of this project.
func CyclomaticComplexity(max int, includeVendor bool, t T) {
	t.Helper()
	found, _ := find.ByName("*.go", ".")
	var files []string
	if includeVendor {
		files = toSlice(found)
	} else {
		files = exclude("vendor/", found)
	}
	result, ok := gocyclo.Assert(files, max)
	if !ok {
		t.Errorf("Exceeded maximum complexity %v", max)
		for _, l := range result {
			t.Error(l)
		}
	}
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
