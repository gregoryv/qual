package qual

import (
	"container/list"
	"github.com/gregoryv/find"
	"github.com/gregoryv/gocyclo"
	"strings"
	"testing"
)

func CyclomaticComplexity(max int, t *testing.T) {
	found, _ := find.ByName("*.go", ".")
	files := exclude("vendor/", found)
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
