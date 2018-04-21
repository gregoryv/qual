// Package find implements search funcs for finding files by name or content
package find

import (
	"container/list"
	"os"
	"path/filepath"
)

// ByName returns a list of files whose names match the shell like pattern
func ByName(pattern, root string) (result *list.List, err error) {
	sp := NewShellPattern(pattern)
	return By(sp, root)
}

// By returns a list of files whose names match
func By(m Matcher, root string) (result *list.List, err error) {
	if root == "" {
		root = "."
	}
	result = list.New()
	err = filepath.Walk(root, newVisitor(m, result))
	return
}

// Returns a visitor that skips directories
func newVisitor(m Matcher, result *list.List) func(string, os.FileInfo, error) error {
	return func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() && m.Match(f.Name()) {
			result.PushBack(path)
		}
		return nil
	}
}
