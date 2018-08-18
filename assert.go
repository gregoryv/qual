package qual

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Vars groups variables which will be logged on test failures
type Vars []interface{}

func Assert(t T, v Vars, checks ...bool) (failed bool) {
	t.Helper()
	return assert(t, above(2), v, checks...)
}

func assert(t T, msg string, v Vars, checks ...bool) (failed bool) {
	t.Helper()
	for i, ok := range checks {
		if ok {
			continue
		}
		if !failed {
			t.Errorf("> %s", msg)
			failed = true
		}
		t.Errorf("  failed assert: %s", trueCase(scanLine(3, 0), i+1))
	}
	if failed {
		// Log all Vars{...} with name and value
		logVars(t, v, strings.Join(funcArgs(3), ","))
	}
	return
}

// returns the variable name of the calling func
func trueCase(str string, nth int) string {
	if assertStatementIsOnSameLine(str) {
		i := strings.Index(str, "}")
		j := strings.Index(str, ")")
		parts := strings.Split(str[i:j], ",")
		return strings.TrimSpace(parts[nth])
	}
	str = scanLine(4, nth)
	// Assuming they are on the same line here
	j := strings.Index(str, ",")
	// if j is -1 then the compiler should fail
	return strings.TrimSpace(str[:j])
}

func assertStatementIsOnSameLine(str string) bool {
	return strings.Index(str, "Assert(") >= 0 && strings.LastIndex(str, ")") > 0
}

func logVars(t T, v Vars, parts string) {
	t.Helper()
	i := strings.Index(parts, "{") + 1
	j := strings.Index(parts, "}")
	vars := strings.Split(parts[i:j], ",")
	for i, v := range v {
		var val string
		switch v := v.(type) {
		case string, error:
			val = fmt.Sprintf("%q", v)
		default:
			val = fmt.Sprintf("%v", v)
		}
		t.Log("   ", strings.TrimSpace(vars[i]), "=", val)
	}
}

// returns the line above the caller
func above(nth int) string {
	str := scanLine(nth+1, -1)
	return strings.TrimSpace(str)
}

func funcArgs(n int) []string {
	str := scanLine(n+1, 0)
	i := strings.Index(str, "(") + 1
	// Assuming they are on the same line here
	j := strings.Index(str, ")")
	if j == -1 {
		return strings.Split(str[i:], ",")
	}
	return strings.Split(str[i:j], ",")
}

func scanLine(caller, back int) string {
	_, file, line, _ := runtime.Caller(caller) // todo, handle error
	fh, _ := os.Open(file)
	scanner := bufio.NewScanner(fh)
	for i := 0; i < line+back; i++ {
		scanner.Scan()
	}
	fh.Close()
	return scanner.Text()
}
