package find

import (
	"path/filepath"
	"regexp"
)

type Matcher interface {
	Match(path string) bool
}

type shellPattern struct {
	pattern string
}

func NewShellPattern(pattern string) Matcher {
	return &shellPattern{pattern: pattern}
}

func (sp *shellPattern) Match(path string) bool {
	res, _ := filepath.Match(sp.pattern, path)
	return res
}

type reg struct {
	ex *regexp.Regexp
}

func NewRegexp(ex *regexp.Regexp) Matcher {
	return &reg{ex: ex}
}

func (rm *reg) Match(path string) bool {
	return rm.ex.Match([]byte(path))
}
